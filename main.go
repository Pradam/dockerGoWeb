package main

import (
	"fmt"
	"io"
	"net/http"
)

type gitHello struct {
	msg, bless string
}

func (gt gitHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf(" Message comes from '%s' \n Blessed from the '%s'\n", gt.msg, gt.bless)
	io.WriteString(w, message)
	fmt.Fprintf(w, fmt.Sprintf(" URL is %v \n", r.URL.Path))
}

func main() {
	gt := &gitHello{
		msg:   "git Say Hello",
		bless: "Hara Hara Mahadeva",
	}
	http.ListenAndServe(":8080", gt)
}
