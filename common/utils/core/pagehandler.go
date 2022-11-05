/**
 * @author:       wangxuebing
 * @fileName:     pagehandler.go
 * @date:         2022/11/5 22:14
 * @description:
 */

package core

import "couch-potato/common/utils/characterutil"

type PageResult struct {
	Page     int     //页码
	PageSize int     //每页数量
	Filter   *string //关键字
}

func PageHandler(page, pageSize int, filter *string) PageResult {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize >= 1000 {
		pageSize = 1000
	}
	if filter != nil {
		_filter := characterutil.StitchingBufStr("%", *filter, "%")
		filter = &_filter
	}
	return PageResult{
		Page:     (page - 1) * pageSize,
		PageSize: pageSize,
		Filter:   filter,
	}
}

func Filter(filter string) string {
	if len(filter) > 0 {
		_filter := characterutil.StitchingBufStr("%", filter, "%")
		return _filter
	}
	return ""
}
