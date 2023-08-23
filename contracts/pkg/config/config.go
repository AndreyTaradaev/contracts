// реализация файла конфигурации создает единственный экземпляр .
package config

import (
	"contracts/internal/tools"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var apiconfig config

// члены структуры закрыты для доступа и изменения , доступ к значениям через геттер
type config struct {
	conf config_imp
}

// пряиой доступ к членам запрещен
type config_imp struct {
	Port     int    `json:"port,omitempty"`     // порт сервиса  по умолчанию 8080
	Logdir   string `json:"logdir,omitempty"`   //  каталог куда пишем лог ,пустой пишем в stdout
	Dbhost   string `json:"dbhost,omitempty"`   //  host  DB   "172.16.0.30",
	Dbport   int    `json:"dbport,omitempty"`   // порт DB    1521,
	Database string `json:"database,omitempty"` // база данных  hydra,
	Dbuser   string `json:"dbuser,omitempty"`   // пользователь db "ais_rpc",
	Dbpasswd string `json:"dbpasswd,omitempty"` // пароль пользователя
}

// справка по программе.
func Help() string {
	ex, _ := os.Executable()
	var h string
	h = `программма ` + tools.GetBase(ex) + ` является  микросервисем получения контрактов . 

	 `
	return h
}

//=====================================================================================

// ctor for config-object
func New() *config {
	once.Do(func() {
		apiconfig = config{
			conf: config_imp{
				Port:     8080,
				Logdir:   "",
				Dbhost:   "172.16.50.2",
				Dbport:   1521,
				Database: "hydra",
				Dbuser:   "ais_rpc",
				Dbpasswd: ""},
		}
	},
	)
	return &apiconfig
}

// загрузка переменных в среду оуружения
/* func (c *config) loadFromEnv() {

	variable := os.Getenv("CENSORPORT")
	if len(variable) != 0 {
		p, err := strconv.Atoi(variable)
		if err == nil {
			c.conf.CensorPort = p
		}
	}
	variable = os.Getenv("Logdir")
	if len(variable) != 0 {
		c.conf.Logdir = variable
	}
} */

// загрузка конфигурации из файла или из Переменных окружения
func (c *config) Load(file string) error {
	//если	 file пустой ищем  конфиг  в  директории исполняемого файла
	if len(file) == 0 {
		exe, err := os.Executable()
		if err != nil {
			return err
		}
		file = filepath.Dir(exe) + "/" + tools.FileConfig()
	}
	b, err := os.ReadFile(file)
	if err == nil {
		json.Unmarshal(b, &c.conf)
	} else {
		fmt.Println(err)
	}
	// проверка на корректность
	if c.conf.Port == 0 {
		c.conf.Port = 8080
	}
	if len(c.conf.Dbhost) == 0 {
		c.conf.Dbhost = "localhost"
	}
	if c.conf.Dbport == 0 {
		c.conf.Dbport = 1521
	}

	if len(c.conf.Database) == 0 {
		c.conf.Database = "hydra"
	}
	if len(c.conf.Dbuser) == 0 {
		c.conf.Dbuser = "ais_rpc"
	}
	return nil
}

func (c *config) isStdOut() string {
	if (len(c.conf.Logdir)) == 0 {
		return "stdout"
	}
	return c.conf.Logdir
}

func (c *config) String() string {
	ret := fmt.Sprintf("Port: %d  \nLogDir: %s \ndbhost: %s\ndbport: %d\nDB: %s\ndbuser: %s\n",
		c.conf.Port, c.isStdOut(), c.conf.Dbhost, c.conf.Dbport, c.conf.Database, c.conf.Dbuser)
	return ret
}

// геттер порт сервера цензуры.
func (c *config) Port() int {
	return c.conf.Port
}

// каталог логов
func (c *config) Log() string {
	return c.conf.Logdir
}

func (c *config) DbHost() string {
	return c.conf.Dbhost
}
func (c *config) DbPort() int {
	return c.conf.Dbport
}
func (c *config) DbUser() string {
	return c.conf.Dbuser
}
func (c *config) Db() string {
	return c.conf.Database
}

func (c *config) Passwd() string {
	return c.conf.Dbpasswd
}
