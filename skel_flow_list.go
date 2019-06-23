package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
	"github.com/simplejia/lib"
)

// SkelFlowListReq 定义输入
type SkelFlowListReq struct {
	LastID string `json:"last_id,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// Regular 用于参数校验
func (skelFlowListReq *SkelFlowListReq) Regular() (ok bool) {
	if skelFlowListReq == nil {
		return
	}

	if skelFlowListReq.Limit <= 0 {
		skelFlowListReq.Limit = 20
	}

	ok = true
	return
}

// SkelFlowListResp 定义输出
type SkelFlowListResp struct {
	List   []*Skel `json:"list,omitempty"`
	LastID string  `json:"last_id,omitempty"`
	Limit  int     `json:"limit,omitempty"`
	More   bool    `json:"more,omitempty"`
	Total  int     `json:"total,omitempty"`
}

func SkelFlowList(name string, req *SkelFlowListReq, trace *lib.Trace) (resp *SkelFlowListResp, result *lib.Resp, err error) {
	addr, err := lib.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/flow_list")
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
		Data *SkelFlowListResp `json:"data"`
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
