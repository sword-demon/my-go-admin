package service

import "my-go-admin/define"

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		Keyword: "",
	}
}
