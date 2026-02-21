package postgresuser

import (
	"goProject/internal/entity"
	"goProject/internal/repository/postgres"
)

func scanUser(scanner postgres.Scanner) (entity.User, error) {
	var user entity.User
	var roleStr string

	err := scanner.Scan(&user.ID ,&user.Email ,&user.Name ,&user.Password ,&roleStr ,&user.Status , &user.CreatedAt ,&user.UpdatedAt,&user.EmailVerify )

	user.Role = entity.MapToRoleEntity(roleStr)

	return user, err
}