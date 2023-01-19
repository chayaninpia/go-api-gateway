package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Init() {

	log.Println(`------- Start Init Configs -------`)
	viper.SetConfigName("api-gateway") // ชื่อ config file
	viper.AddConfigPath("./conf")      // ระบุ path ของ config file
	viper.AutomaticEnv()               // อ่าน value จาก ENV variable

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	log.Println(`------- Inited Configs -------`)

}
