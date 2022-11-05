/**
 * @author:       wangxuebing
 * @fileName:     model.go
 * @date:         2022/11/5 22:14
 * @description:
 */

package core

type CommonId struct {
	ID string `json:"id"`
}

type CommonKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
