package health

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	if dc.isConnected() {
		return nil
	} else {
		return errors.New("can`t connect database")
	}
}

func (dc *DatabaseCheck) isConnected() bool {
	_, err := orm.NewOrm().Raw("select 1").Exec()
	if err != nil {
		return false
	}
	return true
}
