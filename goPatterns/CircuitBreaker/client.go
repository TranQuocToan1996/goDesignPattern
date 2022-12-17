package circuitbreaker

import (
	"errors"
	"fmt"
	"net/http"
)

func DoReq() error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/ping", port))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errors.New("bad response")
	}

	return nil
}
