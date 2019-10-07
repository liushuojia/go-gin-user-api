package conf

import (
	"flag"
	"github.com/Unknwon/goconfig"
	"path/filepath"
)

var configFile *goconfig.ConfigFile
var ConfigApi map[string] string
var ConfigMysql map[string] string
var ConfigRedis map[string] string
var ConfigAdminRedis map[string] string
var ConfigSmsRedis map[string] string
var ConfigEmailRedis map[string] string

var (
	confFile = flag.String("c", "", "set configuration `file`")
)

func init() {

	var err error
	flag.Parse()
	if *confFile == "" {
		dir, _ := filepath.Abs(`.`)
		*confFile =  dir + "/app.conf"
	}

	configFile, err = goconfig.LoadConfigFile( *confFile )

	if err != nil{
		panic("读取配置文件出现错误")
		return
	}

	ConfigApi, err = configFile.GetSection("api" )
	ConfigMysql, err = configFile.GetSection("mysql" )
	ConfigRedis, err = configFile.GetSection("redis" )

	ConfigAdminRedis, err = configFile.GetSection("adminRedis" )
	ConfigSmsRedis, err = configFile.GetSection("smsRedis" )
	ConfigEmailRedis, err = configFile.GetSection("emailRedis" )

	return
}
