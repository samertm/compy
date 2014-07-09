package server

import (
	"github.com/samertm/compy/engine"

	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// warning: modifies req by calling req.ParseForm()
func parseForm(req *http.Request, values ...string) (form url.Values, err error) {
	req.ParseForm()
	form = req.PostForm
	err = checkForm(form, values...)
	return
}

func checkForm(data url.Values, values ...string) error {
	for _, s := range values {
		if len(data[s]) == 0 {
			return errors.New(s + " not passed")
		}
	}
	return nil
}

func handleCommentsAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		form, err := parseForm(r, "pageid", "time", "author", "email", "body")
		if err != nil {
			fmt.Println(err)
			return
		}
		c := engine.NewComment(form["time"][0], form["author"][0],
			form["email"][0], form["body"][0])
		comments.Save <- engine.SavePair{
			PageId:  form["pageid"][0],
			Comment: c,
		}
		io.WriteString(w, "more yay")
	}
}

func handleCommentsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		form, err := parseForm(r, "pageid")
		if err != nil {
			fmt.Println(err)
		}
		c := make(chan []byte)
		comments.Get <- engine.GetPair{
			PageId:   form["pageid"][0],
			Comments: c,
		}
		if comms, ok := <-c; ok {
			// TODO get equiv write-all function for bytes
			io.WriteString(w, string(comms))
		}
	}
}

var comments *engine.Comments

func init() {
	comments = engine.NewComments()
	go comments.Run()
}

func ListenAndServe(ip string) {
	http.HandleFunc("/comments/add", handleCommentsAdd)
	http.HandleFunc("/comments/get", handleCommentsGet)
	http.ListenAndServe(ip, nil)
}
