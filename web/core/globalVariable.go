package core

import (
	"github.com/go-playground/validator/v10"
)

const (
	MysqlDriverFormatter = "%s:%s@%s"
	Int32Min             = ^int(^uint64(0) >> 1)
	Int32Max             = int(^uint64(0) >> 1)
	ServerPortKey        = "server.port"
)

var (
	Validate *validator.Validate
)
