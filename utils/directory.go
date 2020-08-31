package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (error) {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path,os.ModeDir);err != nil{
			return err
		}
		return nil
	}
	return err
}

// @title    createDir
// @description   批量创建文件夹
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		err := PathExists(v)
		if err != nil {
			err = os.MkdirAll(v, os.ModeDir)
			if err != nil {
				return err
			}
		}else {
			return nil
		}
	}
	return err
}

// @title    createDir
// @description   获取运行地址
// @auth                     （2020/04/05  20:22）
// @param     path            string
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("create GetCurrentDirectory ", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}