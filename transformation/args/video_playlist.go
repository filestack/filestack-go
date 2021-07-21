package args

// VideoPlaylistArgs args for VideoPlaylist transformation.
type VideoPlaylistArgs struct {
	width            *int
	height           *int
	preset           *string
	force            *bool
	title            *string
	extname          *string
	upscale          *bool
	aspectMode       *string
	audioSampleRate  *int
	twoPass          *bool
	videoBitrate     *int
	FPS              *int
	keyframeInterval *int
	audioBitrate     *int
	audioChannels    *int
	clipLength       *string
	clipOffset       *string
	watermarkURL     *string
	watermarkTop     *int
	watermarkRight   *int
	watermarkBottom  *int
	watermarkLeft    *int
	frameCount       *int
	filename         *string
	location         *string
	path             *string
	container        *string
	access           *string
}

// NewVideoPlaylistArgs constructor.
func NewVideoPlaylistArgs() *VideoPlaylistArgs {
	return &VideoPlaylistArgs{}
}

// SetWidth setter.
func (a *VideoPlaylistArgs) SetWidth(width int) *VideoPlaylistArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *VideoPlaylistArgs) SetHeight(height int) *VideoPlaylistArgs {
	a.height = &height
	return a
}

// SetPreset setter.
func (a *VideoPlaylistArgs) SetPreset(preset string) *VideoPlaylistArgs {
	a.preset = &preset
	return a
}

// SetForce setter.
func (a *VideoPlaylistArgs) SetForce(force bool) *VideoPlaylistArgs {
	a.force = &force
	return a
}

// SetTitle setter.
func (a *VideoPlaylistArgs) SetTitle(title string) *VideoPlaylistArgs {
	a.title = &title
	return a
}

// SetExtName setter.
func (a *VideoPlaylistArgs) SetExtName(extname string) *VideoPlaylistArgs {
	a.extname = &extname
	return a
}

// SetUpscale setter.
func (a *VideoPlaylistArgs) SetUpscale(upscale bool) *VideoPlaylistArgs {
	a.upscale = &upscale
	return a
}

// SetAspectMode setter.
func (a *VideoPlaylistArgs) SetAspectMode(aspectMode string) *VideoPlaylistArgs {
	a.aspectMode = &aspectMode
	return a
}

// SetAudioSampleRate setter.
func (a *VideoPlaylistArgs) SetAudioSampleRate(audioSampleRate int) *VideoPlaylistArgs {
	a.audioSampleRate = &audioSampleRate
	return a
}

// SetTwoPass setter.
func (a *VideoPlaylistArgs) SetTwoPass(twoPass bool) *VideoPlaylistArgs {
	a.twoPass = &twoPass
	return a
}

// SetVideoBitrate setter.
func (a *VideoPlaylistArgs) SetVideoBitrate(videoBitrate int) *VideoPlaylistArgs {
	a.videoBitrate = &videoBitrate
	return a
}

// SetFPS setter.
func (a *VideoPlaylistArgs) SetFPS(FPS int) *VideoPlaylistArgs {
	a.FPS = &FPS
	return a
}

// SetKeyframeInterval setter.
func (a *VideoPlaylistArgs) SetKeyframeInterval(keyframeInterval int) *VideoPlaylistArgs {
	a.keyframeInterval = &keyframeInterval
	return a
}

// SetAudioBitrate setter.
func (a *VideoPlaylistArgs) SetAudioBitrate(audioBitrate int) *VideoPlaylistArgs {
	a.audioBitrate = &audioBitrate
	return a
}

// SetAudioChannels setter.
func (a *VideoPlaylistArgs) SetAudioChannels(audioChannels int) *VideoPlaylistArgs {
	a.audioChannels = &audioChannels
	return a
}

// SetClipLength setter.
func (a *VideoPlaylistArgs) SetClipLength(clipLength string) *VideoPlaylistArgs {
	a.clipLength = &clipLength
	return a
}

// SetClipOffset setter.
func (a *VideoPlaylistArgs) SetClipOffset(clipOffset string) *VideoPlaylistArgs {
	a.clipOffset = &clipOffset
	return a
}

// SetWatermarkURL setter.
func (a *VideoPlaylistArgs) SetWatermarkURL(watermarkURL string) *VideoPlaylistArgs {
	a.watermarkURL = &watermarkURL
	return a
}

