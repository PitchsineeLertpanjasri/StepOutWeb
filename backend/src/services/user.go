package services

import (
	"errors"

	"stepoutsite/domain/entities"
	"stepoutsite/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository repositories.IUserRepository
}

type IUserService interface {
	GetAllUsers(filter bson.M , studentID string) (*[]entities.UserDataFormat, error)
	CreateUser(user entities.UserDataFormat) error
	GetOneUser(studentID string) (entities.UserDataFormat, error)
	Login(req *entities.UserDataFormat) (string,error)
	CheckPermissionCoreAndAdmin(studentID string) error
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (sv userService) CreateUser(user entities.UserDataFormat) error {
	if user.StudentID == ""{
		return errors.New("please fill in student id")
	}

	err := sv.UserRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (sv userService) GetAllUsers(filter bson.M, studentID string) (*[]entities.UserDataFormat, error){	
	err := sv.CheckPermissionCoreAndAdmin(studentID)
	if err != nil {
		return nil,err
	}

	users,err := sv.UserRepository.GetAllUsers(filter)

	if err != nil {
		return nil,err
	}

	return users,nil
}

func checkPasswords(hashedPwd string, plainPwd string) error {
    byteHash := []byte(hashedPwd)
	pwd := []byte(plainPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, pwd)
    if err != nil {
        return errors.New("passwords do not match")
    }
    
    return nil
}

func (sv userService) GetOneUser(studentID string) (entities.UserDataFormat, error) {
	user,err := sv.UserRepository.GetOneUser(studentID)
	
	if err != nil {
		return entities.UserDataFormat{},err
	}

	return user,nil
}

func (sv userService) Login(req *entities.UserDataFormat) (string,error) {
	user,err := sv.UserRepository.GetOneUser(req.StudentID)

	if err != nil {
		return "",errors.New("user not found")
	}

	if err := checkPasswords(user.Password, req.Password); err != nil {
		return "", errors.New("passwords do not match")
	}

	return sv.UserRepository.Login(req)
}

func (sv userService) CheckPermissionCoreAndAdmin(studentID string) error {
	admin, err := sv.UserRepository.GetOneUser(studentID)
	if err != nil {
		return errors.New("user not found")
	}
	if !(admin.Role == "core" || admin.Role == "admin") {
		return errors.New("unauthorized")
	}
	return nil
}