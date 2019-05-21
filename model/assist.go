package model

type Assist struct {
AssistId int `gorm:"primary_key;AUTO_INCREMENT" json:"assist_id"`
TownId int `json:"town_id"`
VillageId int `json:"village_id"`
Name int `json:"name"`
Mobile int `json:"mobile"`
CreatedAt int `json:"created_at"`
}
func GetAssistById(id int) (Assist, bool) {
	var assist Assist
	ok := db.First(&assist, id).RecordNotFound()

	return assist, ok
}

func GetAssistByOne(condition map[string]interface{}) (Assist, bool) {
	var assist Assist
	ok := db.Where(condition).First(&assist).RecordNotFound()

	return assist, ok
}

func GetAssistList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	assists := make([]*Assist, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&assists).
		Error

	return assists, err
}

func GetAssistCount(condition map[string]interface{}) int {
	var assist Assist
	count := 0
	db.Where(condition).Find(&assist).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetAssistList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetAssistCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteAssistByIds(ids string) (error) {
	var assist Assist
	err := db.Where("id IN  ?", ids).Delete(&assist).Error

	return err
}

func DeleteAssistById(id int64) error {
	var assist Assist
	assist.AssistId = id
	err := db.Delete(&assist).Error

	return err
}

func CreateAssist(assist Assist) (Assist, error) {
	err := db.Create(&assist).Error

	return assist, err
}

func UpdateAssist(condition map[string]interface{}, params map[string]interface{}) error {
	var assist Assist
	err := db.Model(&assist).Where(condition).Updates(params).Error

	return err
}

func SaveAssist(assist Assist) error {
	err := db.Save(&assist).Error

	return err
}