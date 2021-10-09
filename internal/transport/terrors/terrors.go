package terrors

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Success   bool      `json:"success"`
	ErrorBody errorBody `json:"error"`
}

type errorBody struct {
	Code           int    `json:"code,omitempty"`
	Description    string `json:"description,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (e errorBody) Error() string {
	return e.Description
}

func GetErrEncoderFunc(logger *log.Logger) func(ctx context.Context, err error, w http.ResponseWriter) {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		if err == nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := errorResponse{Success: false}
		typedErr, ok := err.(errorBody)
		if ok {
			response.ErrorBody = typedErr
			w.WriteHeader(typedErr.HttpStatusCode)
		} else {
			response.ErrorBody.Description = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Println("FATAL write to http.ResponseWriter: ", err)
		}
	}
}

func MakeBadRequestErr(description string) error {
	return errorBody{
		Description:    description,
		HttpStatusCode: http.StatusBadRequest,
	}
}

func MakeBadRequestErrWithCode(description string, code int) error {
	err := MakeBadRequestErr(description).(errorBody)
	err.Code = code
	return err
}

func MakeForbiddenErr(description string) error {
	return errorBody{
		Description:    description,
		HttpStatusCode: http.StatusForbidden,
	}
}

func MakeForbiddenErrWithCode(description string, code int) error {
	err := MakeForbiddenErr(description).(errorBody)
	err.Code = code
	return err
}

func MakeNotFoundErr(description string) error {
	return errorBody{
		Description:    description,
		HttpStatusCode: http.StatusNotFound,
	}
}
