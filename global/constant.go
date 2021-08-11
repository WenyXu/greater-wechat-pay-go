/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 18:06
*/

package global

import (
	"fmt"
	"strings"

	"github.com/WenyXu/greater-wechat-pay-go/endpoint"
)

func MakeEndpointFormString(url string, method string) endpoint.Endpoint {
	return endpoint.New(
		url,
		method,
	)
}

//type MethodGet string

func MakeEndpointFormFunc(getUrl func(...interface{}) string, method string) func(a ...interface{}) endpoint.Endpoint {
	return func(a ...interface{}) endpoint.Endpoint {
		return endpoint.New(
			getUrl(a...),
			method,
		)
	}

}

const (
	RSA2            = "RSA2"
	RSA             = "RSA"
	PKCS1           = "PKCS1"
	PKCS8           = "PKCS8"
	DefaultEndpoint = "https://api.mch.weixin.qq.com"
	Authorization   = "WECHATPAY2-SHA256-RSA2048"
	// 证书接口
	certificates = "/v3/certificates"

	// 商户接口
	payTransactionsIdFormat              = "/v3/pay/transactions/id/%s"
	payTransactionsOutTradeNoFormat      = "/v3/pay/transactions/out-trade-no/%s"
	payTransactionsJsapiFormat           = "/v3/pay/transactions/jsapi"
	payTransactionsH5Format              = "/v3/pay/transactions/h5"
	payTransactionsNativeFormat          = "/v3/pay/transactions/native"
	payTransactionsAppFormat             = "/v3/pay/transactions/app"
	payTransactionsOutTradeNoCloseFormat = "/v3/pay/transactions/out-trade-no/%s/close"

	billTradeBillFormat = "/v3/bill/tradebill"
	billFundFlowBill    = "/v3/bill/fundflowbill"
	// 服务商接口
	// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis
	payPartnerTransactionsJsapiFormat           = "/v3/pay/partner/transactions/jsapi"
	payPartnerTransactionsIdFormat              = "/pay/partner/transactions/id/%s"
	payPartnerTransactionsOutTradeNoCloseFormat = "/pay/partner/transactions/out-trade-no/%s/close"

	// 公共
	refundDomesticRefundQueryFormat = "/v3/refund/domestic/refunds/%s"
	refundDomesticRefundFormat      = "/v3/refund/domestic/refunds"

	// 分账
	profitSharingOrdersFormat                  = "/v3/profitsharing/orders"
	profitSharingOrdersQueryFormat             = "/v3/profitsharing/orders/%s"
	profitSharingReturnOrdersFormat            = "/v3/profitsharing/return-orders"
	profitSharingReturnOrdersQueryFormat       = "/v3/profitsharing/return-orders/%s"
	profitSharingOrdersUnfreeze                = "/v3/profitsharing/orders/unfreeze"
	profitSharingTransactionsAmountQueryFormat = "/v3/profitsharing/transactions/%s/amounts"
	profitSharingMerchantConfigsFormat         = "/profitsharing/merchant-configs/%s"
	profitSharingReceiversAddFormat            = "/v3/profitsharing/receivers/add"
	profitSharingReceiversDeleteFormat         = "/v3/profitsharing/receivers/delete"
)

func makeMakeUrlFromStringFunc(format string) func(a ...interface{}) string {
	return func(a ...interface{}) string {
		str := format
		if strings.Contains(str, "%") {
			if len(a) == 0 {
				str = fmt.Sprintf(format, "")
			} else {
				str = fmt.Sprintf(format, a...)
			}
		}
		if strings.HasSuffix(str, "/") {
			return str[:len(str)-1]
		}
		return str
	}
}

const (
	HeaderAuthorization = "Authorization"
)

var (
	CertificatesEndpoint           = MethodGetEndpoint(certificates)
	PayTransactionsIdEndpoint      = MethodPostEndpoint(payTransactionsIdFormat)
	PayTransactionsOutTradeNo      = makeMakeUrlFromStringFunc(payTransactionsOutTradeNoFormat)
	PayTransactionsJsapi           = makeMakeUrlFromStringFunc(payTransactionsJsapiFormat)
	PayTransactionsOutTradeNoClose = makeMakeUrlFromStringFunc(payTransactionsOutTradeNoCloseFormat)
	RefundDomesticRefund           = makeMakeUrlFromStringFunc(refundDomesticRefundFormat)
	BillTradeBill                  = makeMakeUrlFromStringFunc(billTradeBillFormat)
	BillFundFlowBill               = makeMakeUrlFromStringFunc(billFundFlowBill)
)
