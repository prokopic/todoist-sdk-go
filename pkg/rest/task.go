package rest

type Task struct {
	Assignee     int     `json:"assignee"`
	Assigner     int     `json:"assigner"`
	CommentCount int     `json:"comment_count"`
	Completed    bool    `json:"completed"`
	Content      string  `json:"content"`
	Description  string  `json:"description"`
	Due          DueDate `json:"due"`
	Id           int64   `json:"id"`
	LabelIds     []int64 `json:"label_ids"`
	Order        int     `json:"order"`
	Priority     int     `json:"priority"`
	ProjectId    int64   `json:"project_id"`
	SectionId    int     `json:"section_id"`
	ParentId     int64   `json:"parent_id"`
	Url          string  `json:"url"`
}
