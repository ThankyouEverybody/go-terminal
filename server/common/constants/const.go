package constants

const (
	DB             string = "db"
	NormalUser            = "normal"     // 普通用户
	AdminUser             = "admin"      // 管理员
	SuperAdminUser        = "superAdmin" // 超级管理员
	N0             int    = 0
	N1             int    = 1
	N2             int    = 2
	N3             int    = 3
)

// GetStatus 获取状态值
func GetStatus(n int) string {
	switch n {
	case N0:
		return "正常"
	case N1:
		return "锁定"
	case N2:
		return "删除"
	default:
		return "未知"
	}

}
