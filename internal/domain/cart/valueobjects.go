package cart

type StatusEnum string

const (
	StatusOpened      StatusEnum = "OPENED"
	StatusDeactivated StatusEnum = "DEACTIVATED"
	StatusLocked      StatusEnum = "LOCKED"
	StatusCompleted   StatusEnum = "COMPLETED"
)
