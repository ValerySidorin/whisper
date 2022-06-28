package converters

import (
	"github.com/ValerySidorin/whisper/internal/domain/dto"
	"github.com/ValerySidorin/whisper/internal/domain/port"
	gitlab "github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/dto"
)

type DeploymentConverter struct {
	Deployment *gitlab.Deployment
}

func NewDeploymentConverter(d *gitlab.Deployment) DeploymentConverter {
	return DeploymentConverter{
		Deployment: d,
	}
}

func (c *DeploymentConverter) Convert() (port.Messageable, error) {
	return &dto.Deployment{
		ID:            c.Deployment.DeploymentID,
		Status:        c.Deployment.Status,
		DeployableURL: c.Deployment.DeployableURL,
		Environment:   c.Deployment.Environment,
		Project: dto.Project{
			ID:          c.Deployment.Project.ID,
			Name:        c.Deployment.Project.Name,
			Description: c.Deployment.Project.Description,
		},
		User: dto.Person{
			Name:     c.Deployment.User.Name,
			UserName: c.Deployment.User.Username,
		},
		CommitURL:   c.Deployment.CommitURL,
		CommitTitle: c.Deployment.CommitTitle,
	}, nil
}
