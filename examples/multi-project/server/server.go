package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Twoot is a small message
type Twoot struct {
	// Msg is the main message
	Msg string `json:"message"`

	// When is the time the Twoot was submitted
	When time.Time `json:"when"`

	// Who is the individual who twooted
	Who string `json:"who"`
}

// Twoots is a collection of Twoot
type Twoots []*Twoot

var twoots = Twoots{
	&Twoot{
		"Someone's birthday?",
		time.Date(1994, 6, 23, 0, 0, 0, 0, nil),
		"Mark",
	},
	&Twoot{
		"Almost time!",
		time.Date(1999, 12, 31, 23, 59, 59, 0, nil),
		"Bob",
	},
}

// ServerPort is the port that the server will run on.
const ServerPort = "8181"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "" || r.Method == http.MethodGet {
			// Get the current list of twoots.
			readTwoots(w)
			return
		}

		if r.Method == http.MethodPost {
			// Read the "message" body
			var buffer []byte
			r.Body.Read(buffer)
			message := string(buffer)

			// Read the "who" header
			who := r.Header["who"][0]

			addTwoot(w, message, who)
			return
		}

		// We don't handle that verb.
		http.NotFound(w, r)
	})

	err := http.ListenAndServe(":"+ServerPort, nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %s\n", err.Error())
	}
}

func addTwoot(w http.ResponseWriter, message, who string) {
	twoot := &Twoot{
		message,
		time.Now(),
		who,
	}

	if len(twoots) >= 10 {
		twoots = append(twoots[len(twoots)-9:], twoot)
	} else {
		twoots = append(twoots, twoot)
	}
}

func readTwoots(w http.ResponseWriter) {
	b, err := json.Marshal(&twoots)
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("Short and stout"))
		return
	}

	w.Write(b)
}
