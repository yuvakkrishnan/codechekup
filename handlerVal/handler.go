package handlerval

import "fmt"

type LolTest struct {
	user string
}

func (l LolTest) GetValues() {
	fmt.Println(l.user)
}
