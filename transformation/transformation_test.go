package transformation

import (
	"testing"
	"time"

	transformationOptions "github.com/filestack/filestack-go/options/transformation"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/transformation/args"
	"github.com/stretchr/testify/assert"
)

const DefaultHandle = "rFVGFv0RYa6d9O7Shncr"
const DefaultAPIKey = "A5PbLMnPT1GMlY1OyoPODz"
const cdnHostURL = "https://cdn.filestackcontent.com/"

func TestTransformation_BuildURL(t *testing.T) {

	security := security.NewSecurity("MBVEKA6ARFGNDDTVY7IQDZ4HYU", &security.Policy{
		Expiry: time.Now().Add(time.Duration(24 * time.Hour)).Unix(),
	})
	handleResource := resource.NewHandle(DefaultHandle)
	apiKey := DefaultAPIKey
	baseURL := cdnHostURL + apiKey

	t.Run("resize", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.addTask("resize", map[string]interface{}{
			"width":  "640",
			"height": "480",
		})
		result := transformation.BuildURL()
		expected := baseURL + "/resize=height:480,width:640/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("resizeWithSecurity", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey, transformationOptions.SecurityPolicy(security))
		transformation.addTask("resize", map[string]interface{}{
			"width":  "640",
			"height": "480",
		})
		result := transformation.BuildURL()
		expected := baseURL + "/" + security.AsString() + "/resize=height:480,width:640/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("shadow", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.addTask("shadow", map[string]interface{}{
			"opacity": "50",
			"blur":    "5",
			"vector":  []string{"5", "10"},
		})
		result := transformation.BuildURL()
		expected := baseURL + "/shadow=blur:5,opacity:50,vector:[5,10]/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("crop_faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.CropFaces(
			args.NewCropFacesArgs().
				SetBuffer(0).
				SetFacesAll().
				SetWidth(100).
				SetHeight(100).
				SetMaxSize(0.35).
				SetMinSize(0.35).
				SetMode("thumb"),
		)
		result := transformation.BuildURL()
		expected := baseURL + "/crop_faces=buffer:0,faces:all,height:100,maxsize:0.35,minsize:0.35,mode:thumb,width:100/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pixelate_faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PixelateFaces(
			args.NewPixelateFacesArgs().
				SetAmount(10).
				SetBlur(4).
				SetBuffer(10).
				SetFacesAll().
				SetMaxSize(0.35).
				SetMinSize(0.35).
				SetType("rect"),
		)
		result := transformation.BuildURL()
		expected := baseURL + "/pixelate_faces=amount:10,blur:4,buffer:10,faces:all,maxsize:0.35,minsize:0.35,type:rect/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("rounded_corners", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.RoundedCorners(args.NewRoundedCornersArgs().SetRadiusMax().SetBlur(0.3).SetBackground("ff0000"))
		result := transformation.BuildURL()
		expected := baseURL + "/rounded_corners=background:ff0000,blur:0.3,radius:max/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("vignette", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Vignette(args.NewVignetteArgs().SetAmount(20).SetBackground("ffccff").SetBlurMode("gaussian"))
		result := transformation.BuildURL()
		expected := baseURL + "/vignette=amount:20,background:ffccff,blurmode:gaussian/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("polaroid", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Polaroid(args.NewPolaroidArgs().SetBackground("ff0000").SetColor("00ff00").SetRotate(90))
		result := transformation.BuildURL()
		expected := baseURL + "/polaroid=background:ff0000,color:00ff00,rotate:90/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("torn_edges", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.TornEdges(args.NewTornEdgesArgs().SetBackground("ff0000").SetSpread(1, 10))
		result := transformation.BuildURL()
		expected := baseURL + "/torn_edges=background:ff0000,spread:[1,10]/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("shadow", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Shadow(args.NewShadowArgs().SetBackground("00ff00").SetBlur(4).SetColor("ff0000").SetOpacity(60).SetVector(4, 4))
		result := transformation.BuildURL()
		expected := baseURL + "/shadow=background:00ff00,blur:4,color:ff0000,opacity:60,vector:[4,4]/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("circle", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Circle(args.NewCircleArgs().SetBackground("00ff00"))
		result := transformation.BuildURL()
		expected := baseURL + "/circle=background:00ff00/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("border", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Border(args.NewBorderArgs().SetBackground("ff0000").SetColor("00ff00").SetWidth(50))
		result := transformation.BuildURL()
		expected := baseURL + "/border=background:ff0000,color:00ff00,width:50/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("sharpen", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Sharpen(args.NewSharpenArgs().SetAmount(10))
		result := transformation.BuildURL()
		expected := baseURL + "/sharpen=amount:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("blur", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Blur(args.NewBlurArgs().SetAmount(10))
		result := transformation.BuildURL()
		expected := baseURL + "/blur=amount:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("monochrome", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Monochrome()
		result := transformation.BuildURL()
		expected := baseURL + "/monochrome/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("blackwhite", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.BlackWhite(args.NewBlackWhiteArgs().SetThreshold(50))
		result := transformation.BuildURL()
		expected := baseURL + "/blackwhite=threshold:50/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("sepia", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Sepia(args.NewSepiaArgs().SetTone(60))
		result := transformation.BuildURL()
		expected := baseURL + "/sepia=tone:60/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pixelate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Pixelate(args.NewPixelateArgs().SetAmount(60))
		result := transformation.BuildURL()
		expected := baseURL + "/pixelate=amount:60/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("oil_paint", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.OilPaint(args.NewOldPaintArgs().SetAmount(10))
		result := transformation.BuildURL()
		expected := baseURL + "/oil_paint=amount:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("negative", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Negative()
		result := transformation.BuildURL()
		expected := baseURL + "/negative/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("modulate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Modulate(args.NewModulateArgs().SetBrightness(10).SetSaturation(20).SetHue(30))
		result := transformation.BuildURL()
		expected := baseURL + "/modulate=brightness:10,hue:30,saturation:20/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("partial_pixelate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		partialPixelateArgs := args.NewPartialPixelateArgs([]args.ImageArea{args.NewImageArea(100, 100, 400, 100)})
		partialPixelateArgs.SetAmount(50)
		partialPixelateArgs.SetBlur(10)
		partialPixelateArgs.SetFilterType("rect")
		transformation.PartialPixelate(partialPixelateArgs)
		result := transformation.BuildURL()

		expected := baseURL + "/partial_pixelate=amount:50,blur:10,objects:[[100,100,400,100]],type:rect/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("partial_blur", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		partialBlurArgs := args.NewPartialBlurArgs([]args.ImageArea{args.NewImageArea(100, 100, 400, 100)})
		transformation.PartialBlur(partialBlurArgs.SetBlur(10).SetAmount(10).SetFilterType("rect"))
		result := transformation.BuildURL()

		expected := baseURL + "/partial_blur=amount:10,blur:10,objects:[[100,100,400,100]],type:rect/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("collage", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		images := []string{
			"aQtNuc6GRVKiZbRnpeoh",
			"JGShhcKrSeyhpmVSE216",
			"Ow3S7Z6dQhe00yoomg8o",
		}
		transformation.Collage(
			args.NewCollageArgs().
				SetColor("000000").
				SetFiles(images).
				SetFit("auto").
				SetHeight(400).
				SetMargin(10).
				SetWidth(600).
				SetAutoRotate(),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/collage=autorotate:true,color:000000,files:[aQtNuc6GRVKiZbRnpeoh,JGShhcKrSeyhpmVSE216,Ow3S7Z6dQhe00yoomg8o],fit:auto,height:400,margin:10,width:600/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("upscale", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Upscale(
			args.NewUpscaleArgs().
				SetNoise("medium").
				SetStyle("artwork").
				SetUpscale(true),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/upscale=noise:medium,style:artwork,upscale:true/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("enhance", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Enhance(args.NewEnhanceArgs().SetPreset("beautify"))
		result := transformation.BuildURL()

		expected := baseURL + "/enhance=preset:beautify/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("redeye", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.RedEye()
		result := transformation.BuildURL()

		expected := baseURL + "/redeye/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("ascii", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.ASCII(args.NewASCIIArgs().
			SetBackground("000000").
			SetColored().
			SetForeground("00cc00").
			SetSize(100),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/ascii=background:000000,colored:true,foreground:00cc00,size:100/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("filetype_conversion", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.FiletypeConversion(args.NewFileTypeConversionArgs().
			SetBackground("000000").
			SetColorSpace("rgb").
			SetCompress(false).
			SetDensity(1).SetFormat("pdf").SetPageFormat("a4").SetPageOrientation("portrait").SetPage(1).SetQualityInput(),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/output=background:000000,colorspace:rgb,compress:false,density:1,format:pdf,page:1,pageformat:a4,pageorientation:portrait,quality:input/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("no_metadata", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.NoMetadata()
		result := transformation.BuildURL()

		expected := baseURL + "/no_metadata/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("metadata", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Metadata()
		result := transformation.BuildURL()

		expected := baseURL + "/metadata/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("tags", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Tags()
		result := transformation.BuildURL()

		expected := baseURL + "/tags/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("sfw", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Sfw()
		result := transformation.BuildURL()

		expected := baseURL + "/sfw/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("run_workflow", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.RunWorkflow(args.NewRunWorkflowArgs().SetID("10"))
		result := transformation.BuildURL()

		expected := baseURL + "/run_workflow=id:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("workflow_status", func(t *testing.T) {
		transformation := MustNewTransformation(nil, apiKey)
		transformation.WorkflowStatus(args.NewWorkflowStatusArgs().SetJobID("10"))
		result := transformation.BuildURL()

		expected := baseURL + "/workflow_status=job_id:10"
		assert.Equal(t, expected, result)
	})

	t.Run("quality", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Quality(5)
		result := transformation.BuildURL()

		expected := baseURL + "/quality=value:5/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("zip", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Zip()
		result := transformation.BuildURL()

		expected := baseURL + "/zip/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("fallback", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Fallback("aQtNuc6GRVKiZbRnpeoh", 3)
		result := transformation.BuildURL()

		expected := baseURL + "/fallback=cache:3,handle:aQtNuc6GRVKiZbRnpeoh/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pdfinfo", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PDFInfo(true)
		result := transformation.BuildURL()

		expected := baseURL + "/pdfinfo=colorinfo:true/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pdfconvert", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PDFConvert(args.NewPDFConvertArgs().
			SetMetadata(true).
			SetPageFormat("a3").
			SetPageOrientation("landscape").
			SetPages("[1,2,3]"),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/pdfconvert=metadata:true,pageformat:a3,pageorientation:landscape,pages:[1,2,3]/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("minify_js", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.MinifyJS(args.NewMinifyJSArgs().
			SetGzip(true).
			SetKeepClassName(true).
			SetKeepFnName(true).
			SetMangle(true).
			SetMergeVars(true).
			SetRemoveConsole(true).
			SetRemoveUndefined(true).
			SetTargets(true).
			SetUseBabelPolyfill(true),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/minify_js=gzip:true,keep_class_name:true,keep_fn_name:true,mangle:true,merge_vars:true,remove_console:true,remove_undefined:true,targets:true,use_babel_polyfill:true/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("minify_css", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.MinifyCSS(args.NewMinifyCSSArgs().SetGzip().SetLevel(5))
		result := transformation.BuildURL()

		expected := baseURL + "/minify_css=gzip:true,level:5/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("av_convert", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.AVConvert(args.AVConvertOptions{Width: 800, Height: 600})
		result := transformation.BuildURL()

		expected := baseURL + "/video_convert=height:600,width:800/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("auto_image", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.AutoImage()
		result := transformation.BuildURL()

		expected := baseURL + "/auto_image/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("doc_detection", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.DocDetection(args.NewDocDetection().SetCoords(true).SetPreprocess(false))
		result := transformation.BuildURL()

		expected := baseURL + "/doc_detection=coords:true,preprocess:false/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("image_sentiment", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.ImageSentiment()
		result := transformation.BuildURL()

		expected := baseURL + "/image_sentiment/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("text_sentiment", func(t *testing.T) {
		text := "this sdk is awesome"
		language := "en"

		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.TextSentiment(args.NewTextSentimentArgs(text).SetLanguage(language))
		result := transformation.BuildURL()

		expected := baseURL + "/text_sentiment=language:" + language + ",text:" + text + "/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("caption", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Caption()
		result := transformation.BuildURL()

		expected := baseURL + "/caption/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("ocr", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.OCR()
		result := transformation.BuildURL()

		expected := baseURL + "/ocr/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("preview", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Preview()
		result := transformation.BuildURL()

		expected := baseURL + "/preview/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("imagesize", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.ImageSize()
		result := transformation.BuildURL()

		expected := baseURL + "/imagesize/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("qr", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.QR(args.NewQR().
			SetVersion(10).
			SetErrorCorrection("H").
			SetFormat("png"),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/qr=error_correction:H,format:png,version:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("animate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Animate(args.NewAnimateArgs().
			SetDelay(10).
			SetLoop(3).
			SetAlign("bottom").
			SetFit("scale").
			SetWidthMax().
			SetHeight(200).
			SetBackground("transparent"),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/animate=align:bottom,background:transparent,delay:10,fit:scale,height:200,loop:3,width:max/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pdfcreate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PDFCreate(args.NewPDFCreateArgs().SetEngine("mupdf"))
		result := transformation.BuildURL()

		expected := baseURL + "/pdfcreate=engine:mupdf/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("doc_to_images", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.DocToImages(
			args.NewDocToImagesArgs().
				SetEngine("imagemagick").
				SetFormat("jpg").
				SetQuality(100).
				SetDensity(72).
				SetHiddenSlides(true).
				SetPages([]string{"1-3", "4-5"}),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/doc_to_images=density:72,engine:imagemagick,format:jpg,hidden_slides:true,pages:[1-3,4-5],quality:100/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pjpg", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PJPG(args.NewPJPGArgs().SetQuality(10).SetMetadata(true))
		result := transformation.BuildURL()

		expected := baseURL + "/pjpg=metadata:true,quality:10/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("blur_faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.BlurFaces(args.NewBlurFacesArgs().
			SetFaces([]int{1, 2, 3}).
			SetBuffer(10).
			SetMaxSize(0.35).
			SetMinSize(0.35).
			SetBlurType("oval").
			SetAmount(0.12),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/blur_faces=amount:0.12,buffer:10,faces:[1,2,3],maxsize:0.35,minsize:0.35,type:oval/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("smart_crop", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.SmartCrop(args.NewSmartCropArgs().
			SetCoords(true).
			SetFillColor("white").
			SetHeight(200).
			SetMode("auto").
			SetWidth(100),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/smart_crop=coords:true,fill_color:white,height:200,mode:auto,width:100/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("slide", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Slide(args.NewSlideArgs().SetEngine("mupdf").SetTheme("dark"))
		result := transformation.BuildURL()

		expected := baseURL + "/slide=engine:mupdf,theme:dark/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("video_playlist", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.VideoPlaylist(
			args.NewVideoPlaylistArgs().
				SetWidth(100).
				SetHeight(200).
				SetPreset("h264").
				SetForce(true).
				SetTitle("myplaylist").
				SetExtName("extname").
				SetUpscale(true).
				SetAspectMode("constrain").
				SetAudioSampleRate(128).
				SetTwoPass(true).
				SetVideoBitrate(128).
				SetFPS(25).
				SetKeyframeInterval(250).
				SetAudioBitrate(128).
				SetAudioChannels(2).
				SetClipLength("02:40:00").
				SetClipOffset("00:01:00").
				SetWatermarkURL("http://127.0.0.1/video.mp4").
				SetWatermarkTop(10).
				SetWatermarkRight(20).
				SetWatermarkBottom(30).
				SetWatermarkLeft(40).
				SetFrameCount(100).
				SetFilename("filename.mp4").
				SetLocation("dropbox").
				SetPath("/home/sdk").
				SetContainer("container").
				SetAccess("public"),
		)
		result := transformation.BuildURL()

		expected := baseURL + "/video_playlist=access:public,aspect_mode:constrain,audio_bitrate:128,audio_channels:2,audio_sample_rate:128,clip_length:02:40:00,clip_offset:00:01:00,container:container,extname:extname,filename:filename.mp4,force:true,fps:25,frame_count:100,height:200,keyframe_interval:250,location:dropbox,path:/home/sdk,preset:h264,title:myplaylist,two_pass:true,upscale:true,video_bitrate:128,watermark_bottom:30,watermark_left:40,watermark_right:20,watermark_top:10,watermark_url:http://127.0.0.1/video.mp4,width:100/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

}

/**
invoke transformations with minimal set of arguments and check BuildURL result
*/
func TestTransformation_ShortestURL(t *testing.T) {

	handleResource := resource.NewHandle(DefaultHandle)
	apiKey := DefaultAPIKey
	baseURL := cdnHostURL + apiKey

	t.Run("urlscreenshot", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.URLScreenshot(args.NewURLScreenshotArgs())
		result := transformation.BuildURL()

		expected := baseURL + "/urlscreenshot/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("resize", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Resize(args.NewResizeArgs().SetWidth(600).SetHeight(400))
		result := transformation.BuildURL()

		expected := baseURL + "/resize=height:400,width:600/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("crop", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Crop(args.NewCropArgs(10, 20, 30, 40))
		result := transformation.BuildURL()

		expected := baseURL + "/crop=height:40,width:30,x:10,y:20/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("rotate", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Rotate(args.NewRotateArgs())
		result := transformation.BuildURL()

		expected := baseURL + "/rotate/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("flip", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Flip()
		result := transformation.BuildURL()

		expected := baseURL + "/flip/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("flop", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Flop()
		result := transformation.BuildURL()

		expected := baseURL + "/flop/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("watermark", func(t *testing.T) {
		fileName := "myfile.txt"
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.Watermark(args.NewWatermarkArgs(fileName))
		result := transformation.BuildURL()

		expected := baseURL + "/watermark=file:" + fileName + "/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("detect faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.DetectFaces(args.NewDetectFacesArgs())
		result := transformation.BuildURL()

		expected := baseURL + "/detect_faces/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("crop faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.CropFaces(args.NewCropFacesArgs())
		result := transformation.BuildURL()

		expected := baseURL + "/crop_faces/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})

	t.Run("pixelate faces", func(t *testing.T) {
		transformation := MustNewTransformation(handleResource, apiKey)
		transformation.PixelateFaces(args.NewPixelateFacesArgs())
		result := transformation.BuildURL()

		expected := baseURL + "/pixelate_faces/" + handleResource.AsString()
		assert.Equal(t, expected, result)
	})
}
