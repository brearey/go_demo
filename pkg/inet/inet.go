package inet

import (
	"demo/app-1/pkg/util"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Curl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		util.PrintError("Curl", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if (resp.StatusCode != http.StatusOK) {
		return "", errors.New("query is not ok")
	}

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CurlPrint(url string) error {
	body, err := Curl(url)
	if (err != nil) {
		util.PrintError("Curl", err)
	}
	fmt.Printf("%s\n", body)
	return nil
}