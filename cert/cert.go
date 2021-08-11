/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/05 19:23
*/

package cert

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/WenyXu/better-alipay-go/global"
	"github.com/WenyXu/greater-wechat-pay-go/sign"
)

// LoadCertSN load root cert sn form path or bytes
func LoadCertSN(certPathOrData interface{}) (sn string, err error) {
	var certData []byte
	switch certPathOrData.(type) {
	case string:
		certData, err = ioutil.ReadFile(certPathOrData.(string))
	case []byte:
		certData = certPathOrData.([]byte)
	}
	if err != nil {
		return sn, err
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return sn, err
		}
		sn = strings.ToUpper(cert.SerialNumber.Text(16))
	}
	if sn == "" {
		return "", errors.New("failed to load cert sn,check the cert path or data")
	}
	return sn, nil
}

func LoadPrivateCertFormBytes(privateKeyType string, input []byte) sign.LoadPrivateKeyFunc {
	return func() (privateKey *rsa.PrivateKey, err error) {
		var (
			block *pem.Block
		)
		if block, _ = pem.Decode(input); block == nil {
			err = errors.New("pem.Decode：privateKey decode error")
			return
		}
		switch privateKeyType {
		case global.PKCS1:
			if privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
				return
			}
		default:
			pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			pk8, ok := pkcs8Key.(*rsa.PrivateKey)
			if !ok {
				err = errors.New("parse PKCS8 key error")
				return nil, err
			}
			privateKey = pk8
		}
		return
	}
}

func LoadPrivateCertFormString(privateKeyType string, input string) sign.LoadPrivateKeyFunc {
	return LoadPrivateCertFormBytes(privateKeyType, []byte(input))
}

func LoadPrivateCertFormPath(privateKeyType string, input string) sign.LoadPrivateKeyFunc {
	return func() (*rsa.PrivateKey, error) {
		bytes, err := ioutil.ReadFile(input)
		if err != nil {
			err = fmt.Errorf("公钥文件读取失败: %w", err)
			return nil, err
		}
		return LoadPrivateCertFormBytes(privateKeyType, bytes)()
	}
}

// LoadPublicCertFormBytes load public cert form byte
func LoadPublicCertFormBytes(input []byte) (publicKey *rsa.PublicKey, err error) {
	var (
		block  *pem.Block
		pubKey *x509.Certificate
		ok     bool
	)
	if block, _ = pem.Decode(input); block == nil {
		err = errors.New("公钥Decode错误")
		return
	}
	if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
		err = fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
		return
	}
	if _, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
		err = errors.New("public key 类型断言错误")
		return
	}
	return pubKey.PublicKey.(*rsa.PublicKey), nil
}

// PublicCertFormString load public key form string
func PublicCertFormString(input string) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return LoadPublicCertFormBytes([]byte(input))
	}
}

// PublicCertFormBytes load public key form byte
func PublicCertFormBytes(input []byte) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return LoadPublicCertFormBytes(input)
	}
}

// PublicCertFormPath load public key form path
func PublicCertFormPath(input string) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		bytes, err := ioutil.ReadFile(input)
		if err != nil {
			err = fmt.Errorf("公钥文件读取失败: %w", err)
			return nil, err
		}
		return LoadPublicCertFormBytes(bytes)
	}
}
