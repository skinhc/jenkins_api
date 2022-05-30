package bean

import "flag"

//CommandLineArguments 命令行参数
var c *CommandLineArguments

type CommandLineArguments struct {
	//JobName job名称
	JobName *string
	//FilePath 文件路径
	FilePath *string
	//Model 模式 0 无交互 1 交互模式
	Model *uint
	//BranchName 分支名称
	BranchName  *string
	ProjectName *string
}

func NewCommandLineArguments() *CommandLineArguments {
	if c == nil {
		c = &CommandLineArguments{}
		c.Model = flag.Uint("m", 1, "模式 0 无交互 1 交互模式 默认交互模式")
		c.JobName = flag.String("n", "", "job名称")
		c.FilePath = flag.String("f", "", "文件路径")
		c.BranchName = flag.String("b", "", "分支名称")

		flag.Parse()
	}

	return c
}
