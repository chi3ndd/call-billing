package main

import (
	"os"
	"os/signal"
	"syscall"

	"call-billing/defs"
	"call-billing/handler"
	"call-billing/model"

	"gopkg.in/yaml.v3"
)

func main() {
	// Read config
	configFile := os.Getenv(defs.EnvConfigFilePath)
	if configFile == "" {
		configFile = defs.DefaultConfigFilePath
	}
	f, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var conf *model.Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&conf); err != nil {
		panic(err)
	}
	// Start
	go handler.APIHandlerStart(conf)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-ch:
		break
	}
}
