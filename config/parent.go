package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Parent ParentConfig
    Child ChildConfig
}

func main() {
    viper.SetConfigName("conf")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading configuration file, %v", err)
    }
    var config conf.Config

    err := viper.Unmarshal(&config)
    if err != nil {
        log.Fatalf("Unable to form a struct, %v", err)
    }
}

