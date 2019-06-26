package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
)

// SkelDelReq 定义输入
type SkelDelReq struct {
	ID int64 `json:"id"`
}

// Regular 用于参数校验
func (skelDelReq *SkelDelReq) Regular() (ok bool) {
	if skelDelReq == nil {
		return
	}

	ok = true
	return
}

// SkelDelResp 定义输出
type SkelDelResp struct {
}

func SkelDel(name string, req *SkelDelReq, trace *utils.Trace) (resp *SkelDelResp, result *utils.Resp, err error) {
	addr, err := utils.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/del")
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
		Data *SkelDelResp `json:"data"`
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
