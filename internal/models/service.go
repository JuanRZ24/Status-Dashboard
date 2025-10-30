package models

import "time"


type ServiceStatus string


const (
	StatusUnknown ServiceStatus = "UNKNOWN"
	StatusOnline  ServiceStatus = "ONLINE"
	StatusOffline ServiceStatus = "OFFLINE"
	StatusDegraded ServiceStatus = "DEGRADED"
)

type Service struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	URL               string        `json:"url"`
	Method            string        `json:"method"`
	ExpectedStatus    int           `json:"expected_status"`
	ExpectedSubstring string        `json:"expected_substring"`
	TimeoutMS         int           `json:"timeout_ms"`
	IntervalSec       int           `json:"interval_sec"`
	Retries           int           `json:"retries"`
	Enabled           bool          `json:"enabled"`
	Status            ServiceStatus `json:"status"`
	LatencyMS         int           `json:"latency_ms"`
	LastCheckedAt     time.Time     `json:"last_checked_at"`
	Error             string        `json:"error"`
}
