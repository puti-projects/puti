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
	"github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/pkg/tickers"
	v "github.com/puti-projects/puti/internal/pkg/version"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configPath = pflag.StringP("config", "c", "", "puti config file path.")
	version    = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
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

	// init config
	if err := config.Init(*configPath); err != nil {
		panic(fmt.Errorf("fatal error init configuration: %s", err))
	}
	logger.Info("configuration load succeeded.", zap.String("config file", viper.ConfigFileUsed()))

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// load default options
	option.LoadOptions()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// create the gin engine
	g := gin.New()

	// routes
	router.Load(g)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		logger.Info("The router has been deployed successfully.")
	}()

	// If open https, start listening https request
	openHTTPS := viper.GetBool("tls.https_open")
	if openHTTPS == true {
		cert := viper.GetString("tls.cert")
		key := viper.GetString("tls.key")
		if cert != "" && key != "" {
			go func() {
				logger.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
				logger.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
			}()
		}
	}

	// init ticker
	tickers.InitCountTicker()
	logger.Info("Start to running the count ticker.")

	logger.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	logger.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router")
}