// SetWatermarkTop setter.
func (a *VideoPlaylistArgs) SetWatermarkTop(watermarkTop int) *VideoPlaylistArgs {
	a.watermarkTop = &watermarkTop
	return a
}

// SetWatermarkRight setter.
func (a *VideoPlaylistArgs) SetWatermarkRight(watermarkRight int) *VideoPlaylistArgs {
	a.watermarkRight = &watermarkRight
	return a
}

// SetWatermarkBottom setter.
func (a *VideoPlaylistArgs) SetWatermarkBottom(watermarkBottom int) *VideoPlaylistArgs {
	a.watermarkBottom = &watermarkBottom
	return a
}

// SetWatermarkLeft setter.
func (a *VideoPlaylistArgs) SetWatermarkLeft(watermarkLeft int) *VideoPlaylistArgs {
	a.watermarkLeft = &watermarkLeft
	return a
}

// SetFrameCount setter.
func (a *VideoPlaylistArgs) SetFrameCount(frameCount int) *VideoPlaylistArgs {
	a.frameCount = &frameCount
	return a
}

// SetFilename setter.
func (a *VideoPlaylistArgs) SetFilename(filename string) *VideoPlaylistArgs {
	a.filename = &filename
	return a
}

// SetLocation setter.
func (a *VideoPlaylistArgs) SetLocation(location string) *VideoPlaylistArgs {
	a.location = &location
	return a
}

// SetPath setter.
func (a *VideoPlaylistArgs) SetPath(path string) *VideoPlaylistArgs {
	a.path = &path
	return a
}

// SetContainer setter.
func (a *VideoPlaylistArgs) SetContainer(container string) *VideoPlaylistArgs {
	a.container = &container
	return a
}

// SetAccess setter.
func (a *VideoPlaylistArgs) SetAccess(access string) *VideoPlaylistArgs {
	a.access = &access
	return a
}

// ToMap converts this data to a map.
func (a *VideoPlaylistArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.width != nil {
		args["width"] = a.width
	}

	if a.height != nil {
		args["height"] = a.height
	}

	if a.preset != nil {
		args["preset"] = a.preset
	}

	if a.force != nil {
		args["force"] = a.force
	}

	if a.title != nil {
		args["title"] = a.title
	}

	if a.extname != nil {
		args["extname"] = a.extname
	}

	if a.upscale != nil {
		args["upscale"] = a.upscale
	}

	if a.aspectMode != nil {
		args["aspect_mode"] = a.aspectMode
	}

	if a.audioSampleRate != nil {
		args["audio_sample_rate"] = a.audioSampleRate
	}

	if a.twoPass != nil {
		args["two_pass"] = a.twoPass
	}

	if a.videoBitrate != nil {
		args["video_bitrate"] = a.videoBitrate
	}

	if a.FPS != nil {
		args["fps"] = a.FPS
	}

	if a.keyframeInterval != nil {
		args["keyframe_interval"] = a.keyframeInterval
	}

	if a.audioBitrate != nil {
		args["audio_bitrate"] = a.audioBitrate
	}

	if a.audioChannels != nil {
		args["audio_channels"] = a.audioChannels
	}

	if a.clipLength != nil {
		args["clip_length"] = a.clipLength
	}

	if a.clipOffset != nil {
		args["clip_offset"] = a.clipOffset
	}

	if a.watermarkURL != nil {
		args["watermark_url"] = a.watermarkURL
	}

	if a.watermarkTop != nil {
		args["watermark_top"] = a.watermarkTop
	}

	if a.watermarkRight != nil {
		args["watermark_right"] = a.watermarkRight
	}

	if a.watermarkBottom != nil {
		args["watermark_bottom"] = a.watermarkBottom
	}

	if a.watermarkLeft != nil {
		args["watermark_left"] = a.watermarkLeft
	}

	if a.frameCount != nil {
		args["frame_count"] = a.frameCount
	}

	if a.filename != nil {
		args["filename"] = a.filename
	}

	if a.location != nil {
		args["location"] = a.location
	}

	if a.path != nil {
		args["path"] = a.path
	}

	if a.container != nil {
		args["container"] = a.container
	}

	if a.access != nil {
		args["access"] = a.access
	}

	return args
}
