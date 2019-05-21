package model

type MemberInfo struct {
MemberInfoId int `gorm:"primary_key;AUTO_INCREMENT" json:"member_info_id"`
MemberId int `json:"member_id"`
Military int `json:"military"`
PoliticalStatus int `json:"political_status"`
Birthday int `json:"birthday"`
PartyRepresent int `json:"party_represent"`
PersonRepresent int `json:"person_represent"`
WomenRepresent int `json:"women_represent"`
VillagePost int `json:"village_post"`
Follow int `json:"follow"`
MedicalInsurance int `json:"medical_insurance"`
Remarks int `json:"remarks"`
Level int `json:"level"`
Address int `json:"address"`
Qrcode int `json:"qrcode"`
CreatedAt int `json:"created_at"`
UpdatedAt int `json:"updated_at"`
}
func GetMemberInfoById(id int) (MemberInfo, bool) {
	var member_info MemberInfo
	ok := db.Where("member_info_id = ?", id).First(&member_info).RecordNotFound()

	return member_info, ok
}

func GetMemberInfoByOne(condition []string) (MemberInfo, bool) {
	var member_info MemberInfo
	ok := db.Where(condition).First(&member_info).RecordNotFound()

	return member_info, ok
}

func GetMemberInfoList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	member_infos := make([]*MemberInfo, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&member_infos).
		Error

	return member_infos, err
}

func GetMemberInfoCount(condition []string) int {
	var member_info MemberInfo
	count := 0
	db.Where(condition).Find(&member_info).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetMemberInfoList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetMemberInfoCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteMemberInfoByIds(ids string) (error) {
	var member_info MemberInfo
	err := db.Where("id IN  ?", ids).Delete(&member_info).Error

	return err
}

func DeleteMemberInfoById(id int64) error {
	var member_info MemberInfo
	member_info.MemberInfoId = id
	err := db.Delete(&member_info).Error

	return err
}

func CreateMemberInfo(member_info MemberInfo) (MemberInfo, error) {
	err := db.Create(&member_info).Error

	return member_info, err
}

func UpdateMemberInfo(condition []string, params map[string]interface{}) error {
	var member_info MemberInfo
	err := db.Model(&member_info).Where(condition).Updates(params).Error

	return err
}

func SaveMemberInfo(member_info MemberInfo) error {
	err := db.Save(&member_info).Error

	return err
}