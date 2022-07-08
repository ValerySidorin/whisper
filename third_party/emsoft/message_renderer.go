package emsoft

import (
	"regexp"
	"strings"

	dto "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"
)

type EmsoftMessageRenderer struct {
}

func RegisterEmsoftMessageRenderer() *EmsoftMessageRenderer {
	return &EmsoftMessageRenderer{}
}

func (r *EmsoftMessageRenderer) RenderMergeRequestEvent(mre *dto.MergeRequestEvent) string {
	trackerLink := ""
	matched, _ := regexp.MatchString(`^\w+-\d+\s`, mre.MergeRequest.Title)
	if matched {
		trackerLink = "https://tracker.yandex.ru/" + mre.MergeRequest.Title[:strings.IndexByte(mre.MergeRequest.Title, ' ')]
	}
	switch mre.MergeRequest.State {
	case "opened":
		if mre.Event == "update" {
			return "Merge Request has been updated! | " + mre.MergeRequest.Project.Name + "\n-- -- -- --\nTitle: " + mre.MergeRequest.Title + "\nDescription: " + mre.MergeRequest.Description + "\n" + mre.MergeRequest.URL + "\n-- -- -- --\nTracker: " + trackerLink + "\n-- -- -- --\nBranch: " + mre.MergeRequest.SourceBranch + "→ " + mre.MergeRequest.TargetBranch + "\nAuthor: " + mre.MergeRequest.Author.Name
		}
		return "New Merge Request! | " + mre.MergeRequest.Project.Name + "\n-- -- -- --\nTitle: " + mre.MergeRequest.Title + "\nDescription: " + mre.MergeRequest.Description + "\n" + mre.MergeRequest.URL + "\n-- -- -- --\nTracker: " + trackerLink + "\n-- -- -- --\nBranch: " + mre.MergeRequest.SourceBranch + "→ " + mre.MergeRequest.TargetBranch + "\nAuthor: " + mre.MergeRequest.Author.Name
	case "merged":
		return "Merge Request has been merged!" + "\n-- -- -- --\n" + mre.MergeRequest.URL
	case "closed":
		return "Merge Request has been closed!" + "\n-- -- -- --\n" + mre.MergeRequest.URL
	default:
		return ""
	}
}

func (r *EmsoftMessageRenderer) RenderDeploymentEvent(de *dto.DeploymentEvent) string {
	switch de.Deployment.Status {
	case "success":
		return "Deployment succeded! | " + de.Deployment.Project.Name + "\n-- -- -- --\nEnv: " + de.Deployment.Environment + "\n-- -- -- --\nJob: " + de.Deployment.Job.Name + "\n-- -- -- --\nCommit: " + de.Deployment.CommitTitle + "\n" + de.Deployment.CommitURL + "\n-- -- -- --\nInitiator: " + de.Deployment.User.Name
	case "failed":
		return "Deployment failed! | " + de.Deployment.Project.Name + "\n-- -- -- --\nEnv: " + de.Deployment.Environment + "\n-- -- -- --\nJob: " + de.Deployment.Job.Name + "\n" + de.Deployment.DeployableURL
	default:
		return ""
	}
}
