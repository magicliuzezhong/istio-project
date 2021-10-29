//
// Package filter
// @Description：web过滤器
// @Author：liuzezhong 2021/6/25 6:01 下午
// @Company cloud-ark.com
//
package filter

import (
	"github.com/kataras/iris/v12/context"
)

//
// WebFilter
// @Description: 普通web过滤器
// @param context iris上下文
//
func WebFilter(context context.Context) {
	var url = context.Request().URL
	if url.String() == "/health" || url.String() == "/metrics" { //如果是健康检查或者普罗米修斯那么不进行全链路追踪
		context.Next()
		return
	}
	context.Next()
}
