// GOOS=linux GOARCH=amd64 go build webhook.go
package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/gitlab"
	"net/http"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := gitlab.New(gitlab.Options.Secret("MyGitHubSuperSecret123456"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, gitlab.PushEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
			}
		}

		switch payload.(type) {

		case gitlab.PushEventPayload:
			release := payload.(gitlab.PushEventPayload)
			fmt.Printf("%+v", release)
			// 在此处实现自己的业务逻辑 即当有push 事件时 要做什么。。

		}

	})
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		return
	}
}
