package main

import (
	"fmt"
	) 

	type User struct {
		ID int
		FirstName string
		LastName string
		Email string
		IsActive bool
	}

	type Group struct {
		Name string
		Admin User
		Users []User
		IsAvailable bool
	}
func main() {
	// // initialize struct of User
	user := User{}
	
	// // method insert data into struct 1
	user.ID = 1
	user.FirstName = "Dedi"
	user.LastName = "Fardiyanto"
	user.Email = "dedif15@gmail.com"
	user.IsActive = true

	// // method insert data into struct 2
	user2 := User{ID: 2, FirstName: "Dedi", LastName: "Fard", Email: "dedif15@gmail.com", IsActive: true}
	fmt.Println("User Struct-------------------------------------------------------")
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println("-------------------------------------------------------")

	fmt.Println(displayName(user))
	fmt.Println(displayName(user2))

	fmt.Println("Group Struct-------------------------------------------------------")
	users := []User{user, user2}
	group := Group{Name: "Group 1", Admin: user, Users: users, IsAvailable: true}

	displayGroup(group)
}

// struct as params func
func displayName(user User) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group Group) {
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("Name Of Group: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Admin Of Group: %s", group.Admin.FirstName)
	fmt.Println("")
	fmt.Printf("Member Count: %d", len(group.Users))
	
	fmt.Println("")
	fmt.Println("")
	fmt.Println("User Of Group: ")
	for _, u := range group.Users {
	fmt.Println(u.FirstName + " " + u.LastName)
	}
}