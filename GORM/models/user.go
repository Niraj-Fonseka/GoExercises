package models

type UserData struct {
	Users []User
}

type UserInterface interface {
	GetAllUsers(autoPreload bool) ([]User, error)
}
type User struct {
	Api_key    string `gorm:"size:100;unique" json:"api_key"`
	Api_secret string `gorm:"size:100" json:"api_secret"`
	//Age int
	// Name       string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	// Num        int    `gorm:"AUTO_INCREMENT"`
	// IgnoreMe   int    `gorm:"-"` // Ignore this field

}

//TableName Creates table with give name
func (User) TableName() string {
	return "user"
}

//CreateUser creates data
func CreateUser(user *User) (id string, err error) {

	if err = DB.Create(user).Error; err == nil {
		return string(user.Api_key), nil
	}
	return "", err
}

//ModifyUser modifies data
func ModifyUser(user *User) (err error) {

	if err = DB.Save(user).Error; err == nil {
		return nil
	}
	return err
}

//GetUsertByID gets User by id
//keep autoPreload true if needs to fetch all related data else keep false
func GetUsertByID(id int, autoPreload bool) (user *User, err error) {
	user = &User{}
	if autoPreload == true {
		if err = DB.Set("gorm:auto_preload", true).First(user, id).Error; err != nil {

			return nil, err
		}
	} else {
		if err = DB.First(user, id).Error; err != nil {

			return nil, err
		}
	}
	return user, nil
}

//GetUserByFeild : can change the method according to requierment
func GetUserByFeild(feildName string, searchValue string) (user *User, err error) {
	user = &User{}
	// use DB.Unscoped() insted of DB to find soft deleted entries
	if err = DB.Where(feildName+" = ?", searchValue).First(user).Error; err != nil {

		return nil, err
	}

	return user, nil
}

//GetAllUsers to get all data
//keep autoPreload true if needs to fetch all related data else keep false
func (u UserData) GetAllUsers(autoPreload bool) (users []User, err error) {
	//users = []User{}
	users = u.Users
	if autoPreload == true {
		err = DB.Set("gorm:auto_preload", true).Find(&users).Error
	} else {
		err = DB.Find(&users).Error
	}

	if err == nil {
		return users, nil
	}
	return nil, err
}

// //DeleteUser deletes data from database
// // for soft delete keep softdelete true
// func DeleteUser(user User, softDelete bool) (err error) {

// 	// Delete record permanently with Unscoped
// 	if softDelete == false {
// 		//db.Unscoped().Where("age = ?", 20).Delete(&User{})

// 		err = DB.Unscoped().Where("api_key = ?", user.Api_key).Delete(&user).Error
// 	} else {
// 		err = DB.Delete(&user).Error
// 		// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id

// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
