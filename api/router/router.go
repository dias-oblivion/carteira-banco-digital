package router

import "github.com/gin-gonic/gin"

func InitRouter() {
	// inicializa o router do gin no modo padrão
	r := gin.Default()

	// inicializa as rotas
	InitRoutes(r)

	// inicializa a api na porta 8081
	r.Run(":8081")
}
