package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error:Cannot find config file (%v)\n", err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file has changed")
	})
	return nil
}
