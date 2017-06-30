package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/object88/bbreloader/examples/multi-project/server"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		if len(text) == 0 {
			readTwoots()
		} else {
			postTwoot(text)
		}
	}
}

func readTwoots() {
	var buffer []byte
	var twoots server.Twoots

	response, _ := netClient.Get("http://localhost:8181")
	response.Body.Read(buffer)
	err := json.Unmarshal(buffer, &twoots)
	if err != nil {
		fmt.Print("Failed to get twoots from server\n")
		return
	}
}

func postTwoot(message string) {

}
