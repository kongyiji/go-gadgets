package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("test.tar.gz")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	gz := gzip.NewWriter(f)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	fr, err := os.Open("测试中文文件名.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fr.Close()

	_, err = io.Copy(tw, fr)
	if err != nil {
		fmt.Println(err)
	}
}
