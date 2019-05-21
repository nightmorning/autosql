package model

type Member struct {
MemberId int `gorm:"primary_key;AUTO_INCREMENT" json:"member_id"`
TownId int `json:"town_id"`
VillageId int `json:"village_id"`
GroupId int `json:"group_id"`
Name int `json:"name"`
Mobile int `json:"mobile"`
Password int `json:"password"`
Token int `json:"token"`
Type int `json:"type"`
IsCadre int `json:"is_cadre"`
Status int `json:"status"`
CreatedAt int `json:"created_at"`
}
func GetMemberById(id int) (Member, bool) {
	var member Member
	ok := db.Where("member_id = ?", id).First(&member).RecordNotFound()

	return member, ok
}

func GetMemberByOne(condition []string) (Member, bool) {
	var member Member
	ok := db.Where(condition).First(&member).RecordNotFound()

	return member, ok
}

func GetMemberList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	members := make([]*Member, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&members).
		Error

	return members, err
}

func GetMemberCount(condition []string) int {
	var member Member
	count := 0
	db.Where(condition).Find(&member).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetMemberList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetMemberCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteMemberByIds(ids string) (error) {
	var member Member
	err := db.Where("id IN  ?", ids).Delete(&member).Error

	return err
}

func DeleteMemberById(id int64) error {
	var member Member
	member.MemberId = id
	err := db.Delete(&member).Error

	return err
}

func CreateMember(member Member) (Member, error) {
	err := db.Create(&member).Error

	return member, err
}

func UpdateMember(condition []string, params map[string]interface{}) error {
	var member Member
	err := db.Model(&member).Where(condition).Updates(params).Error

	return err
}

func SaveMember(member Member) error {
	err := db.Save(&member).Error

	return err
}