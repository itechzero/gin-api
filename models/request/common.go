package request

// Paging common input parameter struct
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// Find by id struct
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id struct
type GetAuthorityId struct {
	AuthorityId string // 角色ID
}

type Empty struct{}
