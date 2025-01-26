package controllers

import (
	"todo/internal/models"
	databse "todo/pkg/database"
	"todo/pkg/response"
	"todo/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	todo := &models.Todo{}
	meta := utils.Meta{}

	err := c.Bind(&todo)
	if err != nil {
		response.InternalErrorResponse(c, utils.MessageBox["binding_error"], todo, meta)
		return
	}

	result := databse.DB.Save(&todo)

	if result.Error != nil {
		response.InternalErrorResponse(c, utils.MessageBox["todo_error"], todo, meta)
		return
	}
	response.SuccessResponse(c, utils.MessageBox["todo_created"], todo, meta)
}

func GetTodo(c *gin.Context) {
	todo := &[]models.Todo{}
	meta := utils.Meta{}

	id := c.Param("id")
	if id != "/" {
		id = id[1:]
		databse.DB.Find(&todo, id)
		response.SuccessResponse(c, utils.MessageBox["todo_fetch"], todo, meta)
		return
	}

	page := utils.StrToInt(c.Query("page"), 1)
	limit := utils.StrToInt(c.Query("limit"), 10)

	result := databse.DB.Limit(limit).Offset((page - 1) * limit).Find(&todo)

	if result.Error != nil {
		response.InternalErrorResponse(c, utils.MessageBox["todo_error"], todo, meta)
		return
	}

	var count int64
	databse.DB.Model(&todo).Count(&count)
	meta.Page = page
	meta.Limit = limit
	meta.Total = count

	response.SuccessResponse(c, utils.MessageBox["todo_fetch"], todo, meta)
}

func UpdateTodo(c *gin.Context) {

	todo := &models.Todo{}
	id := c.Param("id")
	meta := &utils.Meta{}
	databse.DB.First(&todo, id)
	if todo.ID == 0 {
		response.BadResponse(c, utils.MessageBox["todo_not_exist"], todo, meta)
		return
	}
	c.Bind(&todo)
	result := databse.DB.Save(&todo)

	if result.RowsAffected > 0 {
		response.SuccessResponse(c, utils.MessageBox["todo_updated"], todo, meta)
		return
	}
	response.BadResponse(c, utils.MessageBox["badRequest"], todo, meta)

}

func DeleteTodo(c *gin.Context) {
	todo := &models.Todo{}
	id := c.Param("id")
	databse.DB.First(&todo)
	if todo.ID == 0 {
		response.BadResponse(c, utils.MessageBox["todo_not_exist"], todo, utils.Meta{})
		return
	}
	result := databse.DB.Delete(&todo, id)
	if result.RowsAffected > 0 {
		response.SuccessResponse(c, utils.MessageBox["todo_deleted"], todo, utils.Meta{})
		return
	}
	response.BadResponse(c, utils.MessageBox["todo_not_exist"], todo, utils.Meta{})
}
