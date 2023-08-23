// микросервис проверки комментариев
package main

import (
	_ "contracts/censor/docs"
	"contracts/censor/pkg/api"
	"contracts/censor/pkg/config"
	logs "contracts/internal/log"
	tools "contracts/internal/tools"
	"fmt"
	"net/http"
)

// @title Check Contract Hydra API
// @version 1.0
// @description  Получение списка  договоров из Биллинга
// @tag.name  Contracts
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email ataradaev@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func main() {
	//парсим командную строку
	a, err := tools.ParseFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if a.Help {
		fmt.Println(config.Help())
		return
	}
	// загрузка конфига
	conf := config.New()
	err = conf.Load(a.Fileconfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	//инициализация логирования
	log := logs.New()
	logs.InitConfig(conf.Log(), true)
	defer logs.Close()
	log.Info("config loaded, start program....")
	log.Info("Path to app:", tools.GetAppExe())
	log.Infoln("cmdline: ", a)

	log.Debugln("loaded config ", conf)
	// инициализация маршрутизатора HTTP
	log.Info("Init Http router")
	api, err := api.New()
	if err != nil {
		log.Fatal(err)
	}
	Port := fmt.Sprintf(":%d", conf.Port())

	log.Info("Init Http server on ", Port)
	err = http.ListenAndServe(Port, api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
