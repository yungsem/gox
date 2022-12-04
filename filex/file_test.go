package filex

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	file, err := os.Open("eqps")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := ReadLines(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", lines)
}
