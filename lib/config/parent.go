package config

import (
    "github.com/spf13/viper"
    "log"
)

type ParentConfiguration struct {
    name string
}

func main() {
    viper.SetConfigName("parent")
    viper.AddConfigPath("confs")
    var parentconf parent.Configuration

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading parent config, %s", err)
    }

    err := viper.Unmarshal(&parentconf)
    if err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }

    log.Printf("name %s", parentconf.Parent.name)
}
