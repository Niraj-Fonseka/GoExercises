package models

//User struct
type UserRoles struct {
	Api_key string `gorm:"foreignkey:Api_key json:"api_key"`
	Role    string `gorm:"size:100" json:"role"`
	//Age int
	// Name       string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	// Num        int    `gorm:"AUTO_INCREMENT"`
	// IgnoreMe   int    `gorm:"-"` // Ignore this field

}

//TableName Creates table with give name
func (UserRoles) TableName() string {
	return "userrole"
}

//GetAllUsers to get all data
//keep autoPreload true if needs to fetch all related data else keep false
func GetAllUserRoles(autoPreload bool) (userRoles []UserRoles, err error) {
	userRoles = []UserRoles{}
	if autoPreload == true {
		err = DB.Set("gorm:auto_preload", true).Find(&userRoles).Error
	} else {
		err = DB.Find(&userRoles).Error
	}

	if err == nil {
		return userRoles, nil
	}
	return nil, err
}

// //GetUserByFeild : can change the method according to requierment
// func GetUserByFeild(feildName string, searchValue string) (user *User, err error) {
// 	user = &UserRoles{}
// 	// use DB.Unscoped() insted of DB to find soft deleted entries
// 	if err = DB.Where(feildName+" = ?", searchValue).First(user).Error; err != nil {

// 		return nil, err
// 	}

// 	return user, nil
// }

func GetUserRoleById(Api_key string, autoPreload bool) (user *UserRoles, err error) {
	user = &UserRoles{}
	if autoPreload == true {
		if err = DB.Set("gorm:auto_preload", true).First(user, Api_key).Error; err != nil {

			return nil, err
		}
	} else {
		if err = DB.First(user, Api_key).Error; err != nil {

			return nil, err
		}
	}
	return user, nil
}

func CreateUserRole(role *UserRoles) (id string, err error) {

	if err = DB.Create(role).Error; err == nil {
		return string(role.Api_key), nil
	}
	return "", err
}
