package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func (Orders) TableName() string {
	return "t_orders"
}

type Orders struct {
	// '应用Id',
	AppId string `gorm:"primary_key;type:varchar(64);not null;index:index_app_id_transaction_id"`
	// '内容方应用Id',
	ContentAppId string `gorm:"type:varchar(64)"`
	// '订单Id',
	OrderId string `gorm:"primary_key;type:varchar(64);not null;unique;unique_index"`
	// '用户id',
	UserId string `gorm:"type:varchar(64);not null"`
	// '付费类型：2-单笔、3-付费产品包、4-团购、5-单笔的购买赠送、6-产品包的购买赠送、7-问答提问、8-问答偷听、9-购买会员、10-会员的购买赠送、11-付费活动报名、12-打赏类型',
	PaymentType int `gorm:"not null"`
	// '资源类型：0-无（不通过资源的购买入口，如团购）、1-图文、2-音频、3-视频、4-直播、5-活动报名(若payment_type=14时，5-会员)、6-专栏、7-社群、8-大专栏、20-电子书、21-实物商品',
	ResourceType int `gorm:"not null"`
	// '资源物品标识，单笔时为买的资源的id, 付费产品包时为通过哪个资源id而产生的点击购买行为，团购时为团购配置表的id，问答提问和偷听均为问题的id',
	ResourceId string `gorm:"type:varchar(64)"`
	// 'payment_type为2时-NULL, payment_type为3时-绑定的付费产品包id，payment_type为4时留空',
	ProductId string `gorm:"type:varchar(64)"`
	// '购买数量，默认为1，团购时按实际情况填写',
	Count int `gorm:"type:varchar(64);default:1"`
	// '渠道id',
	ChannelId string `gorm:"type:varchar(64)"`
	// '渠道来源',
	ChannelInfo string `gorm:"type:varchar(64)"`
	// '若是来自分享，填写分享人的id',
	ShareUserId string `gorm:"type:varchar(64)"`
	// '分享类型 0-音频分享 1-日签分享 2-专栏分享 4-邀请卡分享 5-推广分销分享 6- 推广分销平台',
	ShareType int
	// '购买的资源名',
	PurchaseName string `gorm:"type:varchar(64)"`
	// '资源配图url',
	ImgUrl string `gorm:"type:varchar(256)"`
	// '用户优惠券id',
	CuId string `gorm:"type:varchar(64);not null"`
	// '优惠价格:分',
	CouPrice int `gorm:"not null"`
	// '折扣id',
	DiscountId string `gorm:"type:varchar(64)"`
	// '折扣减免的价格：分',
	DiscountPrice int
	// '价格（分）',
	Price int
	// '订单状态：0-未支付 1-支付成功 2-支付失败 3-退款 4-预定(如问答的提问，未使用) 5(未使用) 6-订单过期自动取消 7-手动取消订单',
	OrderState int
	// '商品类型(0-虚拟商品 1-实物商品)',
	GoodsType int
	// '发货状态(0-禁止发货 1-待发货 2-已发货 3-已收货) (目前仅用于实物商品)',
	ShipState int
	// '成功支付的外部订单号(out_orders)',
	OutOrderId string `gorm:"type:varchar(64)"`
	// '微信支付交易单号',
	TransactionId string `gorm:"type:varchar(64);index:index_app_id_transaction_id"`
	// '数据来源 0-小程序 1-公众号 10-开放平台 11-PC通用版 12-App',
	WxAppType int
	// '有效期（秒）null则不限时间',
	Period sql.NullInt64
	// '是否是来自代收 0-否 1-是',
	UseCollection int
	// '结算状态（0-未结算；1-结算中；2-结算完成）',
	SettleStatus int
	// '是否为分销订单,默认0-普通 1-分销订单',
	DistributeType int
	// '分销提成（分）',
	DistributePrice int
	// '分销比例(百分点1-100)',
	DistributePercent int
	// '上级分销用户id',
	SuperiorDistributeUserId string `gorm:"type:varchar(64)"`
	// '上级分销提成（分）',
	SuperiorDistributePrice int
	// '上级分销比例(百分点1-100)',
	SuperiorDistributePercent int
	// '订单关联的id 分销时为t_distribute_detail表的id；若payment_type为问答提问，则存储答主id,若payment_type为13,14时存小团id；活动购票则存票种id',
	RelatedId string `gorm:"type:varchar(64)"`
	// '是否是续费订单：默认0-普通订单、1-会员续费、2-过期后购买、3-删除后购买',
	IsRenew int
	// '取消订单超时时间',
	InvalidTime time.Time
	// '支付时间',
	PayTime mysql.NullTime
	// '结算时间',
	SettleTime mysql.NullTime
	// '退款时间',
	RefundTime mysql.NullTime
	// '退款金额',
	RefundMoney int
	// '0 B端客户订单  1 小鹅通精选订单 2 内容市场',
	Source int
	// '设备信息',
	Agent string `gorm:"type:varchar(512)"`
	// '创建时间',
	CreatedAt time.Time
	// '更新时间，有修改自动更新',
	UpdatedAt time.Time
	// 问题发货状态
	QueCheckState int
}
