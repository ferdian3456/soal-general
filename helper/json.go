package helper

import (
	"encoding/json"
	"io"
	"net/http"
	"soal-general/model"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		PanicIfError(err)
	}

	err = json.Unmarshal(body, result)
	PanicIfError(err)
}

func WriteSuccessResponse(writer http.ResponseWriter) {
	webResponse := model.WebResponse{
		Message: "Customer added successfully",
	}

	jsonData, err := json.Marshal(webResponse)
	PanicIfError(err)

	_, err = writer.Write(jsonData)
	PanicIfError(err)
}

func WriteSuccessResponseWithData(writer http.ResponseWriter, response interface{}) {
	webResponse := response

	jsonData, err := json.Marshal(webResponse)
	PanicIfError(err)

	_, err = writer.Write(jsonData)
	PanicIfError(err)
}

func WriteErrorResponse(writer http.ResponseWriter, statusCode int, errorResponse error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	webResponse := model.WebResponse{
		Message: errorResponse.Error(),
	}

	jsonData, err := json.Marshal(webResponse)
	PanicIfError(err)

	_, err = writer.Write(jsonData)
	PanicIfError(err)
}
