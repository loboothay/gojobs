package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o router utilizando as configurações padroes do Gin
	router := gin.Default()
	//Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rodando a Api e especificando a porta
	router.Run(":8080")
}
