package model

type Business struct {
BusinessId int `gorm:"primary_key;AUTO_INCREMENT" json:"business_id"`
TownId int `json:"town_id"`
Content int `json:"content"`
Date int `json:"date"`
AuthorId int `json:"author_id"`
Type int `json:"type"`
CreatedAt int `json:"created_at"`
}
func GetBusinessById(id int) (Business, bool) {
	var business Business
	ok := db.First(&business, id).RecordNotFound()

	return business, ok
}

func GetBusinessByOne(condition map[string]interface{}) (Business, bool) {
	var business Business
	ok := db.Where(condition).First(&business).RecordNotFound()

	return business, ok
}

func GetBusinessList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	businesss := make([]*Business, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&businesss).
		Error

	return businesss, err
}

func GetBusinessCount(condition map[string]interface{}) int {
	var business Business
	count := 0
	db.Where(condition).Find(&business).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetBusinessList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetBusinessCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteBusinessByIds(ids string) (error) {
	var business Business
	err := db.Where("id IN  ?", ids).Delete(&business).Error

	return err
}

func DeleteBusinessById(id int64) error {
	var business Business
	business.BusinessId = id
	err := db.Delete(&business).Error

	return err
}

func CreateBusiness(business Business) (Business, error) {
	err := db.Create(&business).Error

	return business, err
}

func UpdateBusiness(condition map[string]interface{}, params map[string]interface{}) error {
	var business Business
	err := db.Model(&business).Where(condition).Updates(params).Error

	return err
}

func SaveBusiness(business Business) error {
	err := db.Save(&business).Error

	return err
}