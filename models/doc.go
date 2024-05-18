package models

import "time"

type Doc struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Created_at time.Time `json:"created_at"`
	Deleted_at time.Time `json:"deleted_at"`
	Checksum []byte `json:"checksum"`
}

