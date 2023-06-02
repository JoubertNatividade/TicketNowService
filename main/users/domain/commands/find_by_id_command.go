package commands

import (
	"fmt"
	"ticketNowService/main/users/domain/repositories"
)

type FindByIdUserCommand struct {
	repository repositories.IUsersRepository
}

func NewFindByIdUserCommand(repository repositories.IUsersRepository) FindByIdUserCommand {
	return FindByIdUserCommand{repository}
}

func (itself FindByIdUserCommand) Execute() error {
	err := itself.repository.FindById()
	if err != nil {
		fmt.Print(err)
	}
	return nil
}
