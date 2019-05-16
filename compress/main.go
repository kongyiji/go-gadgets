package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	Default  Default                `yaml:"default"`
	Projects map[string]interface{} `yaml:"projects,omitempty"`
	// Projects map[string]string `yaml:"projects,omitempty"`
}

type Default struct {
	BackupPath string `yaml:"backupPath"`
}

func (c *conf) getConf(file string) *conf {
	ymlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(ymlFile, c)
	fmt.Printf("--- c:\n%v\n\n", *c)

	return c
}

// func (c *conf) getProjectName() projectName string {
// }

func main() {
	var c conf
	c.getConf("config.yml")

	// fmt.Printf("%s", c.Default.BackupPath)

	defaultBackupPath := c.Default.BackupPath

	fmt.Println(defaultBackupPath)

	for k, v := range c.Projects {
		fmt.Printf("key: %v\tvalue: %v\n", k, v)
		// if reflect.TypeOf(v) == 'map' {
		// 	fmt.Println("%v is map", v)
		// }
		if i, ok := v.(map[interface{}]interface{}); ok {
			for kk, vv := range i {
				fmt.Printf("key in %v: %v\tvalue in %v: %v\n", i, kk, i, vv)
			}
		}
		// fmt.Println(reflect.TypeOf(v))
		// i, ok := v.(map[interface{}]interface{})
		// fmt.Println(i, ok)
		// fmt.Println(v.backupPath)
	}
}
