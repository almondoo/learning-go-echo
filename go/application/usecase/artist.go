package usecase

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/service"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/interface/validation"
)

type ArtistUsecase interface {
	Login(*validation.ArtistLoginRequest) (*auth.TokenDetails, error)
	Register(*validation.ArtistRegisterRequest) (*auth.TokenDetails, error)
	Logout(*auth.TokenData) error
	Edit(*validation.ArtistEditRequest, uint) (*entity.Artist, error)
}

type artistUsecase struct {
	as    service.ArtistService
	auth  auth.AuthInterface
	token auth.TokenInterface
}

func NewArtistUsecase(as service.ArtistService, auth auth.AuthInterface, token auth.TokenInterface) ArtistUsecase {
	return &artistUsecase{as: as, auth: auth, token: token}
}

func (au *artistUsecase) Login(request *validation.ArtistLoginRequest) (*auth.TokenDetails, error) {
	artist, err := au.as.Login(request)
	if err != nil {
		return nil, err
	}

	token, err := au.token.CreateToken(int(artist.ID))
	if err != nil {
		return nil, err
	}

	if err := au.auth.CreateAuth(int(artist.ID), token); err != nil {
		return nil, err
	}

	return token, nil
}

func (au *artistUsecase) Register(request *validation.ArtistRegisterRequest) (*auth.TokenDetails, error) {
	artist, err := au.as.Register(request)
	if err != nil {
		return nil, err
	}

	token, err := au.token.CreateToken(int(artist.ID))
	if err != nil {
		return nil, err
	}

	if err := au.auth.CreateAuth(int(artist.ID), token); err != nil {
		return nil, err
	}

	return token, nil
}

func (au *artistUsecase) Logout(tokens *auth.TokenData) error {
	if err := au.as.Logout(tokens); err != nil {
		return err
	}

	return nil
}

func (au *artistUsecase) Edit(request *validation.ArtistEditRequest, artistID uint) (*entity.Artist, error) {
	artist, err := au.as.Edit(request, artistID)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
