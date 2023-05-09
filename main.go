package main

import (
	"flag"
	"fmt"
	"github.com/TrHung-297/chat-v2/server"
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	env.SetupConfigEnv()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`Debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	} else {
		fmt.Println("Service RUN on PRODUCTION mode")
	}
}

func main()  {
	flag.Parse()
	app := fiber.New()
	server.InitializeApp(app)

	app.Listen("127.0.0.1:9099")
}
