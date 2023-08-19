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

// QueryRequest 关键字和分页信息结构体
type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

// GetUserListRequest 获取用户列表参数结构体
type GetUserListRequest struct {
	*QueryRequest
}

// GetUserListReply 返回管理员信息结构体
type GetUserListReply struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
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
	Email    string `json:"email"`
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

// ----------------- 角色模块开始 ----------------------

// GetRoleListRequest 获取角色列表参数结构体
type GetRoleListRequest struct {
	*QueryRequest
}

// GetRoleListReply 返回角色列表结构体
type GetRoleListReply struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	IsAdmin   int    `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddRoleRequest 新增角色参数结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Sort    int64  `json:"sort"`
	IsAdmin int8   `json:"isAdmin"`
	Remarks string `json:"remarks"`
}

// ----------------- 角色模块结束 ----------------------
