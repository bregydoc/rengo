package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

const urlPEL = "https://www.peruanosenlinea.com/demo/reniec/example/consulta.php"

func main() {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Localization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	engine.Any("/getfromessalud", func(c *gin.Context) {
		dni := struct {
			Dni string `json:"dni"`
		}{}
		c.BindJSON(&dni)

		pp.Println(dni)

		v := url.Values{}
		v.Set("ndni", dni.Dni)
		v.Set("source", "EsSalud")
		resp, err := http.PostForm(urlPEL, v)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var response interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			panic(err)
		}
		c.JSON(resp.StatusCode, response)
	})

	engine.Run(":3300")
}
