// Package transformation provides Transformation type
// with a collection of transformation methods.
// More information can be found in the API documentation:
// https://www.filestack.com/docs/api/processing/
package transformation

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/internal/url"
	"github.com/filestack/filestack-go/options"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/transformation/args"
)

// Transformation type represents an application of tasks to resources.
type Transformation struct {
	host      string
	resources []resource.Resource
	apiKey    string
	security  *security.Security
	tasks     []url.Task
}

// NewTransformation constructor creates an instance
// for single-resource transformation.
func NewTransformation(
	res resource.Resource,
	apiKey string,
	options ...options.Transformation,
) (*Transformation, error) {
	var resources []resource.Resource
	if res != nil {
		resources = append(resources, res)
	}

	return NewTransformationMultiResource(resources, apiKey, options...)
}

// MustNewTransformation is an alias for NewTransformation.
// The difference is that the function stops execution once an error occurs.
func MustNewTransformation(
	res resource.Resource,
	apiKey string,
	options ...options.Transformation,
) *Transformation {
	tr, err := NewTransformation(res, apiKey, options...)
	if err != nil {
		log.Fatalf("failed to create a transformation: %v", err)
	}

	return tr
}

// NewTransformationMultiResource constructor creates
// an instance for multiple-resource transformation.
func NewTransformationMultiResource(
	resources []resource.Resource,
	apiKey string,
	options ...options.Transformation,
) (*Transformation, error) {
	config := config.NewTransformationConfig()
	for _, option := range options {
		if err := option(config); err != nil {
			return nil, fmt.Errorf("setting transformation options has failed: %w", err)
		}
	}

	return &Transformation{
		host:      config.BaseURL,
		resources: resources,
		apiKey:    apiKey,
		security:  config.SecurityPolicy,
	}, nil
}

// MustNewTransformationMultiResource is an alias for NewTransformationMultiResource.
// The difference is that the function stops execution once an error occurs.
func MustNewTransformationMultiResource(
	resources []resource.Resource,
	apiKey string,
	options ...options.Transformation,
) *Transformation {
	tr, err := NewTransformationMultiResource(resources, apiKey, options...)
	if err != nil {
		log.Fatalf("failed to create a transformation: %v", err)
	}

	return tr
}

// BuildURL generates URL address.
func (t *Transformation) BuildURL() string {
	var security string
	if t.security != nil {
		security = t.security.AsString()
	}

	return url.Build(
		t.host,
		t.apiKey,
		security,
		t.tasks,
		t.resources,
		false,
	)
}

// BuildURLInBase64 generates URL address using base64 encoding.
func (t *Transformation) BuildURLInBase64() string {
	var security string
	if t.security != nil {
		security = t.security.AsString()
	}

	return url.Build(
		t.host,
		t.apiKey,
		security,
		t.tasks,
		t.resources,
		true,
	)
}

// GetSecurity provides access to security policy.
func (t *Transformation) GetSecurity() *security.Security {
	return t.security
}

// URLScreenshot extends list of tasks with urlscreenshot.
func (t *Transformation) URLScreenshot(args *args.URLScreenshotArgs) *Transformation {
	return t.addTask("urlscreenshot", args.ToMap())
}

// Resize extends list of tasks with resize.
func (t *Transformation) Resize(args *args.ResizeArgs) *Transformation {
	return t.addTask("resize", args.ToMap())
}

// Crop extends list of tasks with crop.
func (t *Transformation) Crop(args *args.CropArgs) *Transformation {
	return t.addTask("crop", args.ToMap())
}

// Rotate extends list of tasks with rotate.
func (t *Transformation) Rotate(args *args.RotateArgs) *Transformation {
	return t.addTask("rotate", args.ToMap())
}

// Flip extends list of tasks with flip.
func (t *Transformation) Flip() *Transformation {
	return t.addTask("flip", nil)
}

// Flop extends list of tasks with flop.
func (t *Transformation) Flop() *Transformation {
	return t.addTask("flop", nil)
}

// Watermark extends list of tasks with watermark.
func (t *Transformation) Watermark(args *args.WatermarkArgs) *Transformation {
	return t.addTask("watermark", args.ToMap())
}

// DetectFaces extends list of tasks with detect_faces.
func (t *Transformation) DetectFaces(args *args.DetectFacesArgs) *Transformation {
	return t.addTask("detect_faces", args.ToMap())
}

// CropFaces extends list of tasks with crop_faces.
func (t *Transformation) CropFaces(args *args.CropFacesArgs) *Transformation {
	return t.addTask("crop_faces", args.ToMap())
}

// PixelateFaces extends list of tasks with pixelate_faces.
func (t *Transformation) PixelateFaces(args *args.PixelateFacesArgs) *Transformation {
	return t.addTask("pixelate_faces", args.ToMap())
}

// RoundedCorners extends list of tasks with rounded_corners.
func (t *Transformation) RoundedCorners(args *args.RoundedCornersArgs) *Transformation {
	return t.addTask("rounded_corners", args.ToMap())
}

// Vignette extends list of tasks with vignette.
func (t *Transformation) Vignette(args *args.VignetteArgs) *Transformation {
	return t.addTask("vignette", args.ToMap())
}

// Polaroid extends list of tasks with polaroid.
func (t *Transformation) Polaroid(args *args.PolaroidArgs) *Transformation {
	return t.addTask("polaroid", args.ToMap())
}

// TornEdges extends list of tasks with torn_edges.
func (t *Transformation) TornEdges(args *args.TornEdgesArgs) *Transformation {
	return t.addTask("torn_edges", args.ToMap())
}

// Shadow extends list of tasks with shadow.
func (t *Transformation) Shadow(args *args.ShadowArgs) *Transformation {
	return t.addTask("shadow", args.ToMap())
}

// Circle extends list of tasks with circle.
func (t *Transformation) Circle(args *args.CircleArgs) *Transformation {
	return t.addTask("circle", args.ToMap())
}

// Border extends list of tasks with border.
func (t *Transformation) Border(args *args.BorderArgs) *Transformation {
	return t.addTask("border", args.ToMap())
}

// Sharpen extends list of tasks with sharpen.
func (t *Transformation) Sharpen(args *args.SharpenArgs) *Transformation {
	return t.addTask("sharpen", args.ToMap())
}

// Blur extends list of tasks with blur.
func (t *Transformation) Blur(args *args.BlurArgs) *Transformation {
	return t.addTask("blur", args.ToMap())
}

// Monochrome extends list of tasks with monochrome.
func (t *Transformation) Monochrome() *Transformation {
	return t.addTask("monochrome", nil)
}

// BlackWhite extends list of tasks with blackwhite.
func (t *Transformation) BlackWhite(args *args.BlackWhiteArgs) *Transformation {
	return t.addTask("blackwhite", args.ToMap())
}

// Sepia extends list of tasks with sepia.
func (t *Transformation) Sepia(args *args.SepiaArgs) *Transformation {
	return t.addTask("sepia", args.ToMap())
}

// Pixelate extends list of tasks with pixelate.
func (t *Transformation) Pixelate(args *args.PixelateArgs) *Transformation {
	return t.addTask("pixelate", args.ToMap())
}

// OilPaint extends list of tasks with old_paint.
func (t *Transformation) OilPaint(args *args.OldPaintArgs) *Transformation {
	return t.addTask("oil_paint", args.ToMap())
}

// Negative extends list of tasks with nagative.
func (t *Transformation) Negative() *Transformation {
	return t.addTask("negative", nil)
}

// Modulate extends list of tasks with modulate.
func (t *Transformation) Modulate(args *args.ModulateArgs) *Transformation {
	return t.addTask("modulate", args.ToMap())
}

// PartialPixelate extends list of tasks with partial_pixelate.
func (t *Transformation) PartialPixelate(args *args.PartialPixelateArgs) *Transformation {
	return t.addTask("partial_pixelate", args.ToMap())
}

// PartialBlur extends list of tasks with partial_blur.
func (t *Transformation) PartialBlur(args *args.PartialBlurArgs) *Transformation {
	return t.addTask("partial_blur", args.ToMap())
}

// Collage extends list of tasks with collage.
func (t *Transformation) Collage(args *args.CollageArgs) *Transformation {
	return t.addTask("collage", args.ToMap())
}

// Upscale extends list of tasks with upscale.
func (t *Transformation) Upscale(args *args.UpscaleArgs) *Transformation {
	return t.addTask("upscale", args.ToMap())
}

// Enhance extends list of tasks with enhance.
func (t *Transformation) Enhance(args *args.EnhanceArgs) *Transformation {
	return t.addTask("enhance", args.ToMap())
}

// RedEye extends list of tasks with redeye.
func (t *Transformation) RedEye() *Transformation {
	return t.addTask("redeye", nil)
}

// ASCII extends list of tasks with ascii.
func (t *Transformation) ASCII(args *args.ASCIIArgs) *Transformation {
	return t.addTask("ascii", args.ToMap())
}

// FiletypeConversion extends list of tasks with output.
func (t *Transformation) FiletypeConversion(args *args.FileTypeConversionArgs) *Transformation {
	return t.addTask("output", args.ToMap())
}

// NoMetadata extends list of tasks with no_metadata.
func (t *Transformation) NoMetadata() *Transformation {
	return t.addTask("no_metadata", nil)
}

// Quality extends list of tasks with quality.
func (t *Transformation) Quality(value int) *Transformation {
	return t.addTask("quality", map[string]interface{}{
		"value": value,
	})
}

// Zip extends list of tasks with zip.
func (t *Transformation) Zip() *Transformation {
	return t.addTask("zip", nil)
}

// Fallback extends list of tasks with fallback.
func (t *Transformation) Fallback(handle string, cache int) *Transformation {
	return t.addTask("fallback", map[string]interface{}{
		"handle": handle,
		"cache":  cache,
	})
}

// PDFInfo extends list of tasks with pdfinfo.
func (t *Transformation) PDFInfo(colorInfo bool) *Transformation {
	return t.addTask("pdfinfo", map[string]interface{}{
		"colorinfo": colorInfo,
	})
}

