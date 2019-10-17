package task

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func GetTask() *toolbox.Task {
	return toolbox.NewTask(GetTaskName(), "0/1 * * * * *", task)
}

func GetTaskName() string {
	return "testTask"
}

func task() error {
	logs.Info("the task is running")
	return nil
}
