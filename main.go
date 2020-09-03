package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

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

	// load theme path
	theme.LoadInstalled()

	// Set gin mode.
	if "debug" == config.Server.Runmode {
		gin.SetMode(gin.DebugMode)
	} else if "test" == config.Server.Runmode {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	// init db
	if err := db.InitDB(); err != nil {
		logger.Errorf("sql.Open() error(%v)", err)
		panic(fmt.Sprintf("database connection failed. err: %v", err))
	}
	defer db.DBEngine.Close()

	// load default options
	option.LoadOptions()

	// routers
	router := routers.NewRouter()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up. Error Detail:" + err.Error())
		}
		logger.Info("the router has been deployed successfully")
	}()

	// init ticker
	counter.InitCountTicker()

	// If open https, start listening https request
	if true == config.Server.HttpsOpen {
		cert := config.Server.TlsCert
		key := config.Server.TlsKey
		if cert != "" && key != "" {
			go func() {
				logger.Info("start to listening the incoming https requests", zap.String("port", config.Server.HttpsPort))
				logger.Info(
					http.ListenAndServeTLS(
						"0.0.0.0:"+config.Server.HttpsPort,
						cert,
						key,
						router,
					).Error())
			}()
		} else {
			logger.Errorf("cert and key can not be empty, failed to listen https port")
		}
	}
	logger.Info("start to listening the incoming http requests", zap.String("port", config.Server.HttpPort))
	logger.Info(http.ListenAndServe(
		"0.0.0.0:"+config.Server.HttpPort,
		router,
	).Error())
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
