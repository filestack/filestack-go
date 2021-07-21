package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVideoPlaylistArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewVideoPlaylistArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with language", func(t *testing.T) {

		width := 100
		height := 200
		preset := "h264"
		force := true
		title := "title"
		extname := "extname"
		upscale := true
		aspectMode := "letterbox"
		audioSampleRate := 192
		twoPass := true
		videoBitrate := 1024
		fps := 25
		keyframeInterval := 250
		audioBitrate := 128
		audioChannels := 2
		clipLength := "02:14:00"
		clipOffset := "00:00:15"
		watermarkURL := "http://127.0.0.1/watermark.jpg"
		watermarkTop := 10
		watermarkRight := 20
		watermarkBottom := 30
		watermarkLeft := 40
		frameCount := 10
		filename := "myvideo.mp4"
		location := "s3"
		path := "/home/sdk"
		container := "containerName"
		access := "private"

		mapExpected := map[string]interface{}{
			"width":             width,
			"height":            height,
			"preset":            preset,
			"force":             force,
			"title":             title,
			"extname":           extname,
			"upscale":           upscale,
			"aspect_mode":       aspectMode,
			"audio_sample_rate": audioSampleRate,
			"two_pass":          twoPass,
			"video_bitrate":     videoBitrate,
			"fps":               fps,
			"keyframe_interval": keyframeInterval,
			"audio_bitrate":     audioBitrate,
			"audio_channels":    audioChannels,
			"clip_length":       clipLength,
			"clip_offset":       clipOffset,
			"watermark_url":     watermarkURL,
			"watermark_top":     watermarkTop,
			"watermark_right":   watermarkRight,
			"watermark_bottom":  watermarkBottom,
			"watermark_left":    watermarkLeft,
			"frame_count":       frameCount,
			"filename":          filename,
			"location":          location,
			"path":              path,
			"container":         container,
			"access":            access,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewVideoPlaylistArgs()
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetPreset(preset)
		args.SetForce(force)
		args.SetTitle(title)
		args.SetExtName(extname)
		args.SetUpscale(upscale)
		args.SetAspectMode(aspectMode)
		args.SetAudioSampleRate(audioSampleRate)
		args.SetTwoPass(twoPass)
		args.SetVideoBitrate(videoBitrate)
		args.SetFPS(fps)
		args.SetKeyframeInterval(keyframeInterval)
		args.SetAudioBitrate(audioBitrate)
		args.SetAudioChannels(audioChannels)
		args.SetClipLength(clipLength)
		args.SetClipOffset(clipOffset)
		args.SetWatermarkURL(watermarkURL)
		args.SetWatermarkTop(watermarkTop)
		args.SetWatermarkRight(watermarkRight)
		args.SetWatermarkBottom(watermarkBottom)
		args.SetWatermarkLeft(watermarkLeft)
		args.SetFrameCount(frameCount)
		args.SetFilename(filename)
		args.SetLocation(location)
		args.SetPath(path)
		args.SetContainer(container)
		args.SetAccess(access)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
