package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/puti-projects/puti/internal/common/config"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/router"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/pkg/theme"
	"github.com/puti-projects/puti/internal/pkg/tickers"
	v "github.com/puti-projects/puti/internal/pkg/version"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configPath = pflag.StringP("config", "c", "", "puti config file path.")
	version    = pflag.BoolP("version", "v", false, "show version info.")
)

// init function
func init() {
	pflag.Parse()

	// Show version info
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// load config include logger
	config.Init(*configPath)

	// load theme path
	theme.LoadInstalled()

	// Set gin mode.
	if "debug" == viper.GetString("runmode") {
		gin.SetMode(gin.DebugMode)
	} else if "test" == viper.GetString("runmode") {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	// init db
	model.DB.Init()
	defer model.DB.Close()

	// load default options
	option.LoadOptions()

	// create the gin engine
	g := gin.New()

	// routes
	router.Load(g)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up", err)
		}
		logger.Info("the router has been deployed successfully")
	}()

	// init ticker
	tickers.InitCountTicker()

	// If open https, start listening https request
	if true == viper.GetBool("tls.https_open") {
		cert := viper.GetString("tls.cert")
		key := viper.GetString("tls.key")
		if cert != "" && key != "" {
			go func() {
				logger.Info("start to listening the incoming https requests", zap.String("port", viper.GetString("tls.addr")))
				logger.Info(http.ListenAndServeTLS("0.0.0.0:"+viper.GetString("tls.addr"), cert, key, g).Error())
			}()
		} else {
			logger.Errorf("cert and key can not be empty, failed to listen https port")
		}
	}
	logger.Info("start to listening the incoming http requests", zap.String("port", viper.GetString("addr")))
	logger.Info(http.ListenAndServe("0.0.0.0:"+viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the service is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("ping_max_num"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("ping_url") + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to the router")
}
