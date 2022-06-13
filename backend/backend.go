package backend

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func serveClient(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	connectClient(con, "all")

}

func serveClientByDelegator(c *gin.Context) {
	var delegator = c.Param("delegator")
	log.Print(delegator)
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	connectClient(con, delegator)
}

func Run() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	r.GET("/unbond/:validator", func(c *gin.Context) {
		var validator = c.Param("validator")
		c.JSONP(200, getUnbondFromValidator(validator))
	})
	subscibeInit()
	go handleSubscibe()
	r.GET("/websocket", serveClient)
	r.GET("/websocket/delegator/:delegator", serveClientByDelegator)
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
