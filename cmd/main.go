//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/27 17:56
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"istio-project/internal/app/config"
	"istio-project/internal/pkg/util"
	web_register "istio-project/internal/pkg/web"
	"net/http"
	"strings"
	"time"
)

func main() {
	register()
}

func callTest() {
	var url = "http://istio-project.test.com/test/name"
	var count1 = 0
	var count2 = 0
	for {
		time.Sleep(time.Millisecond * 500)
		var response, _ = http.Get(url)
		var content, _ = ioutil.ReadAll(response.Body)
		var data = string(content)
		if strings.Contains(data, "---1---") {
			count1++
			fmt.Println("微服务1：", count1)
		} else {
			count2++
			fmt.Println("微服务2：", count2)
		}
		_ = response.Body.Close()
	}
}

func register() {
	util.GetLogger().Println("====微服务程序启动====")

	//iris web服务注册
	web_register.InitWeb(func() {
		//程序结束回调函数，此处将执行注销服务注册服务
	}, func(application *iris.Application) {
		//在此处注册mvc，mvc注册逻辑不应该放在pkg包中，因为这个是属于业务逻辑的注册，不是通用逻辑的注册，目前是放在config下
		config.RegisterIrisMvc(application)
		//还可以进行本地过滤器的添加, 如需实现本地过滤器请参考iris官方文档
		// application.Use(filter.WebFilter) //注册web拦截器
	})
}
