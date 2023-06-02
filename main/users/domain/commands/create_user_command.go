package commands

import "ticketNowService/main/users/domain/repositories"

type CreateUserCommand struct {
	repository repositories.IUsersRepository
}

func NewCreateUserCommand(repository repositories.IUsersRepository) CreateUserCommand {
	return CreateUserCommand{repository}
}

func (itself CreateUserCommand) Execute() error {
	err := itself.repository.Create()

	if err != nil {
		return err
	}
	return nil
}
