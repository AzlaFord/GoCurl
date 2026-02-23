package requests

type RequestOptions struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []string
}
