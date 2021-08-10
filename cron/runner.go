package cron

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	RealTimeExecution  = time.Second * 1
	WithdrawPeriod     = time.Second * 60
)

func Run() {
	if beego.BConfig.RunMode == "dev" {
		RealTimeExecution = time.Second * 1
		WithdrawPeriod = time.Second * 10
	}
	go func() {
		for {
			select {
			case <-time.Tick(RealTimeExecution):
				err := RealAssetPrice()
				if err != nil {
					logs.Error("run get asset price error %v", err)
				} else {
					logs.Info("run get asset price success.")
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case now := <-time.Tick(time.Minute):
				if now.Hour() == 0 && now.Minute() == 1 {
					fmt.Println("定点执行任务")
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-time.Tick(WithdrawPeriod):
				fmt.Println("充值提现任务")
			}
		}
	}()
}

