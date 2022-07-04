package domain

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
)

type DefaultMessageRenderer struct {
}

func RegisterDefaultMessageRenderer() *DefaultMessageRenderer {
	return &DefaultMessageRenderer{}
}

func (r *DefaultMessageRenderer) RenderMergeRequest(mr *dto.MergeRequest) string {
	switch mr.State {
	case "opened":
		return "New Merge Request! | " + mr.Project.Name + "\n-- -- -- --\nTitle: " + mr.Title + "\nDescription: " + mr.Description + "\n" + mr.URL + "\n-- -- -- --\nBranch: " + mr.SourceBranch + "→ " + mr.TargetBranch + "\nAuthor: " + mr.Author.Name
	default:
		return ""
	}
}

func (r *DefaultMessageRenderer) RenderDeployment(d *dto.Deployment) string {
	switch d.Status {
	case "success":
		return "Deployment succeded! | " + d.Project.Name + "\n-- -- -- --\nEnv: " + d.Environment + "\n-- -- -- --\nJob: " + d.Job.Name + "\n-- -- -- --\nCommit: " + d.CommitTitle + "\n" + d.CommitURL + "\n-- -- -- --\nInitiator: " + d.User.Name
	case "failed":
		return "Deployment failed! | " + d.Project.Name + "\n-- -- -- --\nEnv: " + d.Environment + "\n-- -- -- --\nJob: " + d.Job.Name + "\n" + d.DeployableURL
	default:
		return ""
	}
}
