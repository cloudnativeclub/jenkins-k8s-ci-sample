package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Kubernetes！I'm   test dev    from Jenkins CI！")
	fmt.Println("BRANCH:", os.Getenv("branch"))
}
