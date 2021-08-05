package service

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/infrastructure/security"
	"learning-go-echo/interface/validation"
)

type ArtistService interface {
	Login(*validation.ArtistLoginRequest) (*entity.Artist, error)
	Register(*validation.ArtistRegisterRequest) (*entity.Artist, error)
	Logout(*auth.TokenData) error
	Edit(*validation.ArtistEditRequest, uint) (*entity.Artist, error)
}

type artistService struct {
	artistRepo repository.ArtistRepository
	auth       auth.AuthInterface
	token      auth.TokenInterface
}

func NewArtistService(artistRepo repository.ArtistRepository, auth auth.AuthInterface, token auth.TokenInterface) ArtistService {
	return &artistService{artistRepo: artistRepo, auth: auth, token: token}
}

func (us *artistService) Login(request *validation.ArtistLoginRequest) (*entity.Artist, error) {
	artist, err := us.artistRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := security.VerifyPassword(artist.Password, request.Password); err != nil {
		return nil, err
	}

	return artist, nil
}

func (us *artistService) Register(request *validation.ArtistRegisterRequest) (*entity.Artist, error) {
	password, err := security.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	artist := &entity.Artist{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	createArtist, err := us.artistRepo.Create(artist)
	if err != nil {
		return nil, err
	}

	return createArtist, nil
}

func (us *artistService) Logout(tokens *auth.TokenData) error {
	detailtoken, err := us.token.ExtractTokenMetadata(tokens.AccessToken)
	if err != nil {
		return err
	}
	if err := us.auth.DeleteTokens(detailtoken); err != nil {
		return err
	}
	return nil
}

func (us *artistService) Edit(request *validation.ArtistEditRequest, artistID uint) (artist *entity.Artist, err error) {
	artist, err = us.artistRepo.FindByID(artistID)
	if err != nil {
		return nil, err
	}
	artist.Name = request.Name
	artist.Email = request.Email
	artist, err = us.artistRepo.Update(artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
