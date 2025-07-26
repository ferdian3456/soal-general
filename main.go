package main

import (
	"context"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"soal-general/config"
	"soal-general/helper"
	"soal-general/model"
)

func main() {
	router := httprouter.New()
	router.POST("/api/customer/add", AddCustomer)
	router.GET("/api/customer/get-all", GetAllCustomer)

	server := http.Server{
		Addr:    "localhost:8095",
		Handler: router,
	}

	log.Println("Server is running on :", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to run http server")
	}
}

func AddCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	payload := model.CustomerAddRequest{}
	helper.ReadFromRequestBody(request, &payload)

	ctx := request.Context()

	AddCustomerToDB(ctx, payload)

	helper.WriteSuccessResponse(writer)

}

func AddCustomerToDB(ctx context.Context, payload model.CustomerAddRequest) {
	db := config.NewPostgresqlConn()

	query := "INSERT INTO users (name,email,phone) VALUES ($1,$2,$3)"
	_, err := db.Exec(ctx, query, payload.Name, payload.Email, payload.Phone)
	if err != nil {
		log.Fatal("Failed to query into database")
	}
}

func GetAllCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()

	response, err := GetAllCustomerFromDB(ctx)
	if err != nil {
		helper.WriteErrorResponse(writer, http.StatusNotFound, err)
		return
	}

	helper.WriteSuccessResponseWithData(writer, response)
}

func GetAllCustomerFromDB(ctx context.Context) ([]model.CustomerResponse, error) {
	db := config.NewPostgresqlConn()
	query := "SELECT id,name,email,phone FROM users"
	rows, err := db.Query(ctx, query)
	if err != nil {
		log.Fatalf("Failed to query into database :%v", err)
	}

	defer rows.Close()

	var customers []model.CustomerResponse
	hasData := false

	for rows.Next() {
		var customer model.CustomerResponse
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone)
		if err != nil {
			log.Fatal("Failed to scan query result")
		}
		hasData = true
		customers = append(customers, customer)
	}

	if hasData == false {
		return customers, errors.New("customer not exist")
	}

	return customers, nil
}
