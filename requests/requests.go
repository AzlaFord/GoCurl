package requests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MethodsType struct {
	Method string
	Action func(req RequestOptions)
}

func GetMethodsActions() []MethodsType {
	return []MethodsType{
		{"GET", GetRequest},
		{"POST", PostRequest},
		{"DELETE", DeleteRequest},
	}
}

func GetRequest(req RequestOptions) {
	res, err := http.Get(req.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s %6s", res.Status, body)
}

func PostRequest(req RequestOptions) {
	combinedBody := strings.Join(req.Body, " ")
	bodyReader := strings.NewReader(combinedBody)
	request, err := http.NewRequest(req.Method, req.URL, bodyReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Printf("%s %6s", response.Status, body)
}

func DeleteRequest(req RequestOptions) {
	request, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s %6s", response.Status, body)
}
