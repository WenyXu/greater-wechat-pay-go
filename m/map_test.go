/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/05 19:06
*/

package m

import (
	"encoding/json"
	"fmt"
	"github.com/WenyXu/better-alipay-go/m"
	"testing"
)

func TestMarshal(t *testing.T) {
	v:=m.NewMap(func(sub m.M) {
		sub.Set("test", func(sub m.M) {
			sub.Set("test",1)
		})
	})
	res,_:=json.Marshal(v)
	fmt.Println(string(res))
}