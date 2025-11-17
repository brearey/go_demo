package inet

import (
	"demo/app-1/pkg/util"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	HTTP_TIMEOUT_SECONDS = 30 * time.Second // 30 sec
	MAX_BYTES = 10 * 1024 * 1024 // 10 MB
)

func Curl(url string) (string, error) {
	err := addProtocolPrefix(&url)
	if err != nil {
		util.PrintError("Curl", err)
		return "", err
	}

	client := &http.Client{
		Timeout: HTTP_TIMEOUT_SECONDS,
	}
	resp, err := client.Get(url)
	if err != nil {
		util.PrintError("Curl", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, MAX_BYTES))

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("query is not ok")
	}

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CurlPrint(url string) error {
	body, err := Curl(url)
	if err != nil {
		util.PrintError("Curl", err)
		return err
	}
	fmt.Printf("%s\n", body)
	return nil
}

func addProtocolPrefix(url *string) error {
	if url == nil || *url == "" {
		return errors.New("url is empty")
	}
	if !strings.HasPrefix(*url, "http://") && !strings.HasPrefix(*url, "https://") {
		*url = "http://" + *url
	}
	return nil
}
