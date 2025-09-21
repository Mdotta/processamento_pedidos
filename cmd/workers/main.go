package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"processamento_pedidos/internal/models"
)

func main() {
	req := models.CreateUserRequest{
		Username: "novo usuário",
		Email:    "email@novo.com",
	}
	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		var responseAPI models.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&responseAPI); err != nil {
			panic(err)
		}
		panic(responseAPI.Reason)
	}
	var responseAPI models.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseAPI); err != nil {
		panic(err)
	}
	fmt.Println("new user created:", responseAPI)
}
