package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	server := http.FileServer(http.Dir("E:\\BD\\"))
	http.Handle("/client/", http.StripPrefix("/client/", server))
	http.HandleFunc("/js/", js)
	http.HandleFunc("/", route)
	http.ListenAndServe(":1789", nil)
}
func route(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "welcome")
	r.Body.Close()
}

//设置单文件访问,不能访问目录
func js(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("不能访问目录:%s\n", r.URL.Path)
	old := r.URL.Path
	name, _ := filepath.Abs(filepath.Clean(strings.TrimLeft(old, "/js")))
	info, err := os.Lstat(name)
	if err == nil {
		if !info.IsDir() {
			http.ServeFile(w, r, name)
		} else {
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}
