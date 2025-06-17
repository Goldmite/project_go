package enums

type Status string

const (
	Pending  Status = "pending"
	Accepted Status = "accepted"
	Declined Status = "declined"
	Expired  Status = "expired"
)
