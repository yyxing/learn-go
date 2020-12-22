package core

const (
	MysqlDriverFormatter = "%s:%s@%s"
	Int32Min             = ^int(^uint64(0) >> 1)
	Int32Max             = int(^uint64(0) >> 1)
	ServerPortKey        = "server.port"
)
