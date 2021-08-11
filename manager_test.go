/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/08 14:28
*/

package wechatpay

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bmizerany/assert"

	"github.com/WenyXu/greater-wechat-pay-go/storage"

	"github.com/joho/godotenv"

	"github.com/WenyXu/greater-wechat-pay-go/options"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(m.Run())
}

func TestNewStorageManager(t *testing.T) {
	s := storage.NewMemoryStorage()
	opt := options.DefaultOptions(
		options.MchID("1587998081"),
		options.APIV3Key(os.Getenv("PRIVATE_KEY")),
		options.PrivateKeyFromPath(os.Getenv("PRIVATE_KEY_PATH")),
		options.ClientCertSnFromPath(os.Getenv("PUBLIC_CERT_PATH")),
	)
	fmt.Println(opt)
	m := NewStorageManager(s, opt.Config)
	resp, err := m.Request()
	assert.Equal(t, nil, err)
	buf, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, nil, err)
	fmt.Println(string(buf))
}
