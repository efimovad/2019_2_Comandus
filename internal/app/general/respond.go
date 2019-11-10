package general

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	// TODO: fix logger
	/*ctx := r.Context()
	reqID := ctx.Value("rIDKey").(string)
	s.logger.Infof("Request ID: %s | error : %s", reqID , err.Error())*/
	Respond(w, r, code, map[string]string{"error": errors.Cause(err).Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}