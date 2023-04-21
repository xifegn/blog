// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
