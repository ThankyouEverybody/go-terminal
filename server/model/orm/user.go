package orm

type User struct {
	ID        string `gorm:"primary_key:varchar(36)" json:"id"`
	UserName  string `gorm:"type:varchar(10);comment:用户名" json:"userName"`
	Password  string `gorm:"type:varchar(20);comment:密码" json:"password"`
	LoginName string `gorm:"type:varchar(10);comment:登陆名" json:"loginName"`
	Online    *bool  `gorm:"default:false;comment" json:"online"`
	Mail      string `gorm:"varchar(50);comment:邮箱" json:"mail"`
	Phone     string `gorm:"varchar(20);comment:手机号" json:"phone"`
	Status    int    `gorm:"default:0;comment:0->正常,1->锁定,2->删除" json:"status"`
	Type      string `gorm:"varchar(20)" json:"type"`
	base
}

func (u *User) TableName() string {
	return "users"
}
