/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 18:06
*/

package global

import (
	"fmt"
	"strings"
)

type endpoint struct {
	url    string
	method string
}

func (e endpoint) Url() string {
	return e.url
}

func (e endpoint) Method() string {
	return e.method
}

type Endpoint interface {
	Url() string
	Method() string
}

func MakeEndpointFormString(url string, method string) Endpoint {
	return endpoint{
		url:    url,
		method: method,
	}
}

func MakeEndpointFormFunc(getUrl func() string, method string) Endpoint {
	return endpoint{
		url:    getUrl(),
		method: method,
	}
}

const (
	payTransactionsIdFormat              = "/v3/pay/transactions/id/%s"
	payTransactionsOutTradeNoFormat      = "/v3/pay/transactions/out-trade-no/%s"
	payTransactionsJsapiFormat           = "/v3/pay/transactions/jsapi"
	payTransactionsH5Format              = "/v3/pay/transactions/h5"
	payTransactionsNativeFormat          = "/v3/pay/transactions/native"
	payTransactionsAppFormat             = "/v3/pay/transactions/app"
	payTransactionsOutTradeNoCloseFormat = "/v3/pay/transactions/out-trade-no/%s/close"
	refundDomesticRefundFormat           = "/v3/refund/domestic/refunds/%s"
	billTradeBillFormat                  = "/v3/bill/tradebill"
	billFundFlowBill                     = "/v3/bill/fundflowbill"
)

func makeMakeUrlFromStringFunc(format string) func(a ...interface{}) string {
	return func(a ...interface{}) string {
		str := format
		if strings.Contains(str, "%") {
			if len(a) == 0 {
				str = fmt.Sprintf(format, "")
			} else {
				str = fmt.Sprintf(format, a)
			}
		}
		if strings.HasSuffix(str, "/") {
			return str[:len(str)-1]
		}
		return str
	}
}

var (
	PayTransactionsId              = makeMakeUrlFromStringFunc(payTransactionsIdFormat)
	PayTransactionsOutTradeNo      = makeMakeUrlFromStringFunc(payTransactionsOutTradeNoFormat)
	PayTransactionsJsapi           = makeMakeUrlFromStringFunc(payTransactionsJsapiFormat)
	PayTransactionsOutTradeNoClose = makeMakeUrlFromStringFunc(payTransactionsOutTradeNoCloseFormat)
	RefundDomesticRefund           = makeMakeUrlFromStringFunc(refundDomesticRefundFormat)
	BillTradeBill                  = makeMakeUrlFromStringFunc(billTradeBillFormat)
	BillFundFlowBill               = makeMakeUrlFromStringFunc(billFundFlowBill)
)
