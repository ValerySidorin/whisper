package converters

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
	"github.com/xanzy/go-gitlab"
)

type DeploymentConverter struct {
	DeploymentEvent *gitlab.DeploymentEvent
	Job             *gitlab.Job
}

func NewDeploymentConverter(d *gitlab.DeploymentEvent, j *gitlab.Job) DeploymentConverter {
	return DeploymentConverter{
		DeploymentEvent: d,
		Job:             j,
	}
}

func (c *DeploymentConverter) Convert() (interface{}, error) {
	return &dto.Deployment{
		Status: c.DeploymentEvent.Status,
		Job: dto.Job{
			ID:   int64(c.Job.ID),
			Name: c.Job.Name,
		},
		DeployableURL: c.DeploymentEvent.DeployableURL,
		Environment:   c.DeploymentEvent.Environment,
		Project: dto.Project{
			ID:          int64(c.DeploymentEvent.Project.ID),
			Name:        c.DeploymentEvent.Project.Name,
			Description: c.DeploymentEvent.Project.Description,
		},
		User: dto.Person{
			Name:     c.DeploymentEvent.User.Name,
			UserName: c.DeploymentEvent.User.Username,
		},
		CommitURL:   c.DeploymentEvent.CommitURL,
		CommitTitle: c.DeploymentEvent.CommitTitle,
	}, nil
}
