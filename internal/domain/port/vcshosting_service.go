package port

import "github.com/ValerySidorin/whisper/internal/domain/dto/vcshosting"

type VCSHostingService interface {
	GetPerson(username string) (*vcshosting.Person, error)
}
