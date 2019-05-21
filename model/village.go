package model

type Village struct {
 VillageId int `gorm:"primary_key;AUTO_INCREMENT" json:" village_id"`
TownId int `json:"town_id"`
Name int `json:"name"`
Sort int `json:"sort"`
CreatedAt int `json:"created_at"`
}
func GetVillageById(id int) (Village, bool) {
	var village Village
	ok := db.Where(" village_id = ?", id).First(&village).RecordNotFound()

	return village, ok
}

func GetVillageByOne(condition []string) (Village, bool) {
	var village Village
	ok := db.Where(condition).First(&village).RecordNotFound()

	return village, ok
}

func GetVillageList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	villages := make([]*Village, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&villages).
		Error

	return villages, err
}

func GetVillageCount(condition []string) int {
	var village Village
	count := 0
	db.Where(condition).Find(&village).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetVillageList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetVillageCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteVillageByIds(ids string) (error) {
	var village Village
	err := db.Where("id IN  ?", ids).Delete(&village).Error

	return err
}

func DeleteVillageById(id int64) error {
	var village Village
	village.VillageId = id
	err := db.Delete(&village).Error

	return err
}

func CreateVillage(village Village) (Village, error) {
	err := db.Create(&village).Error

	return village, err
}

func UpdateVillage(condition []string, params map[string]interface{}) error {
	var village Village
	err := db.Model(&village).Where(condition).Updates(params).Error

	return err
}

func SaveVillage(village Village) error {
	err := db.Save(&village).Error

	return err
}