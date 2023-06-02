package commands

import "ticketNowService/main/users/domain/repositories"

type ListAllUsersCommand struct {
	repository repositories.IUsersRepository
}

func NewListUsersCommand(repository repositories.IUsersRepository) ListAllUsersCommand {
	return ListAllUsersCommand{repository}
}

func (itself ListAllUsersCommand) Execute() error {
	err := itself.repository.FindAll()
	if err != nil {
		return err
	}
	return nil
}
