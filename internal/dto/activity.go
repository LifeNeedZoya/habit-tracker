package dto

type ActivityDTO struct {
	Name          string             `json:"name"`
	SessionNumber int                `json:"session_number"`
	StartTime     string             `json:"start_time"`
	TotalDuration int                `json:"total_duration"`
	Sessions      []CreateSessionDTO `json:"sessions"`
}
