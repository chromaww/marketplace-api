package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Config struct {
	Version      string
	Port         int
	DebugMode    bool
	LogFilePath  string
	DBConnection *gorm.DB
}

func Load(environment string) *Config {
	cfg := new(Config)

	var configFile *viper.Viper = viper.New()
	configFile.SetConfigType("json")
	configFile.SetConfigName(environment)
	configFile.AddConfigPath("config/env")

	var err error = configFile.ReadInConfig()
	if err != nil {
		log.Print("Error while loading configuration file!")
		log.Fatal(err)
	}

	// var dbInstance Database = Database{
	// 	Driver:   configFile.GetString("database.driver"),
	// 	Host:     configFile.GetString("database.host"),
	// 	Port:     configFile.GetString("database.port"),
	// 	Database: configFile.GetString("database.name"),
	// 	Username: configFile.GetString("database.username"),
	// 	Password: configFile.GetString("database.password"),
	// 	SslMode:  configFile.GetString("database.sslmode"),
	// }

	cfg.Version = configFile.GetString("application.version")
	cfg.Port = configFile.GetInt("server.port")
	cfg.DebugMode = configFile.GetBool("application.debug_mode")
	cfg.LogFilePath = configFile.GetString("application.log_file_path")
	// cfg.DBConnection = dbInstance.Connect()

	return cfg
}
