package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/counter"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/pkg/theme"
	v "github.com/puti-projects/puti/internal/pkg/version"
	"github.com/puti-projects/puti/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	configPath = pflag.StringP("config", "c", "", "puti config file path.")
	version    = pflag.BoolP("version", "v", false, "show version info.")
)

// init function
func init() {
	pflag.Parse()

	// if a -v was receive, show version info
	if *version {
		versionParams := v.Get()
		marshalled, err := json.MarshalIndent(&versionParams, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// set up config
	err := config.InitConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("setupConfig err: %v", err))
	}

	// set up logger
	logger.InitLogger(config.Server.Runmode)
	logger.Info("logger construction succeeded")

	// init db
	err = db.InitDB()
	if err != nil {
		logger.Panicf("database connection failed. error(%v)", err)
	}
	// if init db without problem, set up dao engine
	dao.NewDaoEngine()

	// load theme path
	theme.LoadInstalled()
}

func main() {
	// load default options (need db connection)
	if err := option.LoadOptions(); err != nil {
		logger.Errorf("load options failed, %v", err)
	}
	logger.Info("options has been deployed successfully")

	// routers
	router := routers.NewRouter(config.Server.Runmode)

	// Ping the server to make sure the router is working.
	// should before http server set up
	go func() {
		if err := pingServer(); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up. Error Detail:" + err.Error())
		}
		logger.Info("the router has been deployed successfully")
	}()

	// init ticker
	counter.InitCountTicker()

	// listen and serve http
	httpServe(router)
}

// httpServe set up http server
// If https open, should only listen https port
func httpServe(router *gin.Engine) {
	var srv *http.Server
	// if open https
	if true == config.Server.HttpsOpen {
		srv = httpsHandle(router)
	} else {
		srv = httpHandle(router)
	}

	signalHandle(srv)
}

func httpHandle(router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + config.Server.HttpPort,
		Handler: router,
	}

	go func() {
		logger.Info("start to listening the incoming http requests", zap.String("port", config.Server.HttpPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server.ListenAndServe err: %v", err)
		}
	}()

	return srv
}

func httpsHandle(router *gin.Engine) *http.Server {
	if config.Server.TlsCert == "" || config.Server.TlsKey == "" {
		logger.Errorf("https opened but cert and key can not be empty, failed to listen https port")
	}

	srv := &http.Server{
		Addr:    ":" + config.Server.HttpsPort,
		Handler: router,
	}

	go func() {
		logger.Info("start to listening the incoming https requests", zap.String("port", config.Server.HttpsPort))
		if err := srv.ListenAndServeTLS(config.Server.TlsCert, config.Server.TlsKey); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server.ListenAndServeTLS err: %v", err)
		}
	}()

	return srv
}

func signalHandle(srv *http.Server) {
	quit := make(chan os.Signal)
	// receive syscall.SIGINT and syscall.SIGTERM signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// if signal received
	<-quit
	logger.Warn("shuting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %v; the service will be forced to quit", err)
	}
	logger.Warn("server shutdown")
}

// pingServer pings the http server to make sure the service is working.
func pingServer() error {
	for i := 0; i < config.Server.PingMaxNum; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(config.Server.PingUrl + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to the router")
}
