package controllers

import (
	"another_todo_app/config"
	"fmt"
	"html/template"
	"net/http"
)

// 複数の受取を可能にしている ...string
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	//ここでfilesnamesのSliceを受け取るためにFilesでSlice型にしている
	var files []string
	//fileに入れている
	for _, file := range filenames {
		//今回の場合はLayoutを先に入れて、次にTopを入れる
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	//Mustはページが存在しているかNilじゃないかを確認している
	//ParseFilesは解析された内容を基に、*template.Template 型のテンプレートオブジェクトを作成します
	templates := template.Must(template.ParseFiles(files...))
	//laoyutにTopがあるので上でスライスを解析させている。そしてここでdataを渡している。
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
