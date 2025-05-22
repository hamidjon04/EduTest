package postgres

import (
	"database/sql"
	"edutest/pkg/model"
	"fmt"
	"log/slog"
)

type AuthRepo interface {
	Register(req *model.RegisterReq, id string) error
	CreateToken(req model.Tokens) error
	GetUserByUsername(username string) (*model.User, error)
}

type authImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewAuthRepo(db *sql.DB, log *slog.Logger) AuthRepo {
	return &authImpl{
		DB:  db,
		Log: log,
	}
}

func (u *authImpl) Register(req *model.RegisterReq, id string) error {
	query := `
				INSERT INTO users(
					id, username, name, lastname, password)
				VALUES
					($1, $2, $3, $4, $5)`
	_, err := u.DB.Exec(query, id, req.Username, req.Name, req.Lastname, req.Password)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error is create user: %v", err))
		return err
	}
	return nil
}

func (u *authImpl) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	query := `
				SELECT 
					id, username, name, lastname, password, phone_number, created_at
				FROM 
					users
				WHERE 
					username = $1 AND deleted_at IS NULL`
	err := u.DB.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Name, &user.Lastname, &user.Password, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error is get user: %v", err))
		return nil, err
	}
	return &user, nil
}

func (u *authImpl) CreateToken(req model.Tokens) error {
	query := `
				INSERT INTO user_tokens(
					user_id, access_token, refresh_token)
				VALUES
					($1, $2, $3)
				ON CONFLICT (user_id) DO UPDATE
				SET 
					refresh_token = EXCLUDED.refresh_token,
					access_token = EXCLUDED.access_token`
	_, err := u.DB.Exec(query, req.UserId, req.AccessToken, req.RefreshToken)
	if err != nil {
		u.Log.Error(fmt.Sprintf("Error is create user's token's: %v", err))
		return err
	}
	return nil
}
