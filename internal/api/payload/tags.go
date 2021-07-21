package payload

// TagsResponse stores response of Tags method.
type TagsResponse struct {
	Tags struct {
		Auto map[string]int `json:"auto"`
		User string         `json:"user"`
	} `json:"tags"`
}
