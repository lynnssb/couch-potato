/**
 * @author:       wangxuebing
 * @fileName:     errorx.go
 * @date:         2022/11/5 22:18
 * @description:
 */

package errorx

import (
	"errors"
	"gorm.io/gorm"
)

const (
	defaultCode     = 1001
	databaseErrCode = 1002
)

type CodeErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeErr{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func NewDatabaseError() error {
	return NewCodeError(databaseErrCode, "操作失败[database error]")
}

func NewDatabaseNotFoundError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewCodeError(databaseErrCode, "该数据查询为空")
	}
	return NewCodeError(databaseErrCode, "操作失败[database error]")
}

func (e *CodeErr) Error() string {
	return e.Msg
}

func (e *CodeErr) Data() *CodeErrResp {
	return &CodeErrResp{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
