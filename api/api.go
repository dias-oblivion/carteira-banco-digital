package api

import (
	"net/http"

	types "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/database"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	port  string
	store database.Store
}

func NewAPIServer(port string, store database.Store) *APIServer {
	return &APIServer{
		port:  port,
		store: store,
	}
}

func (s *APIServer) Start() {
	router := gin.Default()

	router.POST("/user", func(ctx *gin.Context) {
		var user types.CreateUserRequest
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId, err := s.store.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"id": userId})
	})

	router.Run(s.port)
}
