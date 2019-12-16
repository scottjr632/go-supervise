package entities

// CheckUp is a record of a checkup from a worker
type CheckUp struct {
	Worker       *Worker `json:"worker"`
	ActualResult string  `json:"actualResponse"`
	ResponseCode string  `json:"responseCode"`
}
