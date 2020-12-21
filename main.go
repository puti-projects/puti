package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/puti-projects/puti/internal/web/service"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/counter"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/pkg/theme"
	v "github.com/puti-projects/puti/internal/pkg/version"
	"github.com/puti-projects/puti/internal/routers"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme/autocert"
)

var (
	configPath = pflag.StringP("config", "c", "", "Puti config file path.")
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

	// load theme path
	theme.LoadInstalled()
}

func main() {
	// load cache service
	if err := cache.LoadCache(); err != nil {
		logger.Errorf("init cache failed. %s", err)
	} else {
		logger.Info("cache service has been deployed successfully")
	}

	// load default options (need db connection)
	if err := cache.LoadOptions(); err != nil {
		logger.Panicf("load options failed, %v", err)
	}
	logger.Info("options has been deployed successfully")

	// new service engine for frontend as a global engine
	if err := service.NewServiceEngine(); err != nil {
		logger.Panicf("new service engine failed, %v", err)
	}
	logger.Info("new service engine successfully")

	// routers
	router := routers.NewRouter(config.Server.Runmode)

	// Ping the server to make sure the router is working.
	// should before http server set up
	go func() {
		pingServer()
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

// httpHandle handle HTTP
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

// httpsHandle handle HTTPS; there are two situation
// Situation 1. Open auto cert.
// Situation 2. Specify certification path.
func httpsHandle(router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + config.Server.HttpsPort,
		Handler: router,
	}

	if config.Server.AutoCert {
		// Open auto cert
		// auto cert manager
		m := &autocert.Manager{
			Cache:      autocert.DirCache(config.StaticPath("configs/cert/")),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.Server.PutiDomain...),
		}
		// set auto cert config to tls config
		srv.TLSConfig = m.TLSConfig()

		// Listen and serve
		serveTLS(srv, "", "")
	} else {
		// Specify certification path
		if config.Server.TlsCert == "" || config.Server.TlsKey == "" {
			logger.Errorf("https opened but cert and key can not be empty, failed to listen https port")
		}

		// Listen and serve
		serveTLS(srv, config.Server.TlsCert, config.Server.TlsKey)
	}
	return srv
}

// serveTLS serve https for all situation
func serveTLS(srv *http.Server, certFile string, keyFile string) {
	go func() {
		logger.Info("start to listening the incoming https requests", zap.String("port", config.Server.HttpsPort))
		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server.ListenAndServeTLS err: %v", err)
		}
	}()
}

// signalHandle graceful shutdown based on http.server.Shutdown
func signalHandle(srv *http.Server) {
	quit := make(chan os.Signal)
	// receive syscall.SIGINT and syscall.SIGTERM signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// if signal received
	<-quit
	logger.Warn("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %v; the service will be forced to quit", err)
	}
	logger.Warn("server shutdown")
}

// pingServer pings the http server to make sure the service is working.
func pingServer() {
	var pingURL string
	if true == config.Server.HttpsOpen {
		if config.Server.AutoCert {
			pingURL = "https://" + config.Server.PutiDomain[0] + "/check/health"
		} else {
			pingURL = "https://127.0.0.1:" + config.Server.HttpsPort + "/check/health"
		}
	} else {
		pingURL = "http://127.0.0.1:" + config.Server.HttpPort + "/check/health"
	}

	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(pingURL)
		if err == nil && resp.StatusCode == 200 {
			logger.Info("health check finished and the HTTP service is normal.", zap.String("ping url", pingURL))
			logger.Info("the router has been deployed successfully")
			return
		}

		// Sleep for a second to continue the next ping.
		logger.Warn("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	logger.Error("cannot connect to the router! The router has no response, or it might took too long to start up.", zap.String("ping url", pingURL))
	return
}
