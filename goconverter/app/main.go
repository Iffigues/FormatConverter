package main

import (
	"context"
	"converter/conf"
	"converter/logger"
	"converter/pklhandler"
	"converter/server"
	"converter/uploader"
	"log"
)

func main() {
	confs, err := conf.NewConf(context.Background(), "pkl/local/appConfig.pkl")
	if err != nil {
		log.Fatal(err)
	}
	logs := logger.NewLog(confs.Cfg.LogDir)

	if err != nil {
		log.Fatal(err)
	}
	serve, err := server.NewServer(confs, logs)
	if err != nil {
		logs.Fatal(err.Error())
		return
	}
	p := pklhandler.NewPKLHandler(confs)
	serve.AddRoute(p)
	c := uploader.NewPUploader(confs)
	serve.AddRoute(c)
	serve.Start()
	for {
		serve.Serve()
	}
}
