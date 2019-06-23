package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
	"github.com/simplejia/lib"
)

// SkelPageListReq 定义输入
type SkelPageListReq struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// Regular 用于参数校验
func (skelPageListReq *SkelPageListReq) Regular() (ok bool) {
	if skelPageListReq == nil {
		return
	}

	if skelPageListReq.Limit <= 0 {
		skelPageListReq.Limit = 20
	}

	ok = true
	return
}

// SkelPageListResp 定义输出
type SkelPageListResp struct {
	List   []*Skel `json:"list,omitempty"`
	Offset int     `json:"offset,omitempty"`
	Limit  int     `json:"limit,omitempty"`
	More   bool    `json:"more,omitempty"`
	Total  int     `json:"total,omitempty"`
}

func SkelPageList(name string, req *SkelPageListReq, trace *lib.Trace) (resp *SkelPageListResp, result *lib.Resp, err error) {
	addr, err := lib.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/page_list")
	gpp := &utils.GPP{
		Uri:            uri,
		ConnectTimeout: time.Millisecond * 40,
		Timeout:        time.Second * 60,
		Params:         req,
		Headers: map[string]string{
			"h_trace": trace.Encode(),
		},
	}

	body, err := utils.Post(gpp)
	if err != nil {
		return
	}

	s := &struct {
		lib.Resp
		Data *SkelPageListResp `json:"data"`
	}{}
	err = json.Unmarshal(body, s)
	if err != nil {
		return
	}

	if s.Ret != lib.CodeOk {
		result = &s.Resp
		return
	}

	resp = s.Data
	return
}
