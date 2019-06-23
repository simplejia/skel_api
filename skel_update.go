package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
	"github.com/simplejia/lib"
)

// SkelUpdateReq 定义输入
type SkelUpdateReq Skel

// Regular 用于参数校验
func (skelUpdateReq *SkelUpdateReq) Regular() (ok bool) {
	if skelUpdateReq == nil {
		return
	}

	ok = true
	return
}

// SkelUpdateResp 定义输出
type SkelUpdateResp Skel

func SkelUpdate(name string, req *SkelUpdateReq, trace *lib.Trace) (resp *SkelUpdateResp, result *lib.Resp, err error) {
	addr, err := lib.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/update")
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
		Data *SkelUpdateResp `json:"data"`
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
