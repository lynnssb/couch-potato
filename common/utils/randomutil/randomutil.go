/**
 * @author:       wangxuebing
 * @fileName:     randomutil.go
 * @date:         2022/11/1 11:43
 * @description:
 */

package randomutil

import (
	"math/rand"
	"time"
)

// GetRandomNumStr 随机生成指定长度的数字字符串
func GetRandomNumStr(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomStr 随机生成指定长度的字符串
func GetRandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
