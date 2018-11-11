package main

import (
	"bytes"
	"demo/model"
	"demo/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
	"time"
)

const DbDriver = "outer_select:0RdHT0km40uOSdQi@tcp(sh-cdbro-p4af5rft.sql.tencentcdb.com:63890)/db_ex_business?charset=utf8&parseTime=true&loc=Asia%2FShanghai"

func main() {
	db, err := gorm.Open("mysql", DbDriver)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now()
	operate, _ := time.ParseDuration("-1h")
	createdStart := now.Add(4 * operate)
	createdEnd := now.Add(1 * operate)

	pile := &util.Pile{}
	pile.Start()
	var orders []model.Orders
	err = db.Select("app_id, order_id, price, invalid_time, created_at, updated_at, discount_id").
		Where("created_at between ? and ?", createdStart, createdEnd).
		Where("order_state = 0").
		Where("invalid_time between ? and ?", "2018-01-01 00:00:00", createdEnd).
		Limit(250).
		Find(&orders).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	pile.End()

	if len(orders) > 0 {
		// 渲染模板
		tpl := template.Must(template.ParseFiles("./views/email.html"))
		html := &bytes.Buffer{}
		tpl.Execute(html, map[string]interface{}{
			"created_start": createdStart,
			"created_end":   createdEnd,
			"exc_at":        time.Now(),
			"exc_time":      float32(pile.ExcTime) / 1000,
			"orders":        orders,
		})

		// 发送邮件
		err := util.SendEmail("【OCM 数据异常】C端订单未闭环监控", []string{
			"450512013@qq.com",
			"duduchen@xiaoe-tech.com",
		}, string(html.Bytes()))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("mail send success...")
		return
	}

	fmt.Println("cancel all orders ok!")
	return
}
