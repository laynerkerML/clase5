package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
		id := c.Query("id")
		if id != "" {
			fmt.Println("filter: ", id)
			c.JSON(200, GetFilter("id", id))
			return
		}
		c.JSON(200, GetAll())
	})

	r.GET("/users/:id", getById)

	r.Run()
}

func GetAll() Accesos {
	file := getFile()
	datos := Accesos{}
	_ = json.Unmarshal(file, &datos)
	return datos
}

func GetFilter(fieldFilter string, valueFilter string) Accesos {
	file := getFile()
	datos := Accesos{}
	_ = json.Unmarshal([]byte(file), &datos)
	filter := new(Accesos)
	for _, u := range datos.Users {
		fmt.Println("id: ", u.Id)
		if fieldFilter == "id" {
			id, _ := strconv.Atoi(valueFilter)
			if u.Id == id {
				filter.Users = append(filter.Users, u)
			}
		}
	}
	return *filter
}

func getFile() []byte {
	files, err := os.ReadFile("../users.json")
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"menssge": "Parametro no valido!",
		})
		return
	}
	file := getFile()
	datos := Accesos{}
	_ = json.Unmarshal([]byte(file), &datos)
	filter := new(User)
	for _, u := range datos.Users {
		if u.Id == id {
			*filter = u
		}
	}
	if filter.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"menssge": "No existe el registro!",
		})
		return
	}
	c.JSON(http.StatusOK, filter)
}
