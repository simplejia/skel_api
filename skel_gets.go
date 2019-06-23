package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
	"github.com/simplejia/lib"
)

// SkelGetsReq 定义输入
type SkelGetsReq struct {
	IDS []int64 `json:"ids"`
}

// Regular 用于参数校验
func (skelGetsReq *SkelGetsReq) Regular() (ok bool) {
	if skelGetsReq == nil {
		return
	}

	ok = true
	return
}

// SkelGetsResp 定义输出
type SkelGetsResp map[int64]*Skel

func SkelGets(name string, req *SkelGetsReq, trace *lib.Trace) (resp *SkelGetsResp, result *lib.Resp, err error) {
	addr, err := lib.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/gets")
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
		Data *SkelGetsResp `json:"data"`
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
