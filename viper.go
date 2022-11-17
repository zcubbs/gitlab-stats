package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var cfgFile string

var Config Configuration

type Configuration struct {
	Gitlab `mapstructure:"gitlab" json:"gitlab"`
}

type Gitlab struct {
	Url          string `mapstructure:"url" json:"url"`
	PrivateToken string `mapstructure:"private_token" json:"private_token"`
	GroupIds     []int  `mapstructure:"group_ids" json:"group_ids"`
	ProjectIds   []int  `mapstructure:"project_ids" json:"project_ids"`
}

var (
	defaults = map[string]interface{}{
		"debug.enabled": false,
	}
	envPrefix   = "GS"
	configName  = "config"
	configType  = "yaml"
	configPaths = []string{
		".",
		fmt.Sprintf("%s/.gitlab-stats-cli", getUserHomePath()),
	}
)

var allowedEnvVarKeys = []string{
	"gitlab.url",
	"gitlab.private_token",
	"gitlab.group_ids",
	"gitlab.project_ids",
	"debug.enabled",
}

// Bootstrap reads in config file and ENV variables if set.
func Bootstrap() {
	err := godotenv.Load(".env")

	if err != nil {
		if viper.GetString("debug.enabled") == "true" {
			log.Println("Error loading .env file")
		}
	}

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		for _, p := range configPaths {
			viper.AddConfigPath(p)
		}
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
		err := viper.ReadInConfig()
		if err != nil {
			if viper.GetString("debug.enabled") == "true" {
				fmt.Println(err)
			}
		}
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(envPrefix)

	for _, key := range allowedEnvVarKeys {
		err := viper.BindEnv(key)
		if err != nil {
			fmt.Println(err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}
}

func getUserHomePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}
