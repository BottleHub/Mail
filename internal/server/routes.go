package server

import (
	"mail-client/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db, _ = database.ConnectDB()
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.POST("/add", db.AddEmail())

	return r
}
