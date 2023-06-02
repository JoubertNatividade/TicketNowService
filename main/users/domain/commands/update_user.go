package commands

import "ticketNowService/main/users/domain/repositories"

type UpdateUserCommand struct {
	repository repositories.IUsersRepository
}

func NewUpdateUserCommand(repository repositories.IUsersRepository) UpdateUserCommand {
	return UpdateUserCommand{repository}
}

func (itself UpdateUserCommand) Execute() error {
	err := itself.repository.Update()
	if err != nil {
		return err
	}
	return nil
}
