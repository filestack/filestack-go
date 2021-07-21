package transformation

import "errors"

var (
	ConversionNotCompleted = errors.New("conversion is not completed")
	ParsingURLFailed       = errors.New("parsing url has failed")
)
