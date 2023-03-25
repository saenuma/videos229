package slideshow

func RegisterAll(tmplStore *map[string]string) {

	(*tmplStore)["slideshows/1"] = `// The directory containing the pictures for a slideshow. It must be stored in the working directory
// of videos229.
// All pictures here must be of width 1366px and height 768px
pictures_dir:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video seconds
video_length: 10

// switch_frequency is the number of seconds to switch to a new picture
switch_frequency: 15

	`

	(*tmplStore)["slideshows/2"] = `// The directory containing the pictures for a slideshow. It must be stored in the working directory
// of videos229.
// All pictures here must be of width 1366px and height 768px
pictures_dir:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video seconds
video_length: 10

// switch_frequency is the number of seconds to switch to a new picture
switch_frequency: 15

		`
}
