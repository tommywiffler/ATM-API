package responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type Error struct {
	Error string `json:"error"`
}

// Errors will be stored here for reference
var ErrCatalogNotFound = errors.New("catalog not found")

// JSON converts data object to JSON and encodes onto http.responseWriter
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	//log.Info("response.JSON returned " + strconv.Itoa(statusCode))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		//log.Error(err.Error())
		log.Error("panic", zap.String("stack", string(debug.Stack())))
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR converts error into JSON format
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		//log.Error(err.Error())
		log.Error("panic", zap.String("stack", string(debug.Stack())))
		JSON(w, statusCode, Error{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
