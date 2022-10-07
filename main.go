package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/vikaputri/Gorm/models"

	"github.com/vikaputri/Gorm/database"
)

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data ", err)
		return
	}

	fmt.Println("New user data", User)
}

func getUserByID(id uint) {

	db := database.GetDB()
	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		fmt.Println("Error finding user data:", err)
	}

	fmt.Printf("User data: %+v \n", user)
}

func updateUserByID(id uint, email string) {
	db := database.GetDB()
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data: ", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userID uint, brand, name string) {

	db := database.GetDB()

	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data: ", err.Error())
		return
	}

	fmt.Println("New product data: ", Product)
}

func getUsersWithProducts() {

	db := database.GetDB()

	user := models.User{}

	err := db.Preload("Products").Find(&user).Error

	if err != nil {
		fmt.Println("Error getting user datas with product", err.Error())
		return
	}

	fmt.Println("User Datas with Products")
	fmt.Printf("%+v \n", user)

}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("product with id %d has successfully deleted", id)
}

func main() {
	database.StartDB()

	//createUser("vika@gmail.com")
	//getUserByID(1)
	//updateUserByID(1, "vikaputri@gmail.com")
	//createProduct(1, "YLO", "YIIY")
	//getUsersWithProducts()
	deleteProductById(5)
}
