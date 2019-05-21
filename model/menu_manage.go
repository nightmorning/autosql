package model

type MenuManage struct {
MenuManageId int `gorm:"primary_key;AUTO_INCREMENT" json:"menu_manage_id"`
TownId int `json:"town_id"`
Name int `json:"name"`
Aliases int `json:"aliases"`
Type int `json:"type"`
CreatedAt int `json:"created_at"`
}
func GetMenuManageById(id int) (MenuManage, bool) {
	var menu_manage MenuManage
	ok := db.Where("menu_manage_id = ?", id).First(&menu_manage).RecordNotFound()

	return menu_manage, ok
}

func GetMenuManageByOne(condition []string) (MenuManage, bool) {
	var menu_manage MenuManage
	ok := db.Where(condition).First(&menu_manage).RecordNotFound()

	return menu_manage, ok
}

func GetMenuManageList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	menu_manages := make([]*MenuManage, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&menu_manages).
		Error

	return menu_manages, err
}

func GetMenuManageCount(condition []string) int {
	var menu_manage MenuManage
	count := 0
	db.Where(condition).Find(&menu_manage).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetMenuManageList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetMenuManageCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteMenuManageByIds(ids string) (error) {
	var menu_manage MenuManage
	err := db.Where("id IN  ?", ids).Delete(&menu_manage).Error

	return err
}

func DeleteMenuManageById(id int64) error {
	var menu_manage MenuManage
	menu_manage.MenuManageId = id
	err := db.Delete(&menu_manage).Error

	return err
}

func CreateMenuManage(menu_manage MenuManage) (MenuManage, error) {
	err := db.Create(&menu_manage).Error

	return menu_manage, err
}

func UpdateMenuManage(condition []string, params map[string]interface{}) error {
	var menu_manage MenuManage
	err := db.Model(&menu_manage).Where(condition).Updates(params).Error

	return err
}

func SaveMenuManage(menu_manage MenuManage) error {
	err := db.Save(&menu_manage).Error

	return err
}