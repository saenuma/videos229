package sprites

func RegisterAll(tmplStore *map[string]string) {

	(*tmplStore)["sprites/1"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video in seconds
video_length: 10

//vary to your tastes
radius: 100

// vary to your tastes (higher faster, lower slower)
increment: 10

	`

	(*tmplStore)["sprites/2"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video in seconds
video_length: 10
	`

	(*tmplStore)["sprites/3"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768


// video_length is the length of the output video in seconds
video_length: 10
	`

	(*tmplStore)["sprites/4"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video in seconds
video_length: 10

// increment is the incrememt to add to each movement in this movement
increment: 10

	`

	(*tmplStore)["sprites/5"] = `// background_color is the color of the background image. Example is #af1382
background_color: #ffffff

// sprite_file. A sprite is a unit of a pattern in imagery.
sprite_file:

// video_width is the width of the output video in int
video_width: 1366

// video_height is the height of the output video in width
video_height: 768

// video_length is the length of the output video in seconds
video_length: 10

// increment is the incrememt to add to each movement in this movement
increment: 10

	`

}
