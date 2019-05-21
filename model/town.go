package model

type Town struct {
TownId int `gorm:"primary_key;AUTO_INCREMENT" json:"town_id"`
Name int `json:"name"`
CreatedAt int `json:"created_at"`
}
func GetTownById(id int) (Town, bool) {
	var town Town
	ok := db.Where("town_id = ?", id).First(&town).RecordNotFound()

	return town, ok
}

func GetTownByOne(condition []string) (Town, bool) {
	var town Town
	ok := db.Where(condition).First(&town).RecordNotFound()

	return town, ok
}

func GetTownList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	towns := make([]*Town, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&towns).
		Error

	return towns, err
}

func GetTownCount(condition []string) int {
	var town Town
	count := 0
	db.Where(condition).Find(&town).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetTownList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetTownCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteTownByIds(ids string) (error) {
	var town Town
	err := db.Where("id IN  ?", ids).Delete(&town).Error

	return err
}

func DeleteTownById(id int64) error {
	var town Town
	town.TownId = id
	err := db.Delete(&town).Error

	return err
}

func CreateTown(town Town) (Town, error) {
	err := db.Create(&town).Error

	return town, err
}

func UpdateTown(condition []string, params map[string]interface{}) error {
	var town Town
	err := db.Model(&town).Where(condition).Updates(params).Error

	return err
}

func SaveTown(town Town) error {
	err := db.Save(&town).Error

	return err
}