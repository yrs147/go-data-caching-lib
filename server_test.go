package main

import (
	"net/http"
	"testing"
)

func TestHandleGetUser(t *testing.T) {
	s := NewServer()
	ts := NewServer(http.HandlerFunc(s.handleGetUser))
	nreq := 1000

	for i:= 0;i<nreq;i++ {
		id := i %100 +1
		url := fmt.Sprintf("%s/?id=%d",ts.URL , id)
		resp,err := http.Get(url)
		if err !=nil {
			t.Error(err)
		}

		user := &User{}
		if err := json.NewDecoder()
	}
}
