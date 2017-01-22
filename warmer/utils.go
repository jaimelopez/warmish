package warmer

import (
	"log"
	"net/http"
)

func Purge(url string) error {
	log.Println("Purging " + url)

	request, error := http.NewRequest("PURGE", url, nil)

	if error != nil {
		return error
	}

	_, error = http.DefaultClient.Do(request)

	return error
}

func Warmup(url string) error {
	log.Println("Warming " + url)

	_, error := http.Get(url)

	return error
}
