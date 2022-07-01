package converters

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	gitlab "github.com/xanzy/go-gitlab"
)

type MRConverter struct {
	MergeEvent *gitlab.MergeEvent
}

func NewMRConverter(mr *gitlab.MergeEvent) MRConverter {
	return MRConverter{
		MergeEvent: mr,
	}
}

func (c *MRConverter) Convert() (interface{}, error) {
	labels := make([]string, 0)
	for _, l := range c.MergeEvent.Labels {
		labels = append(labels, l.Name)
	}
	var assignee dto.Person
	if c.MergeEvent.Assignee != nil {
		assignee = dto.Person{
			Name:     c.MergeEvent.Assignee.Name,
			UserName: c.MergeEvent.User.Username,
		}
	}
	return &dto.MergeRequest{
		ID:  int64(c.MergeEvent.ObjectAttributes.ID),
		IID: int64(c.MergeEvent.ObjectAttributes.IID),
		Project: dto.Project{
			ID:          int64(c.MergeEvent.Project.ID),
			Name:        c.MergeEvent.Project.Name,
			Description: c.MergeEvent.Project.Description,
		},
		Title:       c.MergeEvent.ObjectAttributes.Title,
		Description: c.MergeEvent.ObjectAttributes.Description,
		State:       c.MergeEvent.ObjectAttributes.State,
		URL:         c.MergeEvent.ObjectAttributes.URL,
		Author: dto.Person{
			Name:     c.MergeEvent.User.Name,
			UserName: c.MergeEvent.User.Username,
		},
		Assignee:     assignee,
		SourceBranch: c.MergeEvent.ObjectAttributes.SourceBranch,
		TargetBranch: c.MergeEvent.ObjectAttributes.TargetBranch,
		Labels:       labels,
	}, nil
}
