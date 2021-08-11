/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/02/26 18:17
*/

package global

type MethodGetEndpoint string
type MethodPostEndpoint string

func (str MethodGetEndpoint) Url() string {
	return string(str)
}

func (str MethodGetEndpoint) Method() string {
	return "GET"
}

func (str MethodPostEndpoint) Url() string {
	return string(str)
}

func (str MethodPostEndpoint) Method() string {
	return "POST"
}

var (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)
