package main

//por padrão se eu não definir um nome, ele considera o nome do pacote (gin) depois da ultima barra
import "github.com/gin-gonic/gin"

func main() {
	//Inicializa o router com as configurações padrão do gin
	router := gin.Default()
	//Define uma rota GET para o caminho "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
