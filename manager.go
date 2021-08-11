/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/06 20:27
*/

package wechatpay

import (
	"context"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/WenyXu/greater-wechat-pay-go/config"
	"github.com/WenyXu/greater-wechat-pay-go/global"
	"github.com/WenyXu/greater-wechat-pay-go/options"
	"github.com/WenyXu/greater-wechat-pay-go/storage"
)

type CertInfo struct {
	EffectiveTime time.Time
	ExpireTime    time.Time
	PublicKey     *x509.Certificate
}

type Manager interface {
	GetLatestEffectiveCert() (CertInfo, bool)
	GetCertBySn(sn string) (CertInfo, bool)
	StoreCert(sn string, value CertInfo) error
	Request() (*http.Response, error)
}

type manager struct {
	storage.Storage
	config config.Config
}

var (
	latestEffectiveCertKey = "LatestEffective"
	makeReq                = options.NewDefaultMakeReqFunc
	transport              = options.DefaultTransport
)

func (m *manager) Request() (*http.Response, error) {
	ep := global.CertificatesEndpoint
	req, err := makeReq(context.Background(), m.config.URL()+ep.Url(), ep.Method(), nil, m.config)
	if err != nil {
		return nil, err
	}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *manager) GetLatestEffectiveCert() (CertInfo, bool) {
	v, ok := m.Get(latestEffectiveCertKey)
	return v.(CertInfo), ok
}

func (m *manager) tryStoreLatestEffectiveCert(value CertInfo) error {
	c, ok := m.GetLatestEffectiveCert()
	if ok {
		if c.EffectiveTime.Sub(value.EffectiveTime) > 0 {
			return nil
		}
	}
	return m.Set(latestEffectiveCertKey, value)
}

func (m *manager) GetCertBySn(sn string) (CertInfo, bool) {
	v, ok := m.Get(sn)
	return v.(CertInfo), ok
}

func (m *manager) StoreCert(sn string, value CertInfo) error {
	err := m.tryStoreLatestEffectiveCert(value)
	if err != nil {
		return err
	}
	return m.Set(sn, value)
}

func NewStorageManager(storage storage.Storage, config config.Config) Manager {
	return &manager{storage, config}
}
