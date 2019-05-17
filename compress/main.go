package main

import (
	"fmt"
	"io/ioutil"
	"reflect"

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
	// var projects []map[string]interface{}
	c.getConf("config.yml")

	// fmt.Printf("%s", c.Default.BackupPath)

	defaultBackupPath := c.Default.BackupPath

	fmt.Println(defaultBackupPath)
	// projects = c.Projects
	// fmt.Println(c.Projects)

	for k, v := range c.Projects {
		// project := make(map[string]interface{})
		// project[k] = v

		fmt.Printf("key: %v\tvalue: %v\n", k, v)

		if i, ok := v.(map[interface{}]interface{}); ok {
			// fmt.Println(v)
			for kk, vv := range i {
				fmt.Println("kk: ", reflect.TypeOf(kk))
				fmt.Println("vv: ", reflect.TypeOf(vv))
				// fmt.Printf("-- key in %v: %v\n---- value in %v: %v\n", i, kk, i, vv)
			}
		}

		// fmt.Println(reflect.TypeOf(v))
		// fmt.Println(v.backupPath)
		// fmt.Println(projects)
	}
	// fmt.Println(projects)
}
