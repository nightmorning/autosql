package model

type HomeManage struct {
ManageId int `gorm:"primary_key;AUTO_INCREMENT" json:"manage_id"`
MemberId int `json:"member_id"`
Value int `json:"value"`
CreatedAt int `json:"created_at"`
UpdatedAt int `json:"updated_at"`
}
func GetHomeManageById(id int) (HomeManage, bool) {
	var home_manage HomeManage
	ok := db.First(&home_manage, id).RecordNotFound()

	return home_manage, ok
}

func GetHomeManageByOne(condition map[string]interface{}) (HomeManage, bool) {
	var home_manage HomeManage
	ok := db.Where(condition).First(&home_manage).RecordNotFound()

	return home_manage, ok
}

func GetHomeManageList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	home_manages := make([]*HomeManage, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&home_manages).
		Error

	return home_manages, err
}

func GetHomeManageCount(condition map[string]interface{}) int {
	var home_manage HomeManage
	count := 0
	db.Where(condition).Find(&home_manage).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetHomeManageList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetHomeManageCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteHomeManageByIds(ids string) (error) {
	var home_manage HomeManage
	err := db.Where("id IN  ?", ids).Delete(&home_manage).Error

	return err
}

func DeleteHomeManageById(id int64) error {
	var home_manage HomeManage
	home_manage.HomeManageId = id
	err := db.Delete(&home_manage).Error

	return err
}

func CreateHomeManage(home_manage HomeManage) (HomeManage, error) {
	err := db.Create(&home_manage).Error

	return home_manage, err
}

func UpdateHomeManage(condition map[string]interface{}, params map[string]interface{}) error {
	var home_manage HomeManage
	err := db.Model(&home_manage).Where(condition).Updates(params).Error

	return err
}

func SaveHomeManage(home_manage HomeManage) error {
	err := db.Save(&home_manage).Error

	return err
}