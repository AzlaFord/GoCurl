package main

import (
	"GoCurl/requests"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage : <METHOD> [Url]")
		return
	}
	structRequest := requests.RequestOptions{}
	structRequest.Method = args[1]
	structRequest.URL = args[2]
	structRequest.Headers = make(map[string]string)
	for i := 3; i < len(args); i++ {
		if def, ok := requests.Flags[args[i]]; ok {
			def.Handler(&i, args, &structRequest)
		}
	}

	if !strings.HasPrefix(structRequest.URL, "https://") && !strings.HasPrefix(structRequest.URL, "http://") {
		structRequest.URL = "http://" + structRequest.URL
	}

	for _, method := range requests.GetMethodsActions() {
		if structRequest.Method == method.Method {
			method.Action(structRequest)
		}
	}
}
