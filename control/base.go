package control

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	db "tpl/db/sqlc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var statusChan = make(chan string, 20)

//var wg sync.WaitGroup

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:    4096,
// 	WriteBufferSize:   4096,
// 	EnableCompression: true,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

var ua string
var project Pj
var mu sync.Mutex

type MyForm struct {
	Name   string                   `json:"name"`
	Port   int32                    `json:"port"`
	Data   interface{}              `json:"state"`
	Tables []map[string]interface{} `json:"Tables"`
}

type Table struct {
	Name string            `json:"name"`
	Attr map[string]string `json:"attr"`
	Pid  int32             `json:"id"`
}

func CHttp() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.POST("/generateLaravel", func(c *gin.Context) {
		GenerateLaravelByLest()
		c.JSON(http.StatusOK, gin.H{"message": "auto generate laravel successfully"})
	})
	r.POST("/createtable", func(c *gin.Context) {
		var table Table
		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, value := range table.Attr {
			fmt.Println("Original value:", value)
			values, err := url.ParseQuery(value)
			if err != nil {
				fmt.Println("Error parsing query string:", err)
				return
			}
			// 直接使用 json.Unmarshal 解析 JSON 字符串
			// 现在 values 是一个 url.Values 对象，可以通过键获取值
			fieldName := values.Get("fieldname")
			showName := values.Get("showname")
			modelType := values.Get("migration_modelType")
			isRequire := values.Get("isrequire")
			requires, err := strconv.ParseInt(isRequire, 10, 0)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			modelTypeParts := strings.Split(modelType, "_")
			// 输出解析后的值
			a := LaraSetting{
				Field:     fieldName,
				ShowName:  showName,
				Migration: modelTypeParts[0],
				ModelType: modelTypeParts[1],
				IsRequire: int32(requires),
			}

			pr := Pj{
				Pg:        db.ConnDev(),
				ProjectID: table.Pid,
				TempField: a,
			}
			mu.Lock()
			defer mu.Unlock()
			if pr.CheckTable(table.Name, table.Pid) {
				pr.GenTable(table.Name)
			}
			pr.ExecCreateTableField()
		}
	})
	// Handle the form submission
	r.POST("/submit", func(c *gin.Context) {
		var form MyForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		project = Pj{
			Pg:          db.ConnDev(),
			ProjectName: form.Name,
			DockerPort:  form.Port,
		}

		// if !project.ChkProjectName() {
		project.GenProject()
		// }
		// 现在 paramMap 包含了所有的参数
		// data := ParseData(ddata)
		// fmt.Println(data)

		// 解析动态的表格参数
		// 打印动态生成的字段
		//wg.Add(1)
		//go func() {
		//	project_name = form.ProjectName
		//	No = form.Port
		//	//GenLaravel(statusChan)
		//}()
		//wg.Wait()
		//c.HTML(http.StatusOK, "redirect.html", data) // Render the HTML template
		// 回應JSON
		c.JSON(http.StatusOK, gin.H{"id": project.ProjectID, "message": "JSON data received successfully"})
	})

	r.GET("/ws", func(c *gin.Context) {
		// 處理 WebSocket 連接
		handleWebSocket(c.Writer, c.Request)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"resp": "pong",
		})
	})

	// Or, customize the CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))
	r.Run(":8080")
}

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func Gomain() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

					fmt.Println(websocket.TextMessage)
					return
				}
			}
		}
	}()

	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Println("WebSocket error:", err)
				}
				return
			}
			fmt.Printf("Received message: %s\n", string(p))
			if messageType == websocket.TextMessage {
				// 处理文本消息
				fmt.Printf("Received message: %s\n", string(p))
			}
		}
	}()

}
