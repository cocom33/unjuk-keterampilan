package controllers

import (
	"errors"
	"net/http"
	"project/config"
	"project/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StudentGetData(c echo.Context) error {
	DB := config.GetDBInstance()
	var students []models.Student

	result := DB.Find(&students)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Message: "Failed get data students",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success get data students",
		Data: students,
	})
}

func StudentGetDetailData(c echo.Context) error {
	DB := config.GetDBInstance()
	var student models.Student
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status: false,
			Message: "Id tidak valid",
		})
	}
	
	result := DB.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.Response{
				Status:  false,
				Message: "Data tidak ditemukan",
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Message: "tidak ada data",
		})
	}
	
	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "berhasil mengambil data",
		Data: student,
	})
}

func StudentAddData(c echo.Context) error {
	DB := config.GetDBInstance()
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// students := models.Student{Name: name, Email: email}
	students := models.Student{}
	c.Bind(&students)
	
	// if name == "" || email == "" {
	if students.Name == "" || students.Email == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status: false,
			Message: "gagal membuat data",
		})
	}
	
	res := DB.First(&models.Student{}, "email = ?", students.Email)
	if res.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.Response{
			Status: false,
			Message: "Email sudah ada",
		})
	}
	
	res = DB.Create(&students)
	if err := res.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Message: "gagal membuat data",
		})
	}
	
	return c.JSON(http.StatusCreated, models.Response{
		Status: true,
		Message: "berhasil membuat data",
		Data: students,
	})
}

func StudentDeleteData(c echo.Context) error {
	DB := config.GetDBInstance()
	var student models.Student
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status: false,
			Message: "Id tidak valid",
		})
	}
	
	result := DB.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.Response{
				Status:  false,
				Message: "Data tidak ditemukan",
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Message: "tidak ada data",
		})
	}
	
	DB.Delete(&student, id)
	
	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "berhasil menghapus data",
	})
}

func StudentUpdateData(c echo.Context) error {
	DB := config.GetDBInstance()
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	var students models.Student
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status: false,
			Message: "Id tidak valid",
		})
	}
	
	// student := models.Student{Id: id, Name: name, Email: email}
	student := models.Student{}
  if err := c.Bind(&student); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Status:  false,
      Message: "Data tidak valid",
    })
  }
  
  result := DB.First(&students, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.Response{
				Status:  false,
				Message: "Data tidak ditemukan",
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: false,
			Message: "tidak ada data",
		})
	}
	
	result = DB.Model(&students).Where("id = ?", id).Updates(&student)
	    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, models.Response{
            Status:  false,
            Message: "Gagal mengupdate data",
        })
    }
	
	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "berhasil memperbarui data",
		Data: student,
	})
}