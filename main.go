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
	for i := 3; i < len(os.Args); i++ {
		if args[i] == "-H" {
			x := strings.SplitN(args[i+1], ":", 2)
			structRequest.Headers[x[0]] = strings.TrimSpace(x[1])
			i++
		}
		if args[i] == "-D" || args[i] == "-d" {
			structRequest.Body = append(structRequest.Body, args[i+1])
			i++
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
