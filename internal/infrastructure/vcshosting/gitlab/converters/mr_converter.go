package converters

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	gitlab "github.com/xanzy/go-gitlab"
)

type MRConverter struct {
	MergeEvent *gitlab.MergeEvent
	User       *gitlab.User
}

type MREventConverter struct {
	MergeEvent *gitlab.MergeEvent
	User       *gitlab.User
}

func NewMRConverter(mr *gitlab.MergeEvent, u *gitlab.User) MRConverter {
	return MRConverter{
		MergeEvent: mr,
		User:       u,
	}
}

func NewMREventConverter(mr *gitlab.MergeEvent, u *gitlab.User) MRConverter {
	return MRConverter{
		MergeEvent: mr,
		User:       u,
	}
}

func (c *MRConverter) Convert() interface{} {
	labels := make([]string, 0)
	for _, l := range c.MergeEvent.Labels {
		labels = append(labels, l.Name)
	}
	var assignee dto.Person
	if c.MergeEvent.Assignee != nil {
		assignee = dto.Person{
			ID:       int64(c.MergeEvent.Assignee.ID),
			Name:     c.MergeEvent.Assignee.Name,
			UserName: c.MergeEvent.Assignee.Username,
		}
	}
	if assignee.ID == 0 {
		assignee = dto.Person{
			ID: int64(c.MergeEvent.ObjectAttributes.AssigneeID),
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
			ID:       int64(c.User.ID),
			Name:     c.User.Name,
			UserName: c.User.Username,
		},
		Assignee:     assignee,
		SourceBranch: c.MergeEvent.ObjectAttributes.SourceBranch,
		TargetBranch: c.MergeEvent.ObjectAttributes.TargetBranch,
		Labels:       labels,
	}
}

func (c *MREventConverter) Convert() interface{} {
	labels := make([]string, 0)
	for _, l := range c.MergeEvent.Labels {
		labels = append(labels, l.Name)
	}
	var assignee dto.Person
	if c.MergeEvent.Assignee != nil {
		assignee = dto.Person{
			ID:       int64(c.MergeEvent.Assignee.ID),
			Name:     c.MergeEvent.Assignee.Name,
			UserName: c.MergeEvent.Assignee.Username,
		}
	}
	if assignee.ID == 0 {
		assignee = dto.Person{
			ID: int64(c.MergeEvent.ObjectAttributes.AssigneeID),
		}
	}
	mr := &dto.MergeRequest{
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
			ID:       int64(c.User.ID),
			Name:     c.User.Name,
			UserName: c.User.Username,
		},
		Assignee:     assignee,
		SourceBranch: c.MergeEvent.ObjectAttributes.SourceBranch,
		TargetBranch: c.MergeEvent.ObjectAttributes.TargetBranch,
		Labels:       labels,
	}
	return &dto.MergeRequestEvent{
		MergeRequest: mr,
		Event:        c.MergeEvent.ObjectAttributes.Action,
	}
}
