package introduction

import "net/http"

func Server() {
	http.Handle("/", http.FileServer(http.Dir("."))) // 将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	http.ListenAndServe(":1412", nil)
}
