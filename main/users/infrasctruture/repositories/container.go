package repositories

import (
	"github.com/google/wire"
	"ticketNowService/main/users/domain/repositories"
)

var Container = wire.NewSet(
	NewGormUserRepository,
	wire.Bind(new(repositories.IUsersRepository), new(*GormUserRepository)),
)
