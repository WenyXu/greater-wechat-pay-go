/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/06 20:48
*/

package storage

type Storage interface {
	Set(key, value interface{}) error
	Get(key interface{}) (interface{}, bool)
}
