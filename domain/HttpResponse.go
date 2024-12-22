package domain

import "net/http"

type HttpResponse struct {
	Header HttpResponseHeader
	Body   string
}

func (h HttpResponse) Write(w http.ResponseWriter) {
	h.Header.Write(w)
	w.Write([]byte(h.Body))
}

type HttpResponseHeader struct {
	HttpStatusCode int
	Contents       []HttpResponseHeaderContent
}

func (h HttpResponseHeader) Write(w http.ResponseWriter) {
	for _, content := range h.Contents {
		w.Header().Add(content.Key, content.Value)
	}
	w.WriteHeader(h.HttpStatusCode)
}

type HttpResponseHeaderContent struct {
	Key   string
	Value string
}
