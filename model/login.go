package model

type Login struct {
LoginId int `gorm:"primary_key;AUTO_INCREMENT" json:"login_id"`
Mobile int `json:"mobile"`
Date int `json:"date"`
Number int `json:"number"`
CreatedAt int `json:"created_at"`
}
func GetLoginById(id int) (Login, bool) {
	var login Login
	ok := db.Where("login_id = ?", id).First(&login).RecordNotFound()

	return login, ok
}

func GetLoginByOne(condition []string) (Login, bool) {
	var login Login
	ok := db.Where(condition).First(&login).RecordNotFound()

	return login, ok
}

func GetLoginList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	logins := make([]*Login, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&logins).
		Error

	return logins, err
}

func GetLoginCount(condition []string) int {
	var login Login
	count := 0
	db.Where(condition).Find(&login).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetLoginList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetLoginCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteLoginByIds(ids string) (error) {
	var login Login
	err := db.Where("id IN  ?", ids).Delete(&login).Error

	return err
}

func DeleteLoginById(id int64) error {
	var login Login
	login.LoginId = id
	err := db.Delete(&login).Error

	return err
}

func CreateLogin(login Login) (Login, error) {
	err := db.Create(&login).Error

	return login, err
}

func UpdateLogin(condition []string, params map[string]interface{}) error {
	var login Login
	err := db.Model(&login).Where(condition).Updates(params).Error

	return err
}

func SaveLogin(login Login) error {
	err := db.Save(&login).Error

	return err
}