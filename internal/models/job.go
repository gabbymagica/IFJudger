package models

type Job struct {
	ID         string
	Code       string
	Input      string
	WebhookURL string
}

type JobResult struct {
	ID     string
	Status string // "queued", "processing", "success", "error"
	Stdout string
	Stderr string
	Error  string
}
