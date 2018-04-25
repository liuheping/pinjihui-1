package repository

import (
	"database/sql"
	"pinjihui.com/pinjihui/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
    "fmt"
)

type RoleRepository struct {
	db  *sqlx.DB
	log *logging.Logger
}

type RoleWithUser struct {
    model.Role
    UserId string `db:"user_id"`
}
func NewRoleRepository(db *sqlx.DB, log *logging.Logger) *RoleRepository {
	return &RoleRepository{db: db, log: log}
}

func (r *RoleRepository) FindByUserId(userId *string) ([]*model.Role, error) {
	roles := make([]*model.Role, 0)

	roleSQL := `SELECT role.*
	FROM roles role
	INNER JOIN rel_users_roles ur ON role.id = ur.role_id
	WHERE ur.user_id = $1 `
	err := r.db.Select(&roles, roleSQL, userId)
	if err == sql.ErrNoRows {
		return roles, nil
	}
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) FindByUserIds(userIds []string) ([]*RoleWithUser, error) {
    roles := make([]*RoleWithUser, 0)

    rolesSQL := `SELECT roles.*, ur.user_id
    FROM roles, rel_users_roles ur where roles.id = ur.role_id
    and ur.user_id in (?)`
    query, args, err := sqlx.In(rolesSQL, userIds)
    if err != nil {
        fmt.Errorf("sqlx.In(rolesSQL, userIds) err: %v", err)
        return nil, err
    }
    query = r.db.Rebind(query)
    fmt.Println(query, args)
    err = r.db.Select(&roles, query, args...)

    if err == sql.ErrNoRows {
        return roles, nil
    }
    if err != nil {
        fmt.Errorf("r.db.Select(&roles, query, args) err: %v", err)
        return nil, err
    }
    return roles, nil
}
