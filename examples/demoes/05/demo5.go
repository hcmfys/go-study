package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
)

func main() {
	//创建新窗口
	//并设置窗口大小
	w, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{0, 0, 500, 500})
	if err != nil {
		log.Fatal(err)
	}
	//加载文件
	//p,_:=filepath.Abs("")
	w.LoadFile("/Volumes/e/web/go-rpc-app/examples/demoes/05/demo5.html")
	//设置标题
	w.SetTitle("固定大小窗口")
	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}
