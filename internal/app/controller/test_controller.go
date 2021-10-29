//
// Package controller
// @Description：测试控制器
// @Author：liuzezhong 2021/6/25 6:45 下午
// @Company cloud-ark.com
//
package controller

import (
	irisContext "github.com/kataras/iris/v12/context"
	"io/ioutil"
	"net/http"
)

type Test struct {
	Age int `json:"age"`
}

type TestController struct {
	Ctx irisContext.Context
}

func (c TestController) GetName() {
	defer c.Ctx.Next()
	c.Ctx.Values().Set("val", "微服务---2---")
}
func (c TestController) GetSite() {
	defer c.Ctx.Next()
	var path = "http://istio-project-service2.test.svc.cluster.local/test/name"
	var response, err = http.Get(path)
	if err != nil {
		c.Ctx.Values().Set("val", "请求异常，"+err.Error())
		return
	}
	defer func() {
		_ = response.Body.Close()
	}()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.Ctx.Values().Set("val", "解析异常，"+err.Error())
		return
	}
	c.Ctx.Values().Set("val", "请求成功，"+string(data))
}
