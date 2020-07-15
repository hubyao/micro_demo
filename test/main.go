/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 11:16:29
 * @LastEditTime : 2020-07-14 11:32:23
 * @Description  : file content
 */ 
 
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

func main() {
    resp, err := http.Get( //获取数据
        "http://cp.zgzcw.com/lottery/ssq/index.jsp")
    //异常处理
    if err != nil {
        panic(err)
    }
    //清理变量
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { //访问页面是否能访问通
        fmt.Println("Error:status code ",
            resp.StatusCode)
        return
    }

    all, err := ioutil.ReadAll(resp.Body) //获取body数据
    //捕捉异常
    if err != nil {
        panic(err)
	}
	
    //编写正则表达式
    re := regexp.MustCompile(`<span class="animate_cur">[0-9]*</span>`)
    //靠正则规则来获取自己想要的数据
	match := re.FindAllSubmatch(`<span class="animate_cur">[0-9]*</span>`, -1)
    //循环输出数据
    for _, m := range match {
        fmt.Printf("Name: %s, URL: %s \n ", m[2], m[1])
        fmt.Printf("%s\n", m)
    }
    //输出数据长度
    fmt.Printf("Match found:%d\n",
        len(match))
}