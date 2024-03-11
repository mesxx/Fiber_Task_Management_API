package handlers

import (
	"strconv"

	"github.com/mesxx/Fiber_Task_Management_API/helpers"
	"github.com/mesxx/Fiber_Task_Management_API/models"
	"github.com/mesxx/Fiber_Task_Management_API/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	TaskHandler interface {
		Create(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetAllByUser(c *fiber.Ctx) error
		GetByID(c *fiber.Ctx) error
		GetImageByID(c *fiber.Ctx) error
		UpdateByID(c *fiber.Ctx) error
		DeleteImageByID(c *fiber.Ctx) error
		DeleteByID(c *fiber.Ctx) error
		DeleteAll(c *fiber.Ctx) error
	}

	taskHandler struct {
		TaskUsecase usecases.TaskUsecase
	}
)

func NewTaskHandler(usecase usecases.TaskUsecase) TaskHandler {
	return &taskHandler{
		TaskUsecase: usecase,
	}
}

func (handler taskHandler) Create(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	// START request
	var requestCreateTask models.RequestCreateTask
	if err := c.BodyParser(&requestCreateTask); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	requestCreateTask.UserID = userSigned.ID

	// START request file
	fileImage, err1 := c.FormFile("image")
	if err1 == nil {
		// START check type
		fileType := fileImage.Header.Get("Content-Type")
		if err := helpers.UploadSettingType(fileType); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END check type

		// START set filename
		fileName, err := helpers.UploadSettingName(fileImage.Filename)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END set filename

		// START save file
		destination := "./publics/images/" + fileName
		if err := c.SaveFile(fileImage, destination); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END save file

		requestCreateTask.Image = fileName
	}
	// END request file

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestCreateTask); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	res, err2 := handler.TaskUsecase.Create(&requestCreateTask)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GetResponseData(fiber.StatusCreated, "success", res))
}

func (handler taskHandler) GetAll(c *fiber.Ctx) error {
	tasks, err := handler.TaskUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", tasks))
}

func (handler taskHandler) GetAllByUser(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	tasks, err := handler.TaskUsecase.GetAllByUser(userSigned.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", tasks))
}

func (handler taskHandler) GetByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	task, err2 := handler.TaskUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", task))
}

func (handler taskHandler) GetImageByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	task, err2 := handler.TaskUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START set path
	fileName := task.Image.String
	destination := "./publics/images/" + fileName
	// END set path

	return c.Status(fiber.StatusOK).SendFile(destination)
}

func (handler taskHandler) UpdateByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	task, err2 := handler.TaskUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START request
	var requestUpdateByIDTask models.RequestUpdateTask
	if err := c.BodyParser(&requestUpdateByIDTask); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START map request
	var values = make(map[string]interface{})
	// END map request

	if requestUpdateByIDTask.Status != "" {
		values["status"] = requestUpdateByIDTask.Status
	}

	if requestUpdateByIDTask.Title != "" {
		values["title"] = requestUpdateByIDTask.Title
	}

	if requestUpdateByIDTask.Description != nil {
		if *requestUpdateByIDTask.Description == "" {
			values["description"] = nil
		} else {
			values["description"] = requestUpdateByIDTask.Description
		}
	}

	// START request file
	fileImage, err3 := c.FormFile("image")
	if err3 == nil {
		// START check type
		fileType := fileImage.Header.Get("Content-Type")
		if err := helpers.UploadSettingType(fileType); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END check type

		// START set filename
		fileName, err := helpers.UploadSettingName(fileImage.Filename)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END set filename

		// START delete old file
		if err := helpers.DeleteImage(task.Image.String); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END delete old file

		// START save new file
		destination := "./publics/images/" + fileName
		if err := c.SaveFile(fileImage, destination); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END save new file

		values["image"] = fileName
	}
	// END request file

	updateByID, err4 := handler.TaskUsecase.UpdateByID(task, values)
	if err4 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err4.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler taskHandler) DeleteImageByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	task, err2 := handler.TaskUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START delete file
	if err := helpers.DeleteImage(task.Image.String); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete file

	// START map request
	var values = make(map[string]interface{})
	values["image"] = nil
	// END map request

	updateByID, err3 := handler.TaskUsecase.UpdateByID(task, values)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler taskHandler) DeleteByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	task, err2 := handler.TaskUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START delete file
	if err := helpers.DeleteImage(task.Image.String); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete file

	deleteByID, err3 := handler.TaskUsecase.DeleteByID(task)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", deleteByID))
}

func (handler taskHandler) DeleteAll(c *fiber.Ctx) error {
	// START delete all file
	if err := helpers.DeleteAllImage(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete all file

	if err := handler.TaskUsecase.DeleteAll(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponse(fiber.StatusOK, "success"))
}
