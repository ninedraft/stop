package stop_test

import (
	"fmt"

	"github.com/ninedraft/stop"
)

func ExampleStatus() {
	defer fmt.Println("never be printed")
	defer stop.Exit()
	defer fmt.Println("will be printed")

	stop.Stop(1)
}
