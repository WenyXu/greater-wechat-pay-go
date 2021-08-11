/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/05 19:24
*/

package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

type LoadPublicKeyFunc func() (publicKey *rsa.PublicKey, err error)
type LoadPrivateKeyFunc func() (privateKey *rsa.PrivateKey, err error)

func SignRequest(str string, privateKey *rsa.PrivateKey) (sign string, err error) {
	var (
		h         hash.Hash
		encrypted []byte
	)
	h = sha256.New()
	//body, _ := ioutil.ReadAll(req.Body)
	//req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// req.Method
	// req.URL.RawPath
	// time.Now().Unix()
	// strconv.FormatUint(rand.Uint64(), 16)
	// body
	h.Write([]byte(str))

	if encrypted, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encrypted)
	return
}
