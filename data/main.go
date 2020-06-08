package main

type Post struct {
	Id int
	Content string
	Author string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {

	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id:1, Content:"Hello", Author:"Sau Sheong"}

	store(post1)
}
