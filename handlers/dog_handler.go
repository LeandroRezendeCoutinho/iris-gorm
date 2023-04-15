package handlers

import (
	config "main/configs"
	"main/internal/entities"

	"github.com/kataras/iris/v12"
)

func GetDogs(ctx iris.Context) {
	var dogs []entities.Dog
	config.Database.Find(&dogs)

	ctx.JSON(dogs)
}

func GetDog(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var dog entities.Dog
	result := config.Database.First(&dog, id)

	if result.RowsAffected == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.JSON(dog)
}

func AddDog(ctx iris.Context) {
	var dog entities.Dog

	if err := ctx.ReadJSON(&dog); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	ctx.ReadJSON(&dog)
	config.Database.Create(&dog)

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(dog)
}

func UpdateDog(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var dog entities.Dog
	result := config.Database.First(&dog, id)

	if result.RowsAffected == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	if err := ctx.ReadJSON(&dog); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	config.Database.Where("id = ?", id).Updates(&dog)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(dog)
}

func DeleteDog(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var dog entities.Dog
	result := config.Database.First(&dog, id)

	if result.RowsAffected == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	config.Database.Delete(&dog)
	ctx.StatusCode(iris.StatusOK)
}
