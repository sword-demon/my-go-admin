package service

import "my-go-admin/define"

// NewQueryRequest 初始化查询列表参数默认值
func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		Keyword: "",
	}
}
