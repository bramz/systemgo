package config

import (
    "fmt"
    "io/ioutil"
    "log"
    "gopkg.in/yaml.v2"
)

type ParentConfig struct {
    Parent struct {
        Name string `yaml:name`
        Root string `yaml:root`
    } `yaml:"parent"`
}

func main() {
    p := ParentConfig{}

    yamlFile, err := ioutil.ReadFile("confs/parent.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &p)
    fmt.Printf("%+v\n", p)
}
