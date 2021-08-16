package cron

import (
	"blockshop/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MarketDataItem struct {
	CurrentPrice    float64  `json:"current_price"`
	CurrentPriceUsd float64 `json:"current_price_usd"`
	Code  string   `json:"code"`
	Name  string   `json:"name"`
}

type MarketData struct {
	Maxpage int64 `json:"maxpage"`
	Currpage int64 `json:"currpage"`
	Code int64 `json:"code"`
	Msg string `json:"msg"`
	Data []MarketDataItem `json:"data"`
}

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
	req_url := "https://dncapi.bqiapp.com/api/coin/web-coinrank?page=1&type=-1&pagesize=%s&webp=1"
	resp, _ := http.Get(req_url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var ret_data MarketData
	if err := json.Unmarshal(body, &ret_data); err != nil {
		return err
	}
	for _, value := range ret_data.Data {
		if value.Name == "USDT" || value.Name == "BTC"{
			var asset models.Asset
			err := db.QueryTable(models.Asset{}).Filter("name", value.Name).One(&asset)
			if err == nil {
				usd_price_str := strconv.FormatFloat(value.CurrentPriceUsd, 'f', -1, 64)
				cny_price_str := strconv.FormatFloat(value.CurrentPrice,'f', -1, 64)
				asset.UsdPrice = usd_price_str
				asset.CnyPrice = cny_price_str
				err := asset.Update()
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			continue
		}
	}
	if err != nil {
		return err
	}
	return nil
}
