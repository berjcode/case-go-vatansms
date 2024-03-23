package dtos

import "time"

type PlanStatusUpdateDto struct {
	ID        uint
	Name      string
	UpdatedOn *time.Time
	UpdatedBy string
}
