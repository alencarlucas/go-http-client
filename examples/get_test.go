package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	// Initialization

	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	fmt.Println(err)
	fmt.Println(endpoints)
}
