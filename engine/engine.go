package engine

import (
	"fmt"
	"strconv"
	"time"
)

type Comment struct {
	id     int // unique id
	Time   time.Time
	Author string
	Email  string
	Body   string
}

func (c *Comment) String() string {
	s := "id: " + strconv.Itoa(c.id) + "\n" +
		"time: " + c.Time.String() + "\n" +
		"author: " + c.Author + "\n" +
		"email: " + c.Email + "\n" +
		"body: " + c.Body + "\n"
	return s
}

type GetPair struct {
	PageId   string // page id
	Comments chan []*Comment
}

type SavePair struct {
	PageId  string
	Comment *Comment
}

type Comments struct {
	Get chan GetPair
	// TODO somehow get error values from this?
	Save chan SavePair
	// maps from page ids (md5 hashes of urls) to
	// the comments associated with that page.
	// TODO order slice by date
	pages map[string][]*Comment
}

func NewComments() *Comments {
	return &Comments{
		Get:   make(chan GetPair),
		Save:  make(chan SavePair),
		pages: make(map[string][]*Comment),
	}
}

func (c *Comments) Run() {
	for {
		select {
		case g := <-c.Get:
			if comments, ok := c.pages[g.PageId]; ok {
				g.Comments <- comments
			}
			close(g.Comments)
		case s := <-c.Save:
			var comments []*Comment
			if storedComments, ok := c.pages[s.PageId]; ok {
				comments = storedComments
			} else {
				comments = make([]*Comment, 0)
			}
			fmt.Println(*s.Comment)
			c.pages[s.PageId] = append(comments, s.Comment)
		}
	}
}

var commentId int

// TODO add time stuff :D
func NewComment(time, author, email, body string) *Comment {
	// t, err := time.Parse(what do i do???)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	id := commentId
	commentId++
	return &Comment{
		id:     id,
		Author: author,
		Email:  email,
		Body:   body,
	}
}