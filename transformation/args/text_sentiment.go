package args

// TextSentimentArgs args for TextSentiment transformation.
type TextSentimentArgs struct {
	text     string
	language *string
}

// NewTextSentimentArgs constructor.
func NewTextSentimentArgs(text string) *TextSentimentArgs {
	return &TextSentimentArgs{
		text: text,
	}
}

// SetLanguage setter.
func (a *TextSentimentArgs) SetLanguage(language string) *TextSentimentArgs {
	a.language = &language
	return a
}

// ToMap converts this data to a map.
func (a *TextSentimentArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	args["text"] = a.text

	if a.language != nil {
		args["language"] = a.language
	}

	return args
}
