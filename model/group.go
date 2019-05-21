package model

type Group struct {
GroupId int `gorm:"primary_key;AUTO_INCREMENT" json:"group_id"`
TownId int `json:"town_id"`
VillageId int `json:"village_id"`
Name int `json:"name"`
MemberId int `json:"member_id"`
Mobile int `json:"mobile"`
CreatedAt int `json:"created_at"`
}
func GetGroupById(id int) (Group, bool) {
	var group Group
	ok := db.Where("group_id = ?", id).First(&group).RecordNotFound()

	return group, ok
}

func GetGroupByOne(condition []string) (Group, bool) {
	var group Group
	ok := db.Where(condition).First(&group).RecordNotFound()

	return group, ok
}

func GetGroupList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	groups := make([]*Group, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&groups).
		Error

	return groups, err
}

func GetGroupCount(condition []string) int {
	var group Group
	count := 0
	db.Where(condition).Find(&group).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetGroupList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetGroupCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteGroupByIds(ids string) (error) {
	var group Group
	err := db.Where("id IN  ?", ids).Delete(&group).Error

	return err
}

func DeleteGroupById(id int64) error {
	var group Group
	group.GroupId = id
	err := db.Delete(&group).Error

	return err
}

func CreateGroup(group Group) (Group, error) {
	err := db.Create(&group).Error

	return group, err
}

func UpdateGroup(condition []string, params map[string]interface{}) error {
	var group Group
	err := db.Model(&group).Where(condition).Updates(params).Error

	return err
}

func SaveGroup(group Group) error {
	err := db.Save(&group).Error

	return err
}