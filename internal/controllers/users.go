package controllers

import (
	"log"
	"net/http"
	"todo/internal/models"
	databse "todo/pkg/database"
	"todo/pkg/response"
	"todo/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	meta := utils.Meta{}
	if user.Password != "" {
		var errPass error
		user.Password, errPass = utils.HashPassword(user.Password)
		if errPass != nil {
			response.BadResponse(c, utils.MessageBox["password_issue"], user, meta)
			return
		}

	}

	if err != nil {
		response.BadResponse(c, utils.MessageBox["binding_error"], user, meta)
		return
	}

	restult := databse.DB.Create(&user)
	log.Println(restult.RowsAffected, restult.Error)
	if restult.Error != nil {
		response.InternalErrorResponse(c, restult.Error.Error(), user, meta)
		return
	}

	response.SuccessResponse(c, utils.MessageBox["users_created"], user, meta)
}

func GetUsers(c *gin.Context) {
	users := []models.User{}
	id := c.Param("id")
	meta := utils.Meta{}
	if id != "/" {
		id = id[1:]
		result := databse.DB.Find(&users, id)

		response.SuccessResponse(c, utils.MessageBox["user_fetch"], result, meta)
		return
	}
	page := utils.StrToInt(c.Query("page"), 1)
	limit := utils.StrToInt(c.Query("limit"), 10)
	databse.DB.Limit(limit).Offset((page - 1) * limit).Find(&users)
	meta.Page = page
	meta.Limit = limit
	var count int64
	user := models.User{}
	resultData := databse.DB.Model(&user).Count(&count)
	if resultData.Error != nil {
		meta.Total = int64(len(users))
	}

	response.SuccessResponse(c, utils.MessageBox["user_fetch"], users, meta)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	isExist := databse.DB.First(&user, id)
	if isExist.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "User not exist!",
			"error":   isExist.Error,
		})
		return
	}

	c.Bind(&user)
	user.ID = uint(utils.StrToInt(id, 0))
	if user.Password != "" {
		var errPass error
		user.Password, errPass = utils.HashPassword(user.Password)
		if errPass != nil {
			response.BadResponse(c, utils.MessageBox["password_issue"], user, utils.Meta{})
			return
		}

	}
	result := databse.DB.Save(&user)
	if result.RowsAffected > 0 {
		response.SuccessResponse(c, utils.MessageBox["users_updated"], user, utils.Meta{})
		return
	}

	response.BadResponse(c, utils.MessageBox["user_update_error"], user, utils.Meta{})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	isExist := databse.DB.First(&user, id)
	if isExist.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "User not found!.",
			"data":    []models.User{},
		})
		return
	}
	result := databse.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"status":      "Ok",
		"message":     "Users has been deleted successfully.",
		"data":        user,
		"rowAffacted": result.RowsAffected,
	})
}

func Login(c *gin.Context) {
	user := models.User{}
	c.Bind(&user)
	pass := user.Password
	result := databse.DB.First(&user, "email=?", user.Email)
	if result.RowsAffected > 0 {
		err := utils.CheckPassword(user.Password, pass)
		if err != nil {
			response.UnAuthorizeResponse(c, utils.MessageBox["login_fail"], user, utils.Meta{})
			return
		}
		tok := make(map[string]string, 2)

		token, errT := utils.GenrateToken(user.ID)
		tok["token"] = token
		log.Println(token)
		if errT != nil {
			log.Println("Log issue in token", errT.Error())
		}
		response.SuccessResponse(c, utils.MessageBox["login_success"], tok, utils.Meta{})
		return
	}
	response.UnAuthorizeResponse(c, utils.MessageBox["login_fail"], user, utils.Meta{})
}
