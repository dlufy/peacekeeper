package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_msg"`
}
type ApiResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

func addCorsHeader(res http.ResponseWriter) {

}
func HandleApiResponse(res http.ResponseWriter, data interface{}, err error) {

	var apiResponse ApiResponse
	if err != nil {
		apiResponse.Status.Success = false
		apiResponse.Status.Message = "Operation Failed"
		apiResponse.Status.ErrorMessage = err.Error()
	} else {
		apiResponse.Status.Success = true
		apiResponse.Status.Message = "Operation Successful"
	}

	apiResponse.Data = data
	bytes, err := json.Marshal(apiResponse)
	if err != nil {
		log.Println("Error marshalling api response ", err, "data", data, "apiresponse", apiResponse)
	} else {
		_, err = res.Write(bytes)
		if err != nil {
			log.Println("Error writing response for trigger campaign ", err)
		}
	}
	return
}

type RouteHandler func(res http.ResponseWriter, req *http.Request) (data interface{}, err error)

func AddMiddleware(rh RouteHandler) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		addCorsHeader(res)
		data, err := rh(res, req)
		/**
		  Add condition if response contains files
		*/
		if res.Header().Get("Content-Type") == "text/csv" {
			/**
			if error occurs pass the error in csv file
			*/
			if err != nil {
				result := []byte(error.Error(err))
				res.Write(result)
			}
			return
		}
		HandleApiResponse(res, data, err)
		return
	}
}
