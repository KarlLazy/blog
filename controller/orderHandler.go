package controller

import (
	"blog/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// FindOrder find order
func FindOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != model.HTTPMethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	params := r.URL.Query()

	index, err := strconv.Atoi(params.Get("index"))
	if err != nil {
		http.Error(w, "参数 `index` 必须为整数", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		http.Error(w, "参数 `limit` 必须为整数", http.StatusBadRequest)
		return
	}

	key := params.Get("key")
	if key != model.KeyAllOrder && key != model.KeyComment && key != model.KeyRefund {
		http.Error(w, "key invalid", http.StatusBadRequest)
		return
	}

	ret := Order.FindOrder(index, limit, key)

	retBytes, err := json.Marshal(ret)
	if err != nil {
		log.Printf("mashal data error: %s", err)
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("mashal data error"))
		return
	}
	w.Write(retBytes)
}
