package server

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	data "overlaysr/server/internal/pkg/data"

	"github.com/gorilla/websocket"
)

type server interface {
	/**TODO:
	* Server 启动后一直监听端口
	 */
	serving(w http.ResponseWriter, r *http.Request)

	/**TODO:
	接收server 的更新信息并放入 update chan
	*/
	update(update chan data.Message)
}

type WsServer struct {
	addr     string
	message  chan data.Message
	upgrader websocket.Upgrader
}

func (ws *WsServer) Init() {
	ws = &WsServer{
		addr:     *flag.String("addr", "localhost:8080", "http service address"),
		message:  make(chan data.Message, 100),
		upgrader: websocket.Upgrader{},
	}
}

func (ws *WsServer) AddrUp() {
	ws.addr = *flag.String("addr", "localhost:8080", "http service address")
}

func (ws *WsServer) GetAddr() string {
	return ws.addr
}

func (ws *WsServer) Serving(w http.ResponseWriter, r *http.Request) {

	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv from client: %s", message)
		reply := "ok"
		msg, err := json.Marshal(reply)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
