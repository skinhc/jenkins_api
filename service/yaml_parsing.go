package service

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// YamlBean yaml文件操作
type YamlBean struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewYamlBean() *YamlBean {
	return &YamlBean{
		Url:      "http://192.168.11.225:8080/",
		Username: "bqj",
		Password: "bqj",
	}
}

//SetYamlConfig 根据文件路径设置文件
func (y *YamlBean) SetYamlConfig(filePath string) error {
	//判断是否是文件夹
	s, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return errors.New("文件不存在")
	}

	if s.IsDir() {
		return errors.New("禁止填入文件夹")
	}

	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(abs, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(all, y)
	if err != nil {
		return err
	}

	return nil

}
