package dto

import "time"

type Changes struct {
	UpdatedById ChangesFlowId   `json:"updated_by_id"`
	UpdatedAt   ChangesFlowDate `json:"updated_at"`
	Labels      []Label         `json:"labels"`
}

type ChangesFlowId struct {
	Previous int64 `json:"previous"`
	Current  int64 `json:"current"`
}

type ChangesFlowDate struct {
	Previous time.Time `json:"previous"`
	Current  time.Time `json:"current"`
}
