package param

import (
	"fmt"
	"github.com/spf13/viper"
)

type server struct {
	Port string
	Name string
}

type AppParam struct {
	Mode   string
	Server server
	Ver    string
	Pwd    string
}

func (p *AppParam) ReadConfig() {

	//default config
	viper.SetDefault("application.mode", "debug")
	viper.SetDefault("application.server.port", "8080")
	viper.SetDefault("application.server.name", "qihuo-app-api")

	//read properties
	p.Mode = viper.GetString("application.mode")
	//注意server的port数字前面有个冒号
	p.Server.Port = ":" + viper.GetString("application.server.port")
	p.Server.Name = viper.GetString("application.server.name")

	p.Ver = viper.GetString("application.ver")
	p.Pwd = viper.GetString("application.pwd")

	fmt.Println("[Init AppParam Over]...", p)

}
