package requests

import (
	"fmt"
	"net/http"
)

type MethodsType struct {
	Method string
	Action func(req RequestOptions)
}

func GetMethodsActions() []MethodsType {
	return []MethodsType{
		{"GET", GetRequest},
	}
}

func GetRequest(req RequestOptions) {
	res, err := http.Get(req.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
