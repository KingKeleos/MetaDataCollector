package server

import (
	"encoding/json"
	"net/http"
)

type AliveMassage struct {
	Name string
	Body string
	Time int64
}

func handleAlive(w http.ResponseWriter, r *http.Request) {
	ans := AliveMassage{"AliveMessage", "true", 69}
	b, err := json.Marshal(ans)
	if err != nil {
		return
	}

	w.Write(b)
}
