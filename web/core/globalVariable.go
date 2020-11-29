package core

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	MysqlDriverFormatter = "%s:%s@%s"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)
