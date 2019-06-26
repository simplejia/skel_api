package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
)

// SkelAddReq 定义输入
type SkelAddReq Skel

// Regular 用于参数校验
func (skelAddReq *SkelAddReq) Regular() (ok bool) {
	if skelAddReq == nil {
		return
	}

	ok = true
	return
}

// SkelAddResp 定义输出
type SkelAddResp Skel

func SkelAdd(name string, req *SkelAddReq, trace *utils.Trace) (resp *SkelAddResp, result *utils.Resp, err error) {
	addr, err := utils.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/add")
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
		Data *SkelAddResp `json:"data"`
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
