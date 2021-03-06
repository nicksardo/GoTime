package gotime

import (
	"net/http"
	"strconv"
	"time"
)

func NowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Now()))
}

func Now() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
