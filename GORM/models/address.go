package models

type Address struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	City    string `json:"city"`
	Api_key string `gorm:"foreignkey:Api_key json:"api_key"`
}

func GetAllAddresses(autoPreload bool) (addresses []Address, err error) {
	addresses = []Address{}
	if autoPreload == true {
		err = DB.Set("gorm:auto_preload", true).Find(&addresses).Error
	} else {
		err = DB.Find(&addresses).Error
	}

	if err == nil {
		return addresses, nil
	}
	return nil, err
}

func CreateAddress(address *Address) (id string, err error) {

	if err = DB.Create(address).Error; err == nil {
		return string(address.Street), nil
	}
	return "", err
}
