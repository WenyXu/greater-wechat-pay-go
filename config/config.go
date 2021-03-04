/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 19:22
*/

package config

import "time"

type Config struct {
	Loc                *time.Location
	AppId              string
	PrivateKey         string
	PrivateKeyType     string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	SignType           string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	Format             string
	Version            string
	AppAuthToken       string
	AuthToken          string
	Production         bool
}