// PDFConvert extends list of tasks with pdfconvert.
func (t *Transformation) PDFConvert(args *args.PDFConvertArgs) *Transformation {
	return t.addTask("pdfconvert", args.ToMap())
}

// MinifyJS extends list of tasks with minify_js.
func (t *Transformation) MinifyJS(args *args.MinifyJSArgs) *Transformation {
	return t.addTask("minify_js", args.ToMap())
}

// MinifyCSS extends list of tasks with minify_css.
func (t *Transformation) MinifyCSS(args *args.MinifyCSSArgs) *Transformation {
	return t.addTask("minify_css", args.ToMap())
}

// AVConvert extends list of tasks with video_convert.
func (t *Transformation) AVConvert(options args.AVConvertOptions) *Transformation {
	return t.addTask("video_convert", options.ToMap())
}

// AutoImage extends list of tasks with auto_image.
func (t *Transformation) AutoImage() *Transformation {
	return t.addTask("auto_image", nil)
}

// DocDetection extends list of tasks with doc_detection.
func (t *Transformation) DocDetection(args *args.DocDetection) *Transformation {
	return t.addTask("doc_detection", args.ToMap())
}

// ImageSentiment extends list of tasks with image_sentiment.
func (t *Transformation) ImageSentiment() *Transformation {
	return t.addTask("image_sentiment", nil)
}

// TextSentiment extends list of tasks with text_sentiment.
func (t *Transformation) TextSentiment(args *args.TextSentimentArgs) *Transformation {
	return t.addTask("text_sentiment", args.ToMap())
}

// Caption extends list of tasks with caption.
func (t *Transformation) Caption() *Transformation {
	return t.addTask("caption", nil)
}

// OCR extends list of tasks with ocr.
func (t *Transformation) OCR() *Transformation {
	return t.addTask("ocr", nil)
}

// Preview extends list of tasks with preview.
func (t *Transformation) Preview() *Transformation {
	return t.addTask("preview", nil)
}

// ImageSize extends list of tasks with imagesize.
func (t *Transformation) ImageSize() *Transformation {
	return t.addTask("imagesize", nil)
}

// QR extends list of tasks with qr.
func (t *Transformation) QR(args *args.QR) *Transformation {
	return t.addTask("qr", args.ToMap())
}

// Animate extends list of tasks with animate.
func (t *Transformation) Animate(args *args.AnimateArgs) *Transformation {
	return t.addTask("animate", args.ToMap())
}

// PDFCreate extends list of tasks with pdfcreate.
func (t *Transformation) PDFCreate(args *args.PDFCreateArgs) *Transformation {
	return t.addTask("pdfcreate", args.ToMap())
}

// DocToImages extends list of tasks with doc_to_images.
func (t *Transformation) DocToImages(args *args.DocToImagesArgs) *Transformation {
	return t.addTask("doc_to_images", args.ToMap())
}

// PJPG extends list of tasks with pjpg.
func (t *Transformation) PJPG(args *args.PJPGArgs) *Transformation {
	return t.addTask("pjpg", args.ToMap())
}

// BlurFaces extends list of tasks with blur_faces.
func (t *Transformation) BlurFaces(args *args.BlurFacesArgs) *Transformation {
	return t.addTask("blur_faces", args.ToMap())
}

// SmartCrop extends list of tasks with smart_crop.
func (t *Transformation) SmartCrop(args *args.SmartCropArgs) *Transformation {
	return t.addTask("smart_crop", args.ToMap())
}

// Slide extends list of tasks with slide.
func (t *Transformation) Slide(args *args.SlideArgs) *Transformation {
	return t.addTask("slide", args.ToMap())
}

// VideoPlaylist extends list of tasks with video_playlist.
func (t *Transformation) VideoPlaylist(args *args.VideoPlaylistArgs) *Transformation {
	return t.addTask("video_playlist", args.ToMap())
}

// Store extends list of tasks with store.
func (t *Transformation) Store(args *args.StoreArgs) *Transformation {
	return t.addTask("store", args.ToMap())
}

// Metadata extends list of tasks with metadata.
func (t *Transformation) Metadata() *Transformation {
	return t.addTask("metadata", nil)
}

// Tags extends list of tasks with tags.
func (t *Transformation) Tags() *Transformation {
	return t.addTask("tags", nil)
}

// Sfw extends list of tasks with sfw.
func (t *Transformation) Sfw() *Transformation {
	return t.addTask("sfw", nil)
}

// RunWorkflow extends list of tasks with run_workflow.
func (t *Transformation) RunWorkflow(args *args.RunWorkflowArgs) *Transformation {
	return t.addTask("run_workflow", args.ToMap())
}

// WorkflowStatus extends list of tasks with workflow_status.
func (t *Transformation) WorkflowStatus(args *args.WorkflowStatusArgs) *Transformation {
	return t.addTask("workflow_status", args.ToMap())
}

func (t *Transformation) addTask(name string, params map[string]interface{}) *Transformation {
	t.tasks = append(t.tasks, url.NewTask(name, params))
	return t
}
