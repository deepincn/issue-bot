package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/projectdde/issue-bot/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}

var (
	h bool
	c string
	d bool
)

func main() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&c, "c", "", "config file")
	flag.BoolVar(&d, "d", false, "debug mode")

	conf := new(config.Yaml)
	yamlFile, err := ioutil.ReadFile(c)

	if err != nil {
		logrus.Infof("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal([]byte(yamlFile), conf)
	if err != nil {
		logrus.Fatalf("Unmarshal: %v", err)
	}

	handle := WebhookHandle{}

	db := DataBase{}
	db.Ping()

	router := gin.Default()
	router.POST("/webhook", handle.WebhookHandle)
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:3002",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Fatal(srv.ListenAndServe())
}
