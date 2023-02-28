package effis

// Represents a file stored on Effis.
type FileData struct {
	// The id of the file.
	Id string `json:"id"`
	// The name of the file.
	Name string `json:"name"`
	// The bucket to which the file belongs.
	Bucket string `json:"bucket"`
	// Whether this file is spoiler tagged. This is not provided if false.
	Spoiler bool `json:"spoiler"`
	// A `FileMetadata` object containing meta-information about the file.
	Metadata FileMetadata `json:"metadata"`
}

// Text metadata does not contain any special parameters.
type TextFileMetadata struct {
	// The type of file as a lowercase string, always "text".
	Type string `json:"type"`
}

// Image metadata contains the width and height of the image.
type ImageFileMetadata struct {
	// The type of file as a lowercase string, always "image".
	Type string `json:"type"`
	// The width of the image in pixels.
	Width int `json:"width"`
	// The height of the image in pixels.
	Height int `json:"height"`
}

// Video metadata contains the width and height of the video.
type VideoFileMetadata struct {
	// The type of file as a lowercase string, always "video".
	Type string `json:"type"`
	// The width of the video in pixels.
	Width int `json:"width"`
	// The height of the video in pixels.
	Height int `json:"height"`
}

// Metadata for other files only include the file type.
type OtherFileMetadata struct {
	// The type of file as a lowercase string.
	Type string `json:"type"`
}

// Holds file type-dependent information of a file stored on Effis. This can be one of four variants:
// 	- `TextFileMetadata`
// 	- `ImageFileMetadata`
// 	- `VideoFileMetadata`
// 	- `OtherFileMetadata`
type FileMetadata struct {
	// The type of file as a lowercase string. This can be one of four variants:
	// 	- "text"
	// 	- "image"
	// 	- "video"
	// 	- "other"
	Type string `json:"type"`
	// The width of the file in pixels. This is not provided if the file is not an image or video.
	Width int `json:"width"`
	// The height of the file in pixels. This is not provided if the file is not an image or video.
	Height int `json:"height"`
}
