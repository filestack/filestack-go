package payload

// AudioVisualInitResponse stores the response of AudioVisualInit method.
type AudioVisualInitResponse struct {
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
	Status string `json:"status"`
}

// AudioVisualStatusResponse stores the response of AudioVisualStatus method.
type AudioVisualStatusResponse struct {
	Status    string `json:"status"`
	UUID      string `json:"uuid"`
	Timestamp string `json:"timestamp"`
}
