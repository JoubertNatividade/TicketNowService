package commands

import "ticketNowService/main/users/domain/repositories"

type DeleteUserCommand struct {
	repository repositories.IUsersRepository
}

func NewDeleteUserCommand(repository repositories.IUsersRepository) DeleteUserCommand {
	return DeleteUserCommand{repository}
}

func (itself DeleteUserCommand) Execute() error {
	err := itself.repository.Delete()
	if err != nil {
		return err
	}
	return nil
}
