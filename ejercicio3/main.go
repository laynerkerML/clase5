package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id            int
	Nombre        string
	Apellido      string
	Email         string
	Edad          int
	Altura        int
	Activo        bool
	FechaCreacion string
}

type Accesos struct {
	Users []User
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"menssge": "Hola laynerker!",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, GetAll())
	})

	r.Run()
}

func GetAll() Accesos {
	files, err := os.ReadFile("../users.json")
	datos := Accesos{}
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(files, &datos)
	return datos
}
