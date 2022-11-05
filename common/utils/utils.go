/**
 * @author:       wangxuebing
 * @fileName:     utils.go
 * @date:         2022/11/5 22:13
 * @description:
 */

package utils

import (
	"net/url"
	"os"
)

// EncodeURL URL编码
func EncodeURL(api string, params map[string]string) (string, error) {
	reqUrl, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	query := reqUrl.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	reqUrl.RawQuery = query.Encode()
	result, _ := url.QueryUnescape(reqUrl.String())
	return result, nil
}

// PathExists 判断文件夹是否存在,不存在则创建
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err == nil {
			return nil
		} else {
			return err
		}
	}
	return err
}
