package cron

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)
/*
 tickers_conf = [
            ['btc-price', 'api/spot/v3/instruments/BTC-USDT/ticker', 'last'],

        ]
 */

const BaseUrl = "https://www.ouyi.cc/api/futures/v3/rate"

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
	btc_price := "api/spot/v3/instruments/BTC-USDT/ticker"
	rate_price := "api/futures/v3/rate"
	fmt.Println(btc_price, rate_price)
	return nil
}
