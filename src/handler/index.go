package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Scfy-Code/scfy-im/entry"
)

// IndexHandler 首页
type IndexHandler struct {
	ViewPath string
}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1、验证是否登录(登录后返回页面/未登录重定向登录页面)

	self := entry.UserEntry{"123", "scfy", "774250", "scfymail@gmail.com", "林昊天", "/static/images/avatar.png"}
	talkList := []entry.UserEntry{entry.UserEntry{"456", "tom", "774250", "scfymail@gmail.com", "汤姆", "/static/images/avatar.png"}}
	friendList := []entry.UserEntry{entry.UserEntry{"789", "jerry", "774250", "scfymail@gmail.com", "杰瑞", "/static/images/avatar.png"}}
	groupList := []entry.UserEntry{entry.UserEntry{"101112", "jake", "774250", "scfymail@gmail.com", "杰克", "/static/images/avatar.png"}}

	indexEntry := &entry.IndexEntry{self, talkList, friendList, groupList}
	//2、
	view, err := template.ParseFiles(ih.ViewPath)
	if err == nil {
		view.Execute(w, indexEntry)
	} else {
		log.Println(err.Error())
	}
}
