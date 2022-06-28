package vcshosting

import (
	"encoding/json"
	"fmt"

	"github.com/ValerySidorin/whisper/internal/domain/port"
	gitlabconv "github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/converters"
	gitlabdto "github.com/ValerySidorin/whisper/internal/infrastructure/vcshosting/gitlab/dto"
	"github.com/tidwall/gjson"
)

func GetMessageable(provider string, j string) (port.Messageable, error) {
	switch provider {
	case "gitlab":
		objKind := gjson.Get(j, "object_kind").String()
		return GetMessageableFromGitlabDto(objKind, j)
	default:
		return nil, fmt.Errorf("unknown event provider: %v", provider)
	}
}

func GetMessageableFromGitlabDto(objKind string, j string) (port.Messageable, error) {
	switch objKind {
	case "merge_request":
		mr := gitlabdto.MergeRequest{}
		if err := json.Unmarshal([]byte(j), &mr); err != nil {
			return nil, err
		}
		conv := gitlabconv.NewMRConverter(&mr)
		res, err := conv.Convert()
		if err != nil {
			return nil, err
		}
		return res, nil
	case "deployment":
		d := gitlabdto.Deployment{}
		if err := json.Unmarshal([]byte(j), &d); err != nil {
			return nil, err
		}
		conv := gitlabconv.NewDeploymentConverter(&d)
		res, err := conv.Convert()
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unknown object kind: %v", objKind)
	}
}
