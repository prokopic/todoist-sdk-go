package rest

// Project is a structure that represents Todoist projects in Todoist REST API.
type Project struct {
	// Id is Project Id.
	Id int64 `json:"id"`

	// Name is Project name.
	Name string `json:"name"`

	// CommentCount is number of Project comments.
	CommentCount int `json:"comment_count"`

	// Order is Project position under the same parent (read-only).
	Order int `json:"order"`

	// Color is a numeric ID representing the color of the project icon. Refer to the id
	// column in the Color docs for more info.
	Color int `json:"color"`

	// Shared indicates whether the project is shared (read-only, a true or false value).
	Shared bool `json:"shared"`

	// Favorite indicates whether the project is a favorite (a true or false value).
	Favorite bool `json:"favorite"`

	// ParentId is ID of parent Project (read-only, absent for top-level projects).
	ParentId *int `json:"parent_id,omitempty"`

	// Parent is ID of parent project (read-only, absent for top-level projects).
	//
	// Deprecated: Will be removed in the next API version. Use parent_id.
	Parent *int `json:"parent,omitempty"`

	// TeamInbox indicates whether the project is TeamInbox (read-only, true or otherwise
	// this property is not sent).
	TeamInbox bool `json:"team_inbox"`

	// InboxProject indicates whether the project is Inbox (read-only, true or otherwise this property is not sent).
	InboxProject *bool `json:"inbox_project,omitempty"`

	// URL to access this project in the Todoist web or mobile applications.
	URL string `json:"url"`

	// SyncId is identifier to find the match between different copies of shared projects.
	// When you share a project, its copy has a different ID for your collaborators. To find
	// a project in a different account that matches yours, you can use the "sync_id" attribute.
	// For non-shared projects the attribute is set to 0.
	SyncId *int `json:"sync_id,omitempty"`
}
