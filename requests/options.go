package requests

import "strings"

type RequestOptions struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []string
}
type FlagHandler func(i *int, args []string, req *RequestOptions)

type FlagDef struct {
	HasValue bool
	Handler  FlagHandler
}

var Flags = map[string]FlagDef{
	"-H": {true, handlerHeader},
	"-d": {true, handlerBody},
	"-I": {false, handlerIncludeHeaders},
}

func handlerHeader(i *int, args []string, req *RequestOptions) {
	x := strings.SplitN(args[*i+1], ":", 2)
	if len(x) != 2 {
		return
	}
	if req.Headers == nil {
		req.Headers = map[string]string{}
	}
	req.Headers[x[0]] = strings.TrimSpace(x[1])
	*i++
}

func handlerBody(i *int, args []string, req *RequestOptions) {
	req.Body = append(req.Body, args[*i+1])
	*i++
}

func handlerIncludeHeaders(i *int, args []string, req *RequestOptions) {

}
