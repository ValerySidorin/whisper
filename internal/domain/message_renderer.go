package domain

import (
	"fmt"

	"github.com/ValerySidorin/whisper/internal/domain/dto"
)

type DefaultMessageRenderer struct {
}

func RegisterDefaultMessageRenderer() *DefaultMessageRenderer {
	return &DefaultMessageRenderer{}
}

func (r *DefaultMessageRenderer) RenderMergeRequest(mr *dto.MergeRequest) string {
	switch mr.State {
	case "opened":
		return fmt.Sprintf("New Merge Request! | %v\n-- -- -- --\nTitle: %v\nDescription: %v\n%v\n-- -- -- --\nBranch: %v → %v\nAuthor: %v",
			mr.Project.Name, mr.Title, mr.Description, mr.URL, mr.SourceBranch, mr.TargetBranch, mr.Author.Name)
	default:
		return ""
	}
}

func (r *DefaultMessageRenderer) RenderDeployment(d *dto.Deployment) string {
	switch d.Status {
	case "success":
		return fmt.Sprintf("Deployment succeded! | %v\n-- -- -- --\nEnv: %v\n-- -- -- --\nJob: %v\n-- -- -- --\nCommit: %v\n%v\n-- -- -- --\nInitiator: %v", d.Project.Name, d.Environment, d.Job.Name, d.CommitTitle, d.CommitURL, d.User.Name)
	case "failed":
		return fmt.Sprintf("Deployment failed! | %v\n-- -- -- --\nEnv: %v\n-- -- -- --\n→ %v\n-- -- -- --\nJob: %v\n%v", d.Project.Name, d.Environment, d.Status, d.Job.Name, d.DeployableURL)
	default:
		return ""
	}
}
