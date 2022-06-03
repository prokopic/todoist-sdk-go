package rest

import "time"

// Comment can be added to a Task or to a Project.
type Comment struct {
	// Content is Comment content. This value may contain markdown-formatted
	// text and hyperlinks. Details on markdown support can be found in the
	// Text Formatting article in the Help Center.
	Content string `json:"content"`

	// Id of the comment
	Id int64 `json:"id"`

	// Posted is date and time when comment was added.
	//
	// Note: REST API returns RFC3339 format in UTC.
	Posted time.Time `json:"posted"`

	// ProjectId is Comment's Project ID (for project comments).
	ProjectId *int64 `json:"project_id,omitempty"`

	// TaskId is Comment's Task ID (for task comments).
	TaskId *int64 `json:"task_id,omitempty"`

	// Attachment file (optional).
	Attachment *Attachment `json:"attachment,omitempty"`
}

// Attachment (e.g. image, audio, video, file, etc.) can be added to a Comment.
type Attachment struct {
	// FileName is the name of the file.
	FileName string `json:"file_name"`

	// FileType is the MIME type (for example text/plain or image/png).
	FileType string `json:"file_type"`

	// FileUrl is the URL where the file is located.
	//
	// Note that Todoist doesn't cache the remote content on their servers and stream
	// or expose files directly from third party resources. In particular this means
	// that you should avoid providing links to non-encrypted (plain HTTP) resources,
	// as exposing these files in Todoist may issue a browser warning.
	FileUrl string `json:"file_url"`

	// ResourceType is the type of the file (for example image, video, audio, file, etc.)
	ResourceType string `json:"resource_type"`
}
