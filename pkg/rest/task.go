package rest

type Task struct {
	// Assignee is the responsible user ID (if set, and only for shared tasks).
	Assignee *int `json:"assignee,omitempty"`

	// Assigner is the ID of the user who assigned the task. 0 if the task is unassigned.
	Assigner int `json:"assigner"`

	// CommentCount is the number of task comments.
	CommentCount int `json:"comment_count"`

	// Completed is the flag to mark completed tasks.
	Completed bool `json:"completed"`

	// Content is Task content.
	//
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting
	// article in the Help Center.
	Content string `json:"content"`

	// Description is a description for the task.
	//
	// This value may contain markdown-formatted text and hyperlinks. Details
	// on markdown support can be found in the Text Formatting article in the
	// Help Center.
	Description string `json:"description"`

	// Due is an object representing task due date/time (see DueDate for more details).
	Due DueDate `json:"due"`

	// Id is Task ID.
	Id int64 `json:"id"`

	// LabelIds is the array of label IDs, associated with a task.
	LabelIds []int64 `json:"label_ids"`

	// Order is Position under the same parent or project for top-level
	// tasks (read-only).
	Order int `json:"order"`

	// Priority is Task priority from 1 (normal, default value) to 4 (urgent).
	Priority int `json:"priority"`

	// ProjectId is Task's project ID (read-only).
	ProjectId int64 `json:"project_id"`

	// SectionId is ID of section that task belongs to.
	SectionId int `json:"section_id"`

	// ParentId is the ID of parent task (read-only, absent for top-level tasks).
	ParentId *int64 `json:"parent_id,omitempty"`

	// Parent is the ID of parent task (read-only, absent for top-level tasks).
	//
	// Deprecated: Will be removed in the next API version. Use ParentId.
	Parent *int64 `json:"parent,omitempty"`

	// URL to access this task in the Todoist web or mobile applications.
	URL string `json:"url"`
}
