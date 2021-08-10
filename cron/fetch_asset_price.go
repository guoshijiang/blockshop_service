package cron

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

func RealAssetPrice() (err error) {
	logs.Info("start exec RealAssetPrice")
	db := orm.NewOrm()
	err = db.Begin()
	defer func() {
		if err != nil {
			err = db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in RealAssetPrice")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in RealAssetPrice")
		}
	}()

	return nil
}
