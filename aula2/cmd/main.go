package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string
	Data    any
}

func main() {
	logger := log.Default()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			logger.Print("Receive Request from '/' and GET method\n")
			response := Response{Message: "OK", Data: "Data"}
			logger.Print(response)
			//  Transform in a binary of our struct
			json, err := json.Marshal(response)

			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			}
			logger.Print(fmt.Sprint("responseJson:", json), "\n")

			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	})
	http.ListenAndServe("localhost:8082", nil) // Create a server
}

// RESPONSE
// {
//     "Message": "OK",
//     "Data": "Data"
// }
