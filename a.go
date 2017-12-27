package a

import (
	"fmt"
	"github.com/carolynvs/deptest-transcons-b"
)

func A() {
	fmt.Println("a did a thing")
	b.B()
}
