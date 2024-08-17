package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
	uuid,
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
//
// GetUser関数、 id, uuid, name, email, password, created_atを取ってくる
func GetUser(id int) (user *User, err error) {
	user = &User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return user, err
}

// func GetUser(id int) (user User, err error) {
// 	user = User{}
// 	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
// Scan メソッドを使用して、データベースから取得した値を直接 User 構造体のフィールドに代入しています。
// 	err = Db.QueryRow(cmd, id).Scan(
// 		&user.ID,
// 		&user.UUID,
// 		&user.Name,
// 		&user.Email,
// 		&user.PassWord,
// 		&user.CreatedAt,
// 	)
// 	return user, err
// }

// nameとemailの値を変えられる関数を書く
// 	cmd := `update users set name = ?, email = ? where id = ?`

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln()
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
