package model

type Guide struct {
GuideId int `gorm:"primary_key;AUTO_INCREMENT" json:"guide_id"`
TownId int `json:"town_id"`
StationId int `json:"station_id"`
Name int `json:"name"`
CreatedAt int `json:"created_at"`
}
func GetGuideById(id int) (Guide, bool) {
	var guide Guide
	ok := db.Where("guide_id = ?", id).First(&guide).RecordNotFound()

	return guide, ok
}

func GetGuideByOne(condition []string) (Guide, bool) {
	var guide Guide
	ok := db.Where(condition).First(&guide).RecordNotFound()

	return guide, ok
}

func GetGuideList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	guides := make([]*Guide, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&guides).
		Error

	return guides, err
}

func GetGuideCount(condition []string) int {
	var guide Guide
	count := 0
	db.Where(condition).Find(&guide).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetGuideList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetGuideCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteGuideByIds(ids string) (error) {
	var guide Guide
	err := db.Where("id IN  ?", ids).Delete(&guide).Error

	return err
}

func DeleteGuideById(id int64) error {
	var guide Guide
	guide.GuideId = id
	err := db.Delete(&guide).Error

	return err
}

func CreateGuide(guide Guide) (Guide, error) {
	err := db.Create(&guide).Error

	return guide, err
}

func UpdateGuide(condition []string, params map[string]interface{}) error {
	var guide Guide
	err := db.Model(&guide).Where(condition).Updates(params).Error

	return err
}

func SaveGuide(guide Guide) error {
	err := db.Save(&guide).Error

	return err
}