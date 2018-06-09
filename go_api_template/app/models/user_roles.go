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

func GetUserRoleById(Api_key string, autoPreload bool) (user *User, err error) {
	user = &User{}
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
