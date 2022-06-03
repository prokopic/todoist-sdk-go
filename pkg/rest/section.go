package rest

type Section struct {
	// Id is a Section id.
	Id int64 `json:"id"`

	// ProjectId is the ID of the project that the section belongs to.
	ProjectId int64 `json:"project_id"`

	// Order is the Section position among other sections from the same project.
	Order int `json:"order"`

	// Name of the Section.
	Name string `json:"name"`
}
