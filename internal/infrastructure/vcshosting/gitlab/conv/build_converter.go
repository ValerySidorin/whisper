package conv

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	gitlab "github.com/xanzy/go-gitlab"
)

type BuildEventConverter struct {
	BuildEvent *gitlab.BuildEvent
	Project    *gitlab.Project
}

func NewBuildConverter(be *gitlab.BuildEvent, p *gitlab.Project) BuildEventConverter {
	return BuildEventConverter{
		BuildEvent: be,
		Project:    p,
	}
}

func (c *BuildEventConverter) Convert() interface{} {
	b := &dto.Build{
		ID:       int64(c.BuildEvent.BuildID),
		Name:     c.BuildEvent.BuildName,
		Stage:    c.BuildEvent.BuildStage,
		Status:   c.BuildEvent.BuildStatus,
		Duration: c.BuildEvent.BuildDuration,
		Project: dto.Project{
			ID:          int64(c.Project.ID),
			Name:        c.Project.Name,
			Description: c.Project.Description,
		},
		User: dto.Person{
			ID:       int64(c.BuildEvent.User.ID),
			Name:     c.BuildEvent.User.Name,
			UserName: c.BuildEvent.User.Username,
		},
	}
	return &dto.BuildEvent{
		Build: b,
	}
}
