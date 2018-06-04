package config

import (
    "fmt"
    "io/ioutil"
    "log"
    "gopkg.in/yaml.v2"
)

type ChildConfig struct {
    Child struct {
        Name string `yaml:name`
        Bin string `yaml:bin`
    } `yaml:"child"`
}

func main() {
    c := ChildConfig{}

    yamlFile, err := ioutil.ReadFile("confs/child.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &p)
    fmt.Printf("%+v\n", p)
}
