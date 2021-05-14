package enums

type Status string

var (
	Active  Status = "Active"
	Passive Status = "Passive"
	Deleted Status = "Deleted"
)

type ActionType string

var (
	Create ActionType = "Create"
	Delete ActionType = "Delete"
	Update ActionType = "Update"
)
