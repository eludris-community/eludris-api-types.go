package models

type FileMetadataType string

const (
	FileMetadataTypeText  FileMetadataType = "text"
	FileMetadataTypeImage FileMetadataType = "image"
	FileMetadataTypeVideo FileMetadataType = "video"
	FileMetadataTypeOther FileMetadataType = "other"
)

// Represents a file stored on Effis.
type FileData struct {
	// The id of the file.
	Id string `json:"id"`
	// The name of the file.
	Name string `json:"name"`
	// The bucket to which the file belongs.
	Bucket string `json:"bucket"`
	// Whether this file is spoiler tagged.
	Spoiler bool `json:"spoiler"`
	// A `FileMetadata` object containing meta-information about the file.
	Metadata BaseFileMetadata `json:"metadata"`
}

// Holds file type-dependent information of a file stored on Effis. This can be one of four variants:
//   - `TextFileMetadata`
//   - `ImageFileMetadata`
//   - `VideoFileMetadata`
//   - `OtherFileMetadata`
type BaseFileMetadata struct {
	// The type of file as a lowercase string.
	Type FileMetadataType `json:"type"`
}

// Text metadata does not contain any special parameters.
type TextFileMetadata struct {
	BaseFileMetadata
}

// Image metadata contains the width and height of the image.
type ImageFileMetadata struct {
	BaseFileMetadata
	// The width of the image in pixels.
	Width int `json:"width"`
	// The height of the image in pixels.
	Height int `json:"height"`
}

// Video metadata contains the width and height of the video.
type VideoFileMetadata struct {
	BaseFileMetadata
	// The width of the video in pixels.
	Width int `json:"width"`
	// The height of the video in pixels.
	Height int `json:"height"`
}

// Metadata for other files only include the file type.
type OtherFileMetadata struct {
	BaseFileMetadata
}
