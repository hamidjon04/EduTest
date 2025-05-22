package service

import (
	"database/sql"
	"edutest/pkg/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hamidjon04/auth/auth"
)

func (s *Service) Register(req model.RegisterReq) (*model.Tokens, error) {
	id := uuid.NewString()
	claimItems := map[string]interface{}{"user_id": id, "created_at": time.Now().Unix()}

	accessToken, err := auth.GenerateJWT(s.Cfg.JWT_KEY, claimItems, time.Hour*24)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is generate access token: %v", err))
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(s.Cfg.JWT_KEY, claimItems, time.Hour*720)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is generate refresh token: %v", err))
		return nil, err
	}

	hashedPass, err := auth.HashPassword(req.Password)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is hash password: %v", err))
		return nil, err
	}
	req.Password = hashedPass

	err = s.Storage.Auth().Register(&req, id)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is create user: %v", err))
		return nil, err
	}

	err = s.Storage.Auth().CreateToken(model.Tokens{
		UserId:       id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is create user's token: %v", err))
		return nil, err
	}

	return &model.Tokens{
		UserId:       id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Login(req model.LoginReq) (*model.Tokens, error) {
	user, err := s.Storage.Auth().GetUserByUsername(req.Username)
	if user == nil || err == sql.ErrNoRows {
		s.Log.Error(fmt.Sprintf("User is not found: %v", req.Username))
		return nil, fmt.Errorf("User is not found: %v", req.Username)
	} else if err != nil {
		s.Log.Error(fmt.Sprintf("Error is get user: %v", err))
		return nil, err
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		s.Log.Error(fmt.Sprintf("Password is incorrect: %v", req.Password))
		return nil, fmt.Errorf("Password is incorrect")
	}

	claimItems := map[string]interface{}{"user_id": user.Id, "created_at": time.Now().Unix()}
	accessToken, err := auth.GenerateJWT(s.Cfg.JWT_KEY, claimItems, time.Hour*24)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is generate access token: %v", err))
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(s.Cfg.JWT_KEY, claimItems, time.Hour*720)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is generate refresh token: %v", err))
		return nil, err
	}

	err = s.Storage.Auth().CreateToken(model.Tokens{
		UserId:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is save user's tokens: %v", err))
		return nil, err
	}

	return &model.Tokens{
		UserId:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshToken(req model.RefreshTokenReq) (*model.Token, error) {
	_, err := auth.ValidateJWT(s.Cfg.JWT_KEY, req.RefreshToken)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is validate token: %v", err))
		return nil, err
	}
	

	claimItem := map[string]interface{}{"user_id": req.UserId, "created_at": time.Now().Unix()}
	accessToken, err := auth.GenerateJWT(s.Cfg.JWT_KEY, claimItem, time.Hour*24)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is generate access token: %v", err))
		return nil, err
	}

	err = s.Storage.Auth().CreateToken(model.Tokens{
		UserId:       req.UserId,
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is update access token: %v", err))
		return nil, err
	}
	return &model.Token{
		AccessToken: accessToken,
	}, nil
}
