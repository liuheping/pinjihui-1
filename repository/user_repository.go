package repository

import (
	"database/sql"
	"errors"
	"pinjihui.com/pinjihui/context"
	"pinjihui.com/pinjihui/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"github.com/rs/xid"
    "pinjihui.com/pinjihui/service"
)

const (
	defaultListFetchSize = 10
)

type UserRepository struct {
	db          *sqlx.DB
	roleRepository *RoleRepository
	log         *logging.Logger
}

func NewUserRepository(db *sqlx.DB, roleRepository *RoleRepository, log *logging.Logger) *UserRepository {
	return &UserRepository{db: db, roleRepository: roleRepository, log: log}
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	userSQL := `SELECT * FROM users WHERE email = $1`
	row := u.db.QueryRowx(userSQL, email)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return nil, err
	}

	roles, err := u.roleRepository.FindByUserId(&user.ID)
	if err != nil {
		u.log.Errorf("Error in retrieving roles : %v", err)
		return nil, err
	}
	user.Roles = roles
	return user, nil
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	userId := xid.New()
	user.ID = userId.String()
	userSQL := `INSERT INTO users (id, email, password, ip_address) VALUES (:id, :email, :password, :ip_address)`
	user.HashedPassword()
	_, err := u.db.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) List(first *int, after *string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if first == nil {
		*first = defaultListFetchSize
	}
	if after != nil {
		userSQL := `SELECT * FROM users WHERE created_at < (SELECT created_at FROM users WHERE id = $1) ORDER BY created_at DESC LIMIT $2;`
		decodedIndex, _ := service.DecodeCursor(after)
		err := u.db.Select(&users, userSQL, decodedIndex, first)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	userSQL := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1;`
	err := u.db.Select(&users, userSQL, first)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) Count() (int, error) {
	var count int
	userSQL := `SELECT count(*) FROM users`
	err := u.db.Get(&count, userSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserRepository) ComparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	if result := user.ComparePassword(userCredentials.Password); !result {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	return user, nil
}
