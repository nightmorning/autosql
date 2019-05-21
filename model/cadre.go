package model

type Cadre struct {
CadreId int `gorm:"primary_key;AUTO_INCREMENT" json:"cadre_id"`
TownId int `json:"town_id"`
Name int `json:"name"`
Mobile int `json:"mobile"`
Job int `json:"job"`
JobName int `json:"job_name"`
Face int `json:"face"`
Region int `json:"region"`
Sort int `json:"sort"`
CreatedAt int `json:"created_at"`
}
func GetCadreById(id int) (Cadre, bool) {
	var cadre Cadre
	ok := db.Where("cadre_id = ?", id).First(&cadre).RecordNotFound()

	return cadre, ok
}

func GetCadreByOne(condition []string) (Cadre, bool) {
	var cadre Cadre
	ok := db.Where(condition).First(&cadre).RecordNotFound()

	return cadre, ok
}

func GetCadreList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	cadres := make([]*Cadre, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&cadres).
		Error

	return cadres, err
}

func GetCadreCount(condition []string) int {
	var cadre Cadre
	count := 0
	db.Where(condition).Find(&cadre).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetCadreList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetCadreCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteCadreByIds(ids string) (error) {
	var cadre Cadre
	err := db.Where("id IN  ?", ids).Delete(&cadre).Error

	return err
}

func DeleteCadreById(id int64) error {
	var cadre Cadre
	cadre.CadreId = id
	err := db.Delete(&cadre).Error

	return err
}

func CreateCadre(cadre Cadre) (Cadre, error) {
	err := db.Create(&cadre).Error

	return cadre, err
}

func UpdateCadre(condition []string, params map[string]interface{}) error {
	var cadre Cadre
	err := db.Model(&cadre).Where(condition).Updates(params).Error

	return err
}

func SaveCadre(cadre Cadre) error {
	err := db.Save(&cadre).Error

	return err
}