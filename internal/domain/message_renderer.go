package domain

import (
	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
)

type DefaultMessageRenderer struct {
}

func RegisterDefaultMessageRenderer() *DefaultMessageRenderer {
	return &DefaultMessageRenderer{}
}

func (r *DefaultMessageRenderer) RenderMergeRequestEvent(mre *dto.MergeRequestEvent) string {
	switch mre.MergeRequest.State {
	case "opened":
		if mre.Event == "update" {
			return "Merge Request has been updated! | " + mre.MergeRequest.Project.Name + "\n-- -- -- --\nTitle: " + mre.MergeRequest.Title + "\nDescription: " + mre.MergeRequest.Description + "\n" + mre.MergeRequest.URL + "\n-- -- -- --\nBranch: " + mre.MergeRequest.SourceBranch + "→ " + mre.MergeRequest.TargetBranch + "\nAuthor: " + mre.MergeRequest.Author.Name
		}
		if mre.Event == "rebase" {
			return "Merge Request rebase triggered! | " + mre.MergeRequest.Project.Name + "\n-- -- -- --\nTitle: " + mre.MergeRequest.Title + "\nDescription: " + mre.MergeRequest.Description + "\n" + mre.MergeRequest.URL + "\n-- -- -- --\nBranch: " + mre.MergeRequest.SourceBranch + "→ " + mre.MergeRequest.TargetBranch + "\nAuthor: " + mre.MergeRequest.Author.Name
		}
		return "New Merge Request! | " + mre.MergeRequest.Project.Name + "\n-- -- -- --\nTitle: " + mre.MergeRequest.Title + "\nDescription: " + mre.MergeRequest.Description + "\n" + mre.MergeRequest.URL + "\n-- -- -- --\nBranch: " + mre.MergeRequest.SourceBranch + "→ " + mre.MergeRequest.TargetBranch + "\nAuthor: " + mre.MergeRequest.Author.Name
	case "merged":
		return "Congrats! you have been merged!" + "\n-- -- -- --\n" + mre.MergeRequest.URL
	case "closed":
		return "Whoops! Your Merge Request has been closed!" + "\n-- -- -- --\n" + mre.MergeRequest.URL
	default:
		return ""
	}
}

func (r *DefaultMessageRenderer) RenderDeploymentEvent(de *dto.DeploymentEvent) string {
	switch de.Deployment.Status {
	case "success":
		return "Deployment succeded! | " + de.Deployment.Project.Name + "\n-- -- -- --\nEnv: " + de.Deployment.Environment + "\n-- -- -- --\nJob: " + de.Deployment.Job.Name + "\n-- -- -- --\nCommit: " + de.Deployment.CommitTitle + "\n" + de.Deployment.CommitURL + "\n-- -- -- --\nInitiator: " + de.Deployment.User.Name
	case "failed":
		return "Deployment failed! | " + de.Deployment.Project.Name + "\n-- -- -- --\nEnv: " + de.Deployment.Environment + "\n-- -- -- --\nJob: " + de.Deployment.Job.Name + "\n" + de.Deployment.DeployableURL
	default:
		return ""
	}
}
