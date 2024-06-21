package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		//ここでのsessはインスタンス。だけどUUIDの部分しかオブジェクトに埋まっていない状態。ほかは空
		sess = models.Session{UUID: cookie.Value}
		//okはvalidのブール値を保持している
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// これはURLのパスの確認をしている。todo/edit/数字もしくはtodo/update/数字じゃないとValidPathとならない
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9])")

// 正規表現によるパスの検証:
// parseURL 内で正規表現により、/todos/edit/数字 というパスの形式が保証されます。これにより、id の部分が整数であることが確認され、安全に strconv.Atoi(q[2]) を使って整数に変換できます。

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check2")
		// /todos/edit/1をいれる。FindStringSubmatchは同じString部分を返すから、違う所があればNilが返ってくるんだと思う
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			fmt.Println("check3")
			http.NotFound(w, r)
			return
		}
		// AtoiはこのStringをintの整数に変える、q2は正規表現の数字の場所を表している
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			fmt.Println("check4")
			http.NotFound(w, r)
			return
		}
		fmt.Println("check5")
		//いめーじとしてはreturn functionみたいなかんじ。
		fn(w, r, qi)
	}
}

//parseURL

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
