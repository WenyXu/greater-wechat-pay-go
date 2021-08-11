/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/05 19:56
*/

package endpoint

type endpoint struct {
	url    string
	method string
}

func New(url, method string) Endpoint {
	return endpoint{url, method}
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
