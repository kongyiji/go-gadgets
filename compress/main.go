package main

import (
	"os"
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	Default  Default                `yaml:"default"`
	Projects map[string]interface{} `yaml:"projects,omitempty"`
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

func compressToTarGz(filename string) {
	// Create a file to write
	f, err := os.Create(filename)
	// var buf bytes.Buffer
	

	// Create a new gzip writer
	gz := gzip.NewWriter(f)

	// Create a new tar writer
	tw := tar.NewWriter(gz)
	
	def tw.Close()
	
	
}

func main() {
	var c conf
	// var projects []map[string]interface{}
	c.getConf("config.yml")

	// fmt.Printf("%s", c.Default.BackupPath)

	defaultBackupPath := c.Default.BackupPath

	fmt.Println(defaultBackupPath)

	fmt.Println("project: ", c.Projects)
	for k, v := range c.Projects {
		// for k, v := range projects {
		// project := make(map[string]interface{})
		// project[k] = v

		fmt.Printf("key: %v\tvalue: %v\n", k, v)

		if i, ok := v.(map[interface{}]interface{}); ok {
			if _, ok := i["backupPath"]; !ok {
				fmt.Println("backup not ok! Add default")
				i["backupPath"] = defaultBackupPath
			}
			for kk, vv := range i {
				// fmt.Println("kk: ", reflect.TypeOf(kk))
				// fmt.Println("vv: ", reflect.TypeOf(vv))
				fmt.Printf("-- key in %v: %v\n---- value in %v: %v\n",
					i, kk, i, vv)
			}
		}

		// fmt.Println(reflect.TypeOf(v))
		// fmt.Println(v.backupPath)
		// fmt.Println(projects)
	}
	// fmt.Println(projects)
}
