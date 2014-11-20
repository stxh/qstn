package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/daryl/zeus"
)

// Same as bson.M
type M map[string]interface{}

func Map(m *zeus.Mux, n *negroni.Negroni) {
	m.GET("/posts", postsGet)
	m.POST("/posts", postsCreate)

	m.GET("/posts/:id", postsShow)
	m.GET("/posts/:id/download", postsDownload)
	m.GET("/posts/:id/view", postsView)

	m.GET("/users", usersGet)
	m.POST("/users", usersCreate)

	m.GET("/users/:id", usersShow)

	m.GET("/session", sessionGet)
}