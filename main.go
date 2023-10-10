package main

import (
	"github.com/dias-oblivion/PicPay-Simplificado/config"
	"github.com/dias-oblivion/PicPay-Simplificado/router"
)

func main() {

	// init main logger
	logger := config.GetLogger("main")

	// init configs
	configError := config.Init()

	if configError != nil {
		logger.Errorf("error when initializing configs: %v", configError)
		return
	}

	// router initialize
	router.InitRouter()
}
