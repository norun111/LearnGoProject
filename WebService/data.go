package main

import "database/sql"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "username=tomoyaueno dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error){
	var post Post
	err = Db.QueryRow("select id, content, author from posts where id = $1 ", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		panic(err)
	}
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($2, $3) retirning id"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		panic(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	err = Db.Exec("update posts set content = $2 author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) deletek() (err error) {
	err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

