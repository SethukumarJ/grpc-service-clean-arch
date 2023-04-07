package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
	"github.com/SethukumarJ/go-gin-clean-arch/pkg/pb"
	services "github.com/SethukumarJ/go-gin-clean-arch/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	Id       int64  `copier:"must"`
	Email    string `copier:"must"`
	Password string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}


func (cr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.Users{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := cr.userUseCase.Save(ctx, user)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

// save godoc
// @summary Get all users
// @description register user
// @tags users
// @id register
// @param RegisterUser body domain.Users{} true "user signup"
// @produce json
// @Router /api/users [post]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) Save(c *gin.Context) {
	var user domain.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := cr.userUseCase.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}

// func (cr *UserHandler) Delete(c *gin.Context) {
// 	paramsId := c.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Cannot parse id",
// 		})
// 		return
// 	}

// 	ctx := c.Request.Context()
// 	user, err := cr.userUseCase.FindByID(ctx, uint(id))

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// 	if user == (domain.Users{}) {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "User is not booking yet",
// 		})
// 		return
// 	}

// 	cr.userUseCase.Delete(ctx, user)

// 	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
// }


func (cr *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// var user domain.Users

	// if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
	// 	return &pb.LoginResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	// match := utils.CheckPasswordHash(req.Password, user.Password)

	// if !match {
	// 	return &pb.LoginResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	// token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  "token",
	}, nil
}

func (cr *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	// claims, err := s.Jwt.ValidateToken(req.Token)

	// if err != nil {
	// 	return &pb.ValidateResponse{
	// 		Status: http.StatusBadRequest,
	// 		Error:  err.Error(),
	// 	}, nil
	// }

	var user domain.Users

	// if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
	// 	return &pb.ValidateResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (cr *UserHandler) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// // Check if the ID is not empty or invalid
	// if req.Id == 0 {
	// 	return &pb.DeleteResponse{
	// 		Status: http.StatusBadRequest,
	// 		Error:  "Invalid ID",
	// 	}, nil
	// }

	var user domain.Users

	// // Check if the record exists in the database
	// result := s.H.DB.First(&user, "id = ?", req.Id)
	// if result.Error != nil {
	// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		return &pb.DeleteResponse{
	// 			Status: http.StatusNotFound,
	// 			Error:  "Record not found",
	// 		}, nil
	// 	} else {
	// 		return &pb.DeleteResponse{
	// 			Status: http.StatusInternalServerError,
	// 			Error:  result.Error.Error(),
	// 		}, nil
	// 	}
	// }

	// // Delete the record from the database
	// result = s.H.DB.Delete(&user, req.Id)
	// if result.Error != nil {
	// 	return &pb.DeleteResponse{
	// 		Status: http.StatusInternalServerError,
	// 		Error:  result.Error.Error(),
	// 	}, nil
	// }

	return &pb.DeleteResponse{
		Status: http.StatusOK,
		Id:     user.Id,
	}, nil
}


// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @security ApiKeyAuth
// @id FindAll
// @produce json
// @Router /api/users [get]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &users)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	user, err := cr.userUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}
}
