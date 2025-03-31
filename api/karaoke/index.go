package karaoke

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ryo-kagawa/Vercel/domain"
	"github.com/ryo-kagawa/Vercel/services/karaoke"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	response, log, err := handler(r)
	if log != "" {
		fmt.Println(log)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		if response.Header.HttpStatusCode == 0 {
			response.Header.HttpStatusCode = http.StatusInternalServerError
		}
		response.Write(w)
		return
	}
	if response.Header.HttpStatusCode == 0 {
		response.Header.HttpStatusCode = 200
	}
	response.Write(w)
}

func handler(r *http.Request) (domain.HttpResponse, string, error) {
	karaoke := karaoke.Karaoke{}
	if r.URL.Path == "/karaoke/webhook" {
		return karaoke.Webhook(r)
	}
	return domain.HttpResponse{
		Header: domain.HttpResponseHeader{
			HttpStatusCode: http.StatusNotFound,
			Contents:       []domain.HttpResponseHeaderContent{},
		},
		Body: "",
	}, "", nil
}
