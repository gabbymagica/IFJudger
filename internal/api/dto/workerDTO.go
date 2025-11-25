package dto

type ExecutionRequest struct {
	Code       string `json:"code"`
	Input      string `json:"input"`
	WebhookURL string `json:"webhook_url"`
}

type ExecutionResponse struct {
	ID     string
	Status string // "queued", "processing", "success", "error"
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

type ExecutionEnqueuedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
