package postgresuser

import (
	"goProject/internal/entity"
	"goProject/internal/repository/postgres"
)

func scanUser(scanner postgres.Scanner) (entity.User, error) {
	var user entity.User
	var roleStr string

	err := scanner.Scan(&user.ID ,&user.Email ,&user.EmailVerify ,&user.Name ,&user.Password ,&roleStr ,&user.Status , &user.CreatedAt ,&user.UpdatedAt)

	user.Role = entity.MapToRoleEntity(roleStr)

	return user, err
}

const UserColumns = `id, email, email_verify, name, password, role, status, created_at, updated_at`
