package main

//vscode://vscode.github-authentication/did-authenticate?windowid=1&code=c5eaad32ed0b4659969d&state=72c5ce73-68fb-4838-ace8-6b60370c3f74

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type gitHello struct {
	msg, bless string
}

func (gt gitHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf(" Message comes from '%s' \n Blessed from the '%s'\n", gt.msg, gt.bless)
	io.WriteString(w, message)
	if r.URL.Path == "/god" {
		values := r.URL.Query()
		fmt.Fprintf(w, fmt.Sprintf(" God name: %s \n Traditional Name: %s \n", strings.Trim(values["name"][0], ""), strings.Trim(values["alias"][0], "")))
	}
}

func main() {
	gt := &gitHello{
		msg:   "git Say Hello",
		bless: "Hara Hara Mahadeva",
	}
	http.ListenAndServe(":8080", gt)
}
