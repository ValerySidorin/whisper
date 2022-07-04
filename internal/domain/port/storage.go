package port

type UserRepository interface {
	Get(messengerUsername string, vcsHosting string, messenger string)
}
