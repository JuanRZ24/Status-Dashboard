package scheduler

import "net/http"

// worker performs a single check for a given URL. Expand with timing, retries, and recording.
func DoCheck(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
