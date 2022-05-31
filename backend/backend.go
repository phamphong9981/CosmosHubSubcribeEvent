package backend

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func getRealtime(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//log.Print()
	// for {
	// 	msg, err := subscribeAllChannel.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(msg.Channel, msg.Payload)
	// }
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

	http.HandleFunc("/websocket", getRealtime)
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
