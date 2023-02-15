package go_date

import (
	"fmt"
)

type Date struct {
}

func (d Date) Now() {
	fmt.Println("this is now")
}
