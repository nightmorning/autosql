package model

type MobileCode struct {
CodeId int `gorm:"primary_key;AUTO_INCREMENT" json:"code_id"`
Mobile int `json:"mobile"`
Code int `json:"code"`
Type int `json:"type"`
Date int `json:"date"`
CreatedAt int `json:"created_at"`
}
func GetMobileCodeById(id int) (MobileCode, bool) {
	var mobile_code MobileCode
	ok := db.First(&mobile_code, id).RecordNotFound()

	return mobile_code, ok
}

func GetMobileCodeByOne(condition map[string]interface{}) (MobileCode, bool) {
	var mobile_code MobileCode
	ok := db.Where(condition).First(&mobile_code).RecordNotFound()

	return mobile_code, ok
}

func GetMobileCodeList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	mobile_codes := make([]*MobileCode, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&mobile_codes).
		Error

	return mobile_codes, err
}

func GetMobileCodeCount(condition map[string]interface{}) int {
	var mobile_code MobileCode
	count := 0
	db.Where(condition).Find(&mobile_code).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetMobileCodeList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetMobileCodeCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteMobileCodeByIds(ids string) (error) {
	var mobile_code MobileCode
	err := db.Where("id IN  ?", ids).Delete(&mobile_code).Error

	return err
}

func DeleteMobileCodeById(id int64) error {
	var mobile_code MobileCode
	mobile_code.MobileCodeId = id
	err := db.Delete(&mobile_code).Error

	return err
}

func CreateMobileCode(mobile_code MobileCode) (MobileCode, error) {
	err := db.Create(&mobile_code).Error

	return mobile_code, err
}

func UpdateMobileCode(condition map[string]interface{}, params map[string]interface{}) error {
	var mobile_code MobileCode
	err := db.Model(&mobile_code).Where(condition).Updates(params).Error

	return err
}

func SaveMobileCode(mobile_code MobileCode) error {
	err := db.Save(&mobile_code).Error

	return err
}