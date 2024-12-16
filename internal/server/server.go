// server.go
package server

import (
	"html/template"
)

type Server struct {
	Templates *template.Template
}

func NewServer(templates *template.Template) *Server {
	return &Server{Templates: templates}
}
