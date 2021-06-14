package request

// User register struct
type Register struct {
	Username    string `json:"userName" form:"userName"`
	Password    string `json:"passWord" form:"userName"`
	NickName    string `json:"nickName"`
	HeaderImg   string `json:"headerImg"`
	AuthorityId string `json:"authorityId"`
}
