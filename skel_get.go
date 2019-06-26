package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
)

// SkelGetReq 定义输入
type SkelGetReq struct {
	ID int64 `json:"id"`
}

// Regular 用于参数校验
func (skelGetReq *SkelGetReq) Regular() (ok bool) {
	if skelGetReq == nil {
		return
	}

	ok = true
	return
}

// SkelGetResp 定义输出
type SkelGetResp Skel

func SkelGet(name string, req *SkelGetReq, trace *utils.Trace) (resp *SkelGetResp, result *utils.Resp, err error) {
	addr, err := utils.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/get")
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
		utils.Resp
		Data *SkelGetResp `json:"data"`
	}{}
	err = json.Unmarshal(body, s)
	if err != nil {
		return
	}

	if s.Ret != utils.CodeOk {
		result = &s.Resp
		return
	}

	resp = s.Data
	return
}
