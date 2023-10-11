package main

import (
	"github.com/dias-oblivion/PicPay-Simplificado/api/router"
	"github.com/dias-oblivion/PicPay-Simplificado/config"
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
