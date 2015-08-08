package main

import (
	"flag"
	"fmt"
	"github.com/flike/kingshard/config"
	"github.com/flike/kingshard/core/golog"
	"github.com/flike/kingshard/proxy/server"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

var configFile *string = flag.String("config", "/etc/kingshard.conf", "kingshard config file")
var logLevel *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")

const banner string = `
    __   _                  __                   __
   / /__(_)___  ____ ______/ /_  ____ __________/ /
  / //_/ / __ \/ __ \/ ___/ __ \ / __\/ ___/ __  /
 / ,< / / / / / /_/ (__  ) / / / /_/ / /  / /_/ /
/_/|_/_/_/ /_/\__, /____/_/ /_/\__,_/_/   \__,_/
             /____/
`

func main() {
	fmt.Print(banner)
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if len(*configFile) == 0 {
		fmt.Println("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}

	if *logLevel != "" {
		setLogLevel(*logLevel)
	} else {
		setLogLevel(cfg.LogLevel)
	}

	var svr *server.Server
	svr, err = server.NewServer(cfg)
	if err != nil {
		golog.Error("main", "main", err.Error(), 0)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		golog.Info("main", "main", "Got signal", 0, "signal", sig)
		golog.GlobalLogger.Close()
		svr.Close()
	}()

	svr.Run()
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		golog.GlobalLogger.SetLevel(golog.LevelDebug)
	case "info":
		golog.GlobalLogger.SetLevel(golog.LevelInfo)
	case "warn":
		golog.GlobalLogger.SetLevel(golog.LevelWarn)
	case "error":
		golog.GlobalLogger.SetLevel(golog.LevelError)
	default:
		golog.GlobalLogger.SetLevel(golog.LevelError)
	}
}
