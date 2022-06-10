package main

import (
	"fmt"

	"github.com/Zeldoso17/Email-App/handlers"
)

func main() {
	fmt.Println("Server running on port 8080")
	handlers.Managers()
}