/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 19:22
*/

package config

import (
	"crypto/rsa"
	"time"

	"github.com/WenyXu/greater-wechat-pay-go/global"
)

type Config struct {
	Loc          *time.Location
	AppId        string
	MchID        string
	ApiV3Key     string
	PrivateKey   *rsa.PrivateKey
	PublicCertSn string
	SignType     string
	ReturnUrl    string
	NotifyUrl    string
	Charset      string
	Format       string
	Version      string
	AppAuthToken string
	AuthToken    string
	Production   bool
}

func (c Config) URL() string {
	return global.DefaultEndpoint
}
