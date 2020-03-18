package testFile

import (
	"fmt"
	"testing"
	"time"
)

func Testgroutine(t testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("22222", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}