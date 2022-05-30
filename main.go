package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"jk/bean"
	"jk/service"
	"os"
	"strings"
)

var (
	jenkinsClient *gojenkins.Jenkins
)

func init() {

	yamlBean := service.NewYamlBean()

	jenkinsClient = gojenkins.CreateJenkins(nil, yamlBean.Url, yamlBean.Username, yamlBean.Password)

}
func main() {

	GetJob("admin-shopping-mall", "parent")
	yamlBean := service.NewYamlBean()
	arguments := bean.NewCommandLineArguments()
	if *arguments.FilePath != "" {
		err := yamlBean.SetYamlConfig(*arguments.FilePath)
		if err != nil {
			fmt.Println("文件异常 err:" + err.Error())
		}
	}
	//判断模式
	switch *arguments.Model {
	case 0:
		if *arguments.JobName == "" {
			fmt.Println("无交互模式请添加job名称")
			os.Exit(1)
		}
		fmt.Println("无交互模式................................")
		notInteractiveMode := service.NewNotInteractiveMode(arguments, yamlBean)

		err := notInteractiveMode.Run()
		if err != nil {
			panic(err)
		}

	case 1:
		fmt.Println("交互模式................................")
		interactiveMode := service.NewInteractiveMode(arguments, yamlBean)

		err := interactiveMode.Run()
		if err != nil {
			panic(err)
		}

	default:
		fmt.Println("请选择交互模式................................")
		os.Exit(1)

	}
}
func updateConfig(ctx context.Context, config string) {

}
func GetJob(id string, parentIDs ...string) {
	s := "/job/" + strings.Join(append(parentIDs, id), "/job/")
	fmt.Println(s)
}
