package dto

import "time"

type Label struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Color       string    `json:"color"`
	ProjectID   int64     `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Template    bool      `json:"template"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	GroupID     int64     `json:"group_id"`
}
