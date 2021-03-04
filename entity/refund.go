/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 19:02
*/

package entity

import "time"

// /v3/refund/domestic/refunds/{out_refund_no}
type DomesticRefundsResponse struct {
	RefundID            string    `json:"refund_id"`
	OutRefundNo         string    `json:"out_refund_no"`
	TransactionID       string    `json:"transaction_id"`
	OutTradeNo          string    `json:"out_trade_no"`
	Channel             string    `json:"channel"`
	UserReceivedAccount string    `json:"user_received_account"`
	SuccessTime         time.Time `json:"success_time"`
	CreateTime          time.Time `json:"create_time"`
	Status              string    `json:"status"`
	FundsAccount        string    `json:"funds_account"`
	Amount              struct {
		Total            int    `json:"total"`
		Refund           int    `json:"refund"`
		PayerTotal       int    `json:"payer_total"`
		PayerRefund      int    `json:"payer_refund"`
		SettlementRefund int    `json:"settlement_refund"`
		SettlementTotal  int    `json:"settlement_total"`
		DiscountRefund   int    `json:"discount_refund"`
		Currency         string `json:"currency"`
	} `json:"amount"`
	PromotionDetail *[]struct {
		PromotionID  string `json:"promotion_id"`
		Scope        string `json:"scope"`
		Type         string `json:"type"`
		Amount       int    `json:"amount"`
		RefundAmount int    `json:"refund_amount"`
		GoodsDetail  *[]struct {
			MerchantGoodsID  string `json:"merchant_goods_id"`
			WechatpayGoodsID string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			UnitPrice        int    `json:"unit_price"`
			RefundAmount     int    `json:"refund_amount"`
			RefundQuantity   int    `json:"refund_quantity"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
}
