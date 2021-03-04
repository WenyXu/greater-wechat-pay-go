/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 19:03
*/

package entity

// /v3/bill/tradebill
type TradeBillResponse struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
}

// /v3/bill/fundflowbill
type FundFlowBillResponse struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
}
