package models

//ParamLogin 用户名，密码登录需要的参数
type ParamLogin struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	ReCAPTCHA string `form:"g_recaptcha_response" json:"g_recaptcha_response" binding:"required"`
}

//ParamsAddModule 添加xss模块需要的参数
type ParamsAddModule struct {
	Name       string   `json:"name" binding:"required,max=50"`
	Desc       string   `json:"desc" binding:"required,max=300"`
	XSSPayload string   `json:"xss_payload" binding:"required"`
	ParamList  []string `json:"param_list" binding:"required"`
	OptionList []string `json:"option_list" binding:"required"`
	IsCommon   *bool    `json:"is_common" binding:"required"`
}

//ParamsUpdateModule 更新xss模块需要的参数
type ParamsUpdateModule struct {
	ID         int64    `json:"id" binding:"required"`
	Name       string   `json:"name" binding:"required,max=50"`
	Desc       string   `json:"desc" binding:"required,max=300"`
	XSSPayload string   `json:"xss_payload" binding:"required"`
	ParamList  []string `json:"param_list" binding:"required"`
	OptionList []string `json:"option_list" binding:"required"`
	IsCommon   *bool    `json:"is_common" binding:"required"`
}

//ParamsRegister 注册需要参数
type ParamsRegister struct {
	Password  string `json:"password" binding:"required"`
	ReCAPTCHA string `form:"g_recaptcha_response" json:"g_recaptcha_response" binding:"required"`
}

//ParamLogout 退出登录
type ParamLogout struct {
	Token string `json:"token" binding:"required"`
}

//ParamCreateProject 创建项目需要的参数
type ParamCreateProject struct {
	Name       string       `json:"name" binding:"required"`
	Desc       string       `json:"desc" binding:"required"`
	ModuleList []ModuleItem `json:"module_list" binding:"required"`
}

type ModuleItem struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	IsChoosed  bool         `json:"is_choosed"`
	OptionList []OptionItem `json:"option_list"`
}

type OptionItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
