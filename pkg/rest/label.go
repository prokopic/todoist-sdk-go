package rest

type Label struct {
	// Id is Label ID.
	Id int64 `json:"id"`

	// Name is Label name.
	Name string `json:"name"`

	// Color is a numeric ID representing the color of the label icon.
	// Refer to the id column in the Colors guide for more info.
	Color int `json:"color"`

	// Order is a number used by clients to sort list of labels.
	Order int `json:"order"`

	// Favorite indicates whether the label is a favorite.
	Favorite bool `json:"favorite"`
}

const (
	LabelsUrl = APIUrl + "/labels"
)
