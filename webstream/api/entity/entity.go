package entity

import "time"

type User struct {
	Id         int64
	Username   string
	Password   string
	CreateTime time.Time
	UpdateTime time.Time
}

type Video struct {
}

type Comment struct {
}
