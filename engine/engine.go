package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Comment struct {
	id     int       `json:"id"` // unique id
	Time   time.Time `json:"time"`
	Author string    `json:"author"`
	Email  string    `json:"email"`
	Body   string    `json:"body"`
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
	Comments chan []byte
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
				j, err := json.Marshal(comments)
				if err != nil {
					log.Print(err)
				} else {
					g.Comments <- j
				}
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

func NewComment(author, email, body string) *Comment {
	id := commentId
	commentId++
	return &Comment{
		id:     id,
		Time: time.Now(),
		Author: author,
		Email:  email,
		Body:   body,
	}
}
