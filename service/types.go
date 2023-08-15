package service

// LoginPasswordRequest 登录参数结构体
type LoginPasswordRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// LoginPasswordReply 登录后的token结构体
type LoginPasswordReply struct {
	Token        string `json:"token"`         // token
	RefreshToken string `json:"refresh_token"` // 用于刷新 token 的 token
}

// GetUserListRequest 获取用户列表参数结构体
type GetUserListRequest struct {
	*QueryRequest
}

// QueryRequest 关键字和分页信息结构体
type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"key_word" form:"keyword"`
}

// GetUserListReply 返回管理员信息结构体
type GetUserListReply struct {
	Username  string `json:"username"`
	Avatar    string `json:"Avatar"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddUserRequest 添加管理员结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Remarks  string `json:"remarks"`
}

// GetUserDetailReply 获取管理员结构体
type GetUserDetailReply struct {
	ID uint `json:"id"`
	AddUserRequest
}

// UpdateUserRequest 更新管理员信息结构体
type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}
