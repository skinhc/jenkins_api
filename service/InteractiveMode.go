package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"jk/bean"
	"strconv"
	"strings"
	"time"
)

//InteractiveMode 交互模式操作
type InteractiveMode struct {
	arguments *bean.CommandLineArguments
	yamlBean  *YamlBean
}

func NewInteractiveMode(arguments *bean.CommandLineArguments, yamlBean *YamlBean) *InteractiveMode {
	return &InteractiveMode{arguments: arguments, yamlBean: yamlBean}
}

//Run 交互模式运行
func (i *InteractiveMode) Run() error {
	jenkins := NewJenkinsService(i.yamlBean.Username, i.yamlBean.Password, i.yamlBean.Url)
	jobs := jenkins.GetJobs()
	fmt.Println(fmt.Sprintf("任务索引\t任务名称"))
	for index, job := range jobs {
		fmt.Println(fmt.Sprintf("%d\t%s", index, job.GetName()))
	}
	scanIndex, err := scanFunc("请输入索引")
	if err != nil {
		return err
	}
	index, err := strconv.Atoi(scanIndex)
	if err != nil {
		return err
	}
	j := jobs[index]
	jenkins.SetJob(j)

	updateBranch, err := scanFunc("是否需要修改分支")
	if err != nil && err.Error() != "unexpected newline" {

		return err
	}
	err = nil
	switch strings.ToLower(updateBranch) {
	case "y":
		fallthrough
	case "yes":
		branchName, err := scanFunc("请输入分支名称")
		if err != nil {
			return err
		}
		err = setBranch(branchName, jenkins)
		if err != nil {
			return err
		}
	default:

	}
	jobId, err := jenkins.BuildJob()
	if err != nil {
		fmt.Println(fmt.Sprintf("构建失败,id:%d,jobName:%s,error:%s", jobId, j.GetName(), err.Error()))
		return err
	}
	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	fmt.Println(fmt.Sprintf("%s:%s", nowTime, j.Raw.URL))

	return nil

}

func setBranch(branchName string, jenkins *JenkinsService) error {
	configuration := jenkins.GetJobConfiguration()
	project := bean.NewProject()
	err := xml.Unmarshal(bytes.NewBufferString(configuration).Bytes(), project)
	if err != nil {
		return err
	}
	project.Scm.Branches.HudsonPluginsGitBranchSpec.Name.Text = branchName
	newConfig, err := xml.Marshal(project)
	if err != nil {
		return err
	}
	err = jenkins.UpdateConfig(string(newConfig))
	if err != nil {
		return err
	}
	return nil
}
func scanFunc(message string) (string, error) {
	fmt.Print(message + ":")
	var scanIndex string
	_, err := fmt.Scanln(&scanIndex)
	if err != nil {
		return "", err
	}
	return scanIndex, nil

}
