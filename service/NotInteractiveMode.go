package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/bndr/gojenkins"
	"jk/bean"
	"time"
)

//NotInteractiveMode 非交互模式操作
type NotInteractiveMode struct {
	arguments *bean.CommandLineArguments
	yamlBean  *YamlBean
}

func NewNotInteractiveMode(arguments *bean.CommandLineArguments, yamlBean *YamlBean) *NotInteractiveMode {
	return &NotInteractiveMode{arguments: arguments, yamlBean: yamlBean}
}

//Run 无交互模式运行
func (i *NotInteractiveMode) Run() error {
	jenkins := NewJenkinsService(i.yamlBean.Username, i.yamlBean.Password, i.yamlBean.Url)

	jobs := jenkins.GetJobs()
	var j *gojenkins.Job
	for _, job := range jobs {
		if job.GetName() == *(i.arguments.JobName) {
			j = job

			break
		}
	}
	jenkins.SetJob(j)

	//判断是否要修改配置
	if *(i.arguments.BranchName) != "" {
		configuration := jenkins.GetJobConfiguration()

		project := bean.NewProject()
		err := xml.Unmarshal(bytes.NewBufferString(configuration).Bytes(), project)
		if err != nil {
			return err
		}
		project.Scm.Branches.HudsonPluginsGitBranchSpec.Name.Text = *(i.arguments.BranchName)
		newConfig, err := xml.Marshal(project)
		if err != nil {
			return err
		}
		err = jenkins.UpdateConfig(string(newConfig))
		if err != nil {
			return err
		}
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
