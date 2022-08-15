package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/miaojiaxi2004/go_server/global"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 解组
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
