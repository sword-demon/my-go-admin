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

// GetRoleDetailReply 获取角色详情结构体
type GetRoleDetailReply struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// UpdateRoleRequest 更新角色信息结构体
type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// ----------------- 角色模块结束 ----------------------

// ----------------- 菜单模块开始 ----------------------

// MenuReply 菜单列表返回结构体
type MenuReply struct {
	ID       int          `json:"id"`
	ParentId int          `json:"parent_id"`
	Name     string       `json:"name"`
	WebIcon  string       `json:"web_icon"`
	Sort     int          `json:"sort"`
	Path     string       `json:"path"`
	Level    int          `json:"level"` // 菜单等级，{0：目录，1：菜单，2：按钮}
	SubMenus []*MenuReply `json:"sub_menus"`
}

// AllMenu 所有菜单数据结构体
type AllMenu struct {
	ID       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	WebIcon  string `json:"web_icon"`
	Sort     int    `json:"sort"`
	Path     string `json:"path"`
	Level    int    `json:"level"`
}

// AddMenuRequest 新增菜单结构体
type AddMenuRequest struct {
	ParentId uint   `json:"parent_id"` // 父级唯一标识，不填默认为顶级菜单
	Name     string `json:"name"`      // 菜单名称
	WebIcon  string `json:"web_icon"`  // 网页图标
	Path     string `json:"path"`      // 路径
	Sort     int    `json:"sort"`      // 排序
	Level    int    `json:"level"`     // 菜单等级，{0：目录，1：菜单，2：按钮}
}

// GetMenuDetailReply 返回菜单详情结构体
type GetMenuDetailReply struct {
	ID uint `json:"id"`
	AddMenuRequest
}

// ----------------- 菜单模块结束 ----------------------
