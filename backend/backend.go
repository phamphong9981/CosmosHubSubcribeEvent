package backend

import (
	"encoding/json"
	"log"
	"net/http"

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
	r.GET("/unbond/:validator", func(c *gin.Context) {
		var validator = c.Param("validator")
		c.JSON(200, getUnbondFromValidator(validator))
	})

	r.GET("/test", func(c *gin.Context) {
		data := Data{validator:"1213246545",time: "dgfdsgfd", amount:"fgfdsgdfhgfhgf"}
		jData, _ := json.Marshal(data)
		log.Print(jData)
		c.JSON(200, jData)
	})
	http.HandleFunc("/websocket", getRealtime)
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
