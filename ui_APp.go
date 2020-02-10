package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("测试")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("输入姓名:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)

		//创建window窗口。并设置长宽。
		window := ui.NewWindow("第一个应用程序。", 600, 500, false)
		//mac不支持居中。
		//https://github.com/andlabs/ui/issues/162
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			//可以直接打印日志。
			fmt.Println("get name :", name.Text())
			greeting.SetText("Hello, " + name.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
