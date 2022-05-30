package service

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"net/http"
	"strings"
)

var client *gojenkins.Jenkins

type JenkinsService struct {
	ctx context.Context
	job *gojenkins.Job
}

func NewJenkinsService(username string, password string, url string) *JenkinsService {

	ctx := context.Background()
	j := &JenkinsService{ctx: ctx}

	client = gojenkins.CreateJenkins(http.DefaultClient, url, username, password)
	_, err := client.Init(ctx)
	if err != nil {
		panic("jenkins init error: " + err.Error())
	}

	return j
}

//GetJobs 获取所有的jobs
func (j *JenkinsService) GetJobs() []*gojenkins.Job {
	fmt.Println("正在读取job列表>>>>>>>>>>>>>>>>")
	jobs, err := client.GetAllJobs(j.ctx)
	if err != nil {
		panic("获取任务列表时发生异常:" + err.Error())
	}
	return jobs
}

//SetJob 设置要操作的指定job
func (j *JenkinsService) SetJob(job *gojenkins.Job) {
	j.job = job
	j.checkJob()

}

//GetJobConfiguration 获取某个指定job的配置
func (j *JenkinsService) GetJobConfiguration() string {
	fmt.Println(fmt.Sprintf("正在获取 %s 的配置项>>>>>>>>>>", j.job.GetName()))
	j.checkJob()
	config, err := j.job.GetConfig(j.ctx)
	if err != nil {
		panic("获取任务配置时异常:" + err.Error())
	}

	return strings.Replace(config, "1.1", "1.0", 1)
}

//UpdateConfig 更新某个job的配置
func (j *JenkinsService) UpdateConfig(config string) error {
	fmt.Println(fmt.Sprintf("正在更新 %s 的配置项>>>>>>>>>>", j.job.GetName()))

	j.checkJob()
	err := j.job.UpdateConfig(j.ctx, config)
	return err
}
func (j *JenkinsService) BuildJob() (int64, error) {
	fmt.Println(fmt.Sprintf("正在构建 %s 的配置项>>>>>>>>>>", j.job.GetName()))

	j.checkJob()
	//name := job.GetName()
	//client.GetJob()
	jobName := j.job.GetName()
	jobId, err := client.BuildJob(j.ctx, jobName, nil)
	return jobId, err
}

//checkJob 检查struct中的job指针
func (j *JenkinsService) checkJob() {
	if j.job == nil {
		panic("Job not found")
	}

}
