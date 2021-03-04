/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 18:55
*/

package entity

import "time"

// /v3/pay/transactions/app
type TransactionsAppResponse struct {
	PrepayID string `json:"prepay_id"`
}

// /v3/pay/transactions/id/{transaction_id}
// /v3/pay/transactions/out-trade-no/{out_trade_no}
type TransactionsResponse struct {
	TransactionID string `json:"transaction_id"`
	Amount        *struct {
		PayerTotal    int    `json:"payer_total"`
		Total         int    `json:"total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
	MchId           string `json:"mchid"`
	TradeState      string `json:"trade_state"`
	BankType        string `json:"bank_type"`
	PromotionDetail *[]struct {
		Amount              int    `json:"amount"`
		WechatpayContribute int    `json:"wechatpay_contribute"`
		CouponID            string `json:"coupon_id"`
		Scope               string `json:"scope"`
		MerchantContribute  int    `json:"merchant_contribute"`
		Name                string `json:"name"`
		OtherContribute     int    `json:"other_contribute"`
		Currency            string `json:"currency"`
		StockID             string `json:"stock_id"`
		GoodsDetail         *[]struct {
			GoodsRemark    string `json:"goods_remark"`
			Quantity       int    `json:"quantity"`
			DiscountAmount int    `json:"discount_amount"`
			GoodsID        string `json:"goods_id"`
			UnitPrice      int    `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
	SuccessTime time.Time `json:"success_time"`
	Payer       struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	OutTradeNo     string `json:"out_trade_no"`
	AppId          string `json:"appid"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeType      string `json:"trade_type"`
	Attach         string `json:"attach"`
	SceneInfo      *struct {
		DeviceID string `json:"device_id"`
	} `json:"scene_info"`
}
