package args

import "encoding/json"

// AVConvertOptions aggregates options for AVConvert method.
type AVConvertOptions struct {
	Width            int    `json:"width,omitempty"`
	Height           int    `json:"height,omitempty"`
	Preset           string `json:"preset,omitempty"`
	Force            bool   `json:"force,omitempty"`
	Title            string `json:"title,omitempty"`
	Extname          string `json:"extname,omitempty"`
	Upscale          bool   `json:"upscale,omitempty"`
	AspectMode       string `json:"aspect_mode,omitempty"`
	AudioSampleRate  int    `json:"audio_sample_rate,omitempty"`
	TwoPass          bool   `json:"two_pass,omitempty"`
	VideoBitrate     int    `json:"video_bitrate,omitempty"`
	FPS              int    `json:"fps,omitempty"`
	KeyframeInterval int    `json:"keyframe_interval,omitempty"`
	AudioBitrate     int    `json:"audio_bitrate,omitempty"`
	AudioChannels    int    `json:"audio_channels,omitempty"`
	ClipLength       string `json:"clip_length,omitempty"`
	ClipOffset       string `json:"clip_offset,omitempty"`
	WatermarkURL     string `json:"watermark_url,omitempty"`
	WatermarkTop     int    `json:"watermark_top,omitempty"`
	WatermarkRight   int    `json:"watermark_right,omitempty"`
	WatermarkBottom  int    `json:"watermark_bottom,omitempty"`
	WatermarkLeft    int    `json:"watermark_left,omitempty"`
	FrameCount       int    `json:"frame_count,omitempty"`
	Filename         string `json:"filename,omitempty"`
	Location         string `json:"location,omitempty"`
	Path             string `json:"path,omitempty"`
	Container        string `json:"container,omitempty"`
	Access           string `json:"access,omitempty"`
}

// ToMap creates a map of non-zero option values.
func (a *AVConvertOptions) ToMap() map[string]interface{} {
	var result map[string]interface{}

	bytes, _ := json.Marshal(a)
	json.Unmarshal(bytes, &result)

	return result
}
