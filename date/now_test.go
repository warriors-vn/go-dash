package date

import (
	"fmt"
	"testing"
)

func Test_now(t *testing.T) {
	result := now()
	fmt.Println(result)
}
