package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (content, user_id, created_at) values (?,?,?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `select id, content, user_id, created_at from todos where id = ?`
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

// func GetTodos() (todos []Todo, err error) {
// 	// すべてのテーブルの内容を取っている
// 	cmd := `select id, content, user_id, created_at from todos`
// 	rows, err := Db.Query(cmd)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	for rows.Next() {
// 		var todo Todo
// 		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)

// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		todos = append(todos, todo)
// 	}
// 	rows.Close()

// 	return todos, err
// }

//全てのテーブルの内容を取得
//ループでスキャンして取る
//todosにappendする
//cmd := `select id, content, user_id, created_at from todos`

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err

}

// UserIdによってセレクトをする。
// UserIDが同じものが複数ある可能性があるからループして取り出す
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	//user_idがあっている場所を探す
	cmd := `select id, content, user_id, created_at from todos
		where user_id = ?`
	// cmdとu.IDでselect
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	//rowsの行のデータ取得の準備をする
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//UserIdによってセレクトをする。
//UserIDが同じものが複数ある可能性があるからループして取り出す

func GetTodosByUser(userId int) (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, userId)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
