package globalkey

const (
	SysPermMenuPrefix          = "/admin/"
	SysJwtUserId               = "userId"
	SysPermMenuCachePrefix     = "cache:arkAdmin:permMenu:"
	SysOnlineUserCachePrefix   = "cache:arkAdmin:online:"
	SysLoginCaptchaCachePrefix = "cache:arkAdmin:captcha:"
	SysUserIdCachePrefix       = "cache:arkAdmin:sysUser:id:"
	SysDateFormat              = "2006-01-02 15:04:05"
	SysNewUserDefaultPassword  = "123456"
	SysShowSystemError         = true
	SysProtectPermMenuMaxId    = 44
	SysProtectDictionaryMaxId  = 4
	SysSuperUserId             = 1
	SysSuperRoleId             = 1
	SysDefaultPermType         = 2
	SysEnable                  = 1
	SysDisable                 = 0
	SysLoginLogType            = 1
	SysTopParentId             = 0

	// 课本对应的单词有序结合key前缀 记录课本对应的单词id 拼接课本id
	BookWordPrefix = "bookWord:"
	// 用户对应课本单词 有序集合key前缀 记录用户学习的课本下所有单词id,单词排序 拼接用户id 课本id
	UserBookSortSetPrefix = "userBookSortSet:"
	// 用户对应课本单词学习记录key前缀 记录用户学习的单词次数,下次复习时间,单词id  拼接用户id  课本id
	UserStudyBookPrefix = "userStudyBook:"

	// 用户课本已掌握的单词数量 记录用户对应课本已经掌握的单词数量 凭借用户id, 课本id
	UserBookWordCountPrefix = "userBookWordCount:"

	// 用户学习对应的课本id 拼接用户id=课本id
	UserBookPrefix = "userBook:"
)
