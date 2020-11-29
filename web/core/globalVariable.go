package core

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	MysqlDriverFormatter = "%s:%s@%s"
	Int32Min             = ^int(^uint32(0) >> 1)
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)
