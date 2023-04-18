package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       int
	Username string
}

type Server struct {
	db    map[int]*User
	dbhit int
}

func NewServer() *Server {
	db := make(map[int]*User)

	for i := 1; i <= 100; i++ {
		db[i] = &User{
			ID:       i,
			Username: fmt.Sprintf("user_%d", i),
		}
	}

	return &Server{
		db: db,
	}
}


func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	//hit the database
	user, ok := s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbhit++

	json.NewEncoder(w).Encode(user)
}

func main() {

}
