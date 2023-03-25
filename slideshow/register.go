package slideshow

func RegisterAll(tmplStore *map[string]string) {

	(*tmplStore)["slideshows/1"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// The directory containing the pictures for a slideshow. It must be stored in the working directory
// of videos229.
// All pictures here must be of width 1366px and height 768px
pictures_dir:

// video_width is the width of the output video in int
video_width: videoWidth

// video_height is the height of the output video in width
video_height: videoHeight

// video_length is the length of the output video in this format (mm:ss)
video_length:
	`

	(*tmplStore)["slideshows/2"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// The directory containing the pictures for a slideshow. It must be stored in the working directory
// of videos229.
// All pictures here must be of width 1366px and height 768px
pictures_dir:

// video_width is the width of the output video in int
video_width: videoWidth

// video_height is the height of the output video in width
video_height: videoHeight

// video_length is the length of the output video in this format (mm:ss)
video_length:
		`
}
