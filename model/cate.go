package model

type Cate struct {
CateId int `gorm:"primary_key;json:"cate_id"`
ParentId int `json:"parent_id"`
Name int `json:"name"`
Sort int `json:"sort"`
CreatedAt int `json:"created_at"`
}
func GetCateById(id int) (Cate, bool) {
	var cate Cate
	ok := db.First(&cate, id).RecordNotFound()

	return cate, ok
}

func GetCateByOne(condition map[string]interface{}) (Cate, bool) {
	var cate Cate
	ok := db.Where(condition).First(&cate).RecordNotFound()

	return cate, ok
}

func GetCateList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	cates := make([]*Cate, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&cates).
		Error

	return cates, err
}

func GetCateCount(condition map[string]interface{}) int {
	var cate Cate
	count := 0
	db.Where(condition).Find(&cate).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetCateList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetCateCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteCateByIds(ids string) (error) {
	var cate Cate
	err := db.Where("id IN  ?", ids).Delete(&cate).Error

	return err
}

func DeleteCateById(id int64) error {
	var cate Cate
	cate.CateId = id
	err := db.Delete(&cate).Error

	return err
}

func CreateCate(cate Cate) (Cate, error) {
	err := db.Create(&cate).Error

	return cate, err
}

func UpdateCate(condition map[string]interface{}, params map[string]interface{}) error {
	var cate Cate
	err := db.Model(&cate).Where(condition).Updates(params).Error

	return err
}

func SaveCate(cate Cate) error {
	err := db.Save(&cate).Error

	return err
}