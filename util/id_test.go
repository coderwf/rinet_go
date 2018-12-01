package util

import (
	"fmt"
	"testing"
)

func TestId(t *testing.T){
    fmt.Println(RandId(10))
    fmt.Println(SecureRandId(10))
}
