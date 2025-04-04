package user

import (
	"context"
	"errors"
	"log"

	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
	"github.com/SahilBheke25/quick-farm-backend/internal/repository"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	Authenticate(ctx context.Context, username, password string) (models.User, error)
	RegisterUser(ctx context.Context, user models.User) error
	UserProfile(ctx context.Context, userId int) (models.User, error)
	OwnerByEquipmentId(ctx context.Context, equipId int) (user models.User, err error)
	UpdateUserProfile(ctx context.Context, updateUser models.User) (models.User, error)
}

func NewService(user repository.UserStorer) Service {
	return service{userRepo: user}
}

func (s service) Authenticate(ctx context.Context, username, password string) (models.User, error) {

	// DB call
	resp, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Service: error occured while calling getUserByUsername DB opeartion, err : %v\n", err)
		return models.User{}, err
	}

	if resp.Password != password {
		return models.User{}, apperrors.ErrInvalidCredentials
	}
	// Password Verification
	// if !utils.CheckPasswordHash(password, resp.Password) {
	// 	log.Println(resp.Password, " ", password)
	// 	return models.User{}, apperrors.ErrInvalidCredentials
	// }

	return resp, nil
}

func (s service) RegisterUser(ctx context.Context, user models.User) error {

	// hashedPassword, err := utils.HashPassword(user.Password)
	// if err != nil {
	// 	log.Printf("Service: error occured while hashing password, err : %v\n", err)
	// 	return err
	// }
	// user.Password = hashedPassword

	//DB call
	err := s.userRepo.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Service: error occured while calling CreateUser DB opeartion, err : %v\n", err)
		return err
	}

	return nil
}

func (s service) UserProfile(ctx context.Context, userId int) (user models.User, err error) {

	// DB call
	user, err = s.userRepo.UserProfile(ctx, userId)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return
		}

		log.Printf("Service: Failed to fetch user with ID %d, err: %v\n", userId, err)
		err = apperrors.ErrInternal
		return
	}

	return
}

func (s service) OwnerByEquipmentId(ctx context.Context, equipmentID int) (user models.User, err error) {

	// DB call
	user, err = s.userRepo.OwnerByEquipmentId(ctx, equipmentID)
	if err != nil {
		log.Printf("Service: Error fetching owner for EquipmentID %d, err: %v\n", equipmentID, err)
		return
	}

	return
}

func (s service) UpdateUserProfile(ctx context.Context, updateUser models.User) (models.User, error) {
	// Validate if user ID is provided
	if updateUser.Id <= 0 {
		log.Println("Service: Invalid user ID")
		return models.User{}, apperrors.ErrInvalidUserID
	}

	// Call the repository layer to update the user profile
	updatedUser, err := s.userRepo.UpdateUserProfile(ctx, updateUser)
	if err != nil {
		log.Printf("Service: Error updating user profile, err: %v\n", err)
		return models.User{}, err
	}

	return updatedUser, nil
}
