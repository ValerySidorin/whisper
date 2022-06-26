package converters

import (
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	gitlab "github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/dto"
)

type MRConverter struct {
	MergeRequest *gitlab.MergeRequest
}

func NewMRConverter(mr *gitlab.MergeRequest) MRConverter {
	return MRConverter{
		MergeRequest: mr,
	}
}

func (c *MRConverter) Convert() (port.Messageable, error) {
	labels := make([]string, 0)
	for _, l := range c.MergeRequest.ObjectAttributes.Changes.Labels {
		labels = append(labels, l.Title)
	}
	return &dto.MergeRequest{
		ID:          c.MergeRequest.ObjectAttributes.ID,
		IID:         c.MergeRequest.ObjectAttributes.IID,
		ProjectID:   c.MergeRequest.Project.ID,
		Title:       c.MergeRequest.ObjectAttributes.Title,
		Description: c.MergeRequest.ObjectAttributes.Description,
		State:       c.MergeRequest.ObjectAttributes.State,
		CreatedDate: c.MergeRequest.ObjectAttributes.CreatedAt,
		UpdatedDate: c.MergeRequest.ObjectAttributes.UpdatedAt,
		URL:         c.MergeRequest.ObjectAttributes.URL,
		Author: dto.Person{
			Name:     c.MergeRequest.User.Name,
			UserName: c.MergeRequest.User.Username,
		},
		Assignee: dto.Person{
			Name:     c.MergeRequest.ObjectAttributes.Assignee.Name,
			UserName: c.MergeRequest.ObjectAttributes.Assignee.Username,
		},
		SourceBranch: c.MergeRequest.ObjectAttributes.SourceBranch,
		TargetBranch: c.MergeRequest.ObjectAttributes.TargetBranch,
		Labels:       labels,
	}, nil
}
