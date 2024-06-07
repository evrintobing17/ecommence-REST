package usecase

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/evrintobing17/ecommence-REST/app/helpers/bcrypthelper"
	"github.com/evrintobing17/ecommence-REST/app/helpers/jwthelper"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/seller"
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

type UC struct {
	repo seller.SellerRepository
}

var (
	JWTDuration = 1 * time.Hour

	PasswordHashCost = 14

	//ErrInvalidCredential - Standard error message for invalid credential
	ErrInvalidCredential    = errors.New("invalid credential")
	ErrInvalidToken         = errors.New("invalid token")
	ErrUserNotFound         = errors.New("user not found")
	ErrUIDNotYetRegistered  = errors.New("the user id is not yet registered to DB")
	ErrFirebaseTokenInvalid = errors.New("firebase authentication token is invalid")
	ErrGameKeyInvalid       = errors.New("error game key invalid")
	ErrEmailAlreadyExist    = errors.New("email already exist")
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrPhoneAlreadyExist    = errors.New("phone already exist")
	ErrURLCallbackNotSet    = errors.New("callback url not set")
	//time in second
	// minute int64 = 60
	// hour         = minute * 60
	// day          = 24 * hour
)

func NewSellerUsecase(repo seller.SellerRepository) seller.SellerUsecase {
	return &UC{
		repo: repo,
	}
}

func (uc *UC) Login(email, password string) (user *models.Seller, token string, err error) {
	user, err = uc.repo.GetByEmail(email)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, "", ErrInvalidCredential
		}
		return nil, "", err
	}

	//compare password against Hash
	err = bcrypthelper.CompareBcrypt(user.Password, password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, "", ErrInvalidCredential
		} else {
			return nil, "", err
		}
	}
	jwtExpirationDurationDayString := os.Getenv("jwt.expirationDurationDay")
	var jwtExpirationDurationDay int
	jwtExpirationDurationDay, err = strconv.Atoi(jwtExpirationDurationDayString)
	if err != nil {
		return nil, "", err
	}

	// Conversion to seconds
	jwtExpiredAt := time.Now().Unix() + int64(jwtExpirationDurationDay*3600*24)

	userClaims := jwthelper.AccessJWTClaims{Id: user.ID, ExpiresAt: jwtExpiredAt}
	jwtToken, err := jwthelper.NewWithClaims(userClaims)
	if err != nil {
		return nil, "", err
	}

	return user, jwtToken, nil
}

func (uc *UC) RefreshAccessJWT(userId int) (newAccessJWT string, err error) {
	//create new AccessJWT
	accessJWT, err := uc.generateUserJWTDriver(userId)
	if err != nil {
		return "", err
	}

	return accessJWT, nil
}

func (uc *UC) generateUserJWTDriver(userId int) (token string, err error) {
	//Create JWT
	jwtExpiredAt := time.Now().Add(JWTDuration).Unix()

	userClaims := jwthelper.AccessJWTClaims{
		Id: userId, ExpiresAt: jwtExpiredAt,
	}
	jwtToken, err := jwthelper.NewWithClaims(userClaims)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (uc *UC) Register(username, email, password, address string) (driver *models.Seller, err error) {
	//Cast ReqRegister to model.User
	userModel := models.Seller{
		Name:     username,
		Email:    email,
		Password: password,
		Address:  address,
	}
	errUser := uc.isSellerExist(userModel)
	if errUser != nil {
		if err == ErrEmailAlreadyExist {
			return nil, ErrEmailAlreadyExist
		}
		if err == ErrPhoneAlreadyExist {
			return nil, ErrPhoneAlreadyExist
		}
		return nil, errUser
	}

	//Convert Plain password to bcrypt
	hashedPassword, err := bcrypthelper.GenerateBcrypt(password)
	if err != nil {
		return nil, err
	}
	userModel.Password = hashedPassword

	userData, errUser := uc.repo.Insert(userModel)
	if errUser != nil {
		return nil, errUser
	}

	return userData, nil
}

func (uc *UC) DeleteByID(userId int) error {
	err := uc.repo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UC) Update(updateData map[string]interface{}) (*models.Seller, error) {
	updateData["updated_at"] = time.Now()
	userData, err := uc.repo.UpdatePartial(updateData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (uc *UC) isSellerExist(user models.Seller) error {
	exist, _ := uc.repo.ExistByUsername(user.Name)
	if exist {
		return ErrUsernameAlreadyExist
	}

	exist, _ = uc.repo.ExistByEmail(user.Email)
	if exist {
		return ErrEmailAlreadyExist
	}

	return nil
}
