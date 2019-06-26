package skel_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simplejia/utils"
)

// SkelUpsertReq 定义输入
type SkelUpsertReq Skel

// Regular 用于参数校验
func (skelUpsertReq *SkelUpsertReq) Regular() (ok bool) {
	if skelUpsertReq == nil {
		return
	}

	ok = true
	return
}

// SkelUpsertResp 定义输出
type SkelUpsertResp Skel

func SkelUpsert(name string, req *SkelUpsertReq, trace *utils.Trace) (resp *SkelUpsertResp, result *utils.Resp, err error) {
	addr, err := utils.NameWrap(name)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("http://%s/%s", addr, "skel/upsert")
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
		Data *SkelUpsertResp `json:"data"`
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
