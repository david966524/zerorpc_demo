// Code generated by goctl. DO NOT EDIT.
package types

type DeleteUserInfoId struct {
	Id int64 `path:"id"`
}

type GetUserInfoReq struct {
	Id int64 `json:"id"`
}

type GetUserInfoResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type RespDomainInfo struct {
	Id         string `json:"id"`
	DomainName string `json:"domainName"`
	Status     string `json:"status"`
}

type UpdateUserInfoReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
