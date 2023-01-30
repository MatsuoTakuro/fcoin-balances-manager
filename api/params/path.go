package params

import "fmt"

type pathParam struct {
	Name  string
	regex string
}

func (p pathParam) Path() string {
	return fmt.Sprintf("%s:%s", p.Name, p.regex)
}

var UserID = &pathParam{
	Name:  "user_id",
	regex: "[0-9]+",
}
