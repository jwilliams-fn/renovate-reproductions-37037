package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Sprintf("Hello %s", uuid.NewString())
}
