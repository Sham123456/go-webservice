package main

import (
	"fmt"

	"github.com/Sham123456/go-webservice.git/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Shamik",
		LastName:  "Mukherjee",
	}
	fmt.Println(u)
}
