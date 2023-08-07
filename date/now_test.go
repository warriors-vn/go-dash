package date

import (
	"fmt"
	"testing"
)

func Test_now(t *testing.T) {
	result := Now()
	fmt.Println(result)
}
