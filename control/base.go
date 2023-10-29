package control

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var statusChan = make(chan string, 20)
var wg sync.WaitGroup

type MyForm struct {
	ProjectName string `form:"project_name"`
	Port        int    `form:"port"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func CHttp() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html") // Load HTML templates from the "templates" directory

	r.GET("/", func(c *gin.Context) {
		data := gin.H{
			"Title": "Welcome to My Web Page",
		}
		c.HTML(http.StatusOK, "index.html", data) // Render the HTML template
	})
	// Handle the form submission
	r.POST("/submit", func(c *gin.Context) {
		var form MyForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//wg.Add(1)
		go func() {
			project_name = form.ProjectName
			No = form.Port
			GenLaravel(statusChan)
			//GenLaravel()
		}()
		//wg.Wait()
		//c.HTML(http.StatusOK, "redirect.html", data) // Render the HTML template
		// 回應JSON

		c.JSON(http.StatusOK, gin.H{
			"ProjectName": form.ProjectName,
			"Port":        form.Port,
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		// 處理 WebSocket 連接
		handleWebSocket(c.Writer, c.Request)
	})

	r.Run(":8080")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	go func() {
		for {
			select {
			case status := <-statusChan:
				// 在这里将状态信息发送到WebSocket连接
				if err := conn.WriteMessage(websocket.TextMessage, []byte(status)); err != nil {
					return
				}
			}
		}
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println("hi", messageType)
		fmt.Println("p=", p)
		// 这里可以处理来自WebSocket客户端的消息
	}
}
