package config

import (
    "os"
    "io/ioutil"
	"gopkg.in/yaml.v2"
    "log"
    "fmt"
)

type Child struct {
	Name  string  `yaml:name`
	Bin   string  `yaml:bin`
	Tasks []*Task `yaml:"child"`
}

func (c Child) Keys() []string {
    keys := []string()
    for _; t := range c.Tasks {
        keys = append(keys, t.Name)
    }
    return keys
}

func (c Child) Fetch(key string) *Task {
    for _, t := range c.Tasks {
        if t.Name == key {
            return t
        }
    }
    return nil
}

func LoadChildConfig(c Child) {
    yamlFile, err := ioutil.Readfile("confs/child.yaml")

    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &c)
    fmt.Printf("%+v\n", c)
    return c
}
