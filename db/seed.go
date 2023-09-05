package db

import (
	"fmt"
	"go-ecom/models"
)

func Seed() {
	users := []models.Users{
		{
			Name:  "first",
			Email: "first@mail.com",
			Phone: "2135156214",
			Password: "first_password",
		},
		{
			Name:  "second",
			Email: "second@mail.com",
			Phone: "5626212",
			Password: "second_password",
		},
		{
			Name:  "third",
			Email: "third@mail.com",
			Phone: "52873921",
			Password: "third_password",
		},
		{
			Name:  "fourth",
			Email: "fourth@mail.com",
			Phone: "628967847",
			Password: "fourth_password",
		},
	}

	orders := []models.Orders{
		{
			UsersID: 1,
			Status: "вработе",
			OrdersItems: []models.OrdersItems{
				{ItemsID: 1, Quantity: 10},
				{ItemsID: 4, Quantity: 10},
				{ItemsID: 5, Quantity: 10},
			},
		},
		{
			UsersID: 1,
			Status: "вработе",
			OrdersItems: []models.OrdersItems{
				{ItemsID: 1, Quantity: 10},
				{ItemsID: 4, Quantity: 10},
				{ItemsID: 3, Quantity: 10},
			},
		},
		{
			UsersID: 2,
			Status: "вдоставке",
			OrdersItems: []models.OrdersItems{
				{ItemsID: 5, Quantity: 10},
				{ItemsID: 3, Quantity: 10},
			},
		},
		{
			UsersID: 3,
			Status: "навыдачи",
			OrdersItems: []models.OrdersItems{
				{ItemsID: 2, Quantity: 10},
			},
		},
		{
			UsersID: 4,
			Status: "готовится",
			OrdersItems: []models.OrdersItems{
				{ItemsID: 1, Quantity: 5},
				{ItemsID: 2, Quantity: 5},
				{ItemsID: 3, Quantity: 5},
			},
		},
}

	items := []models.Items{
		{
			Name: "кола",
			Price: 50,
		},
		{
			Name: "шаурма",
			Price: 250,
		},
		{
			Name: "салат",
			Price: 150,
		},
		{
			Name: "пончик",
			Price: 100,
		},
		{
			Name: "чай",
			Price: 30,
		},
	}

	DB.Model(models.Users{}).Create(&users)
	DB.Model(models.Items{}).Create(&items)
	DB.Model(models.Orders{}).Create(&orders)


	fmt.Printf("USERS, ORDERS, ITEMS,  created in DB \n")

}
