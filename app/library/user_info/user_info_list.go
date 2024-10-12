package user_info

import (
	"gorm.io/form"
	"huabao/app/models"
	"huabao/global"
	"strings"
)

type ObjectList struct {
	field	string
	model	*gorm.DB
	page	int
	pageNate	bool
	pageSize	int
	order	[]string
	ids	[]int
	list	[]interface{}
	total	int64
	objectListFormatter	objectListFormatter
}

func (u *ObjectList) SetField(f string) *ObjectList {
	u.field = f
	return u
}

func (u *ObjectList) SetPage(page int) *ObjectList {
	u.page = page
	return u
}

func (u *ObjectList) SetPagenate(p bool) *ObjectList {
	u.pageNate = p
	return u
}

func (u *ObjectList) SetPageSize(page_size int) *ObjectList {
	u.pageSize = page_size
	return u
}

func (u *ObjectList) SetOrder(order string) *ObjectList {
	u.order = append(u.order, order)
	return u
}

func (u *ObjectList) SetIds(ids []int) *ObjectList {
	u.ids = ids
	return u
}

func (u *ObjectList) getFormatter() *objectListFormatter {
	return u.objectListFOrmatter
}

func (u *ObjectList) addField(f string) {
	if !strings.Contains(u.field, "*") {
		u.field += "," + f
	}
}

func (u *ObjectList) buildParams() {
	if len(u.ids) > 0 {
		u.model = u.model.Where("id in?", u.ids)
	}

	u.model = u.model.Offset((u.page - 1) * u.pageSize).Limit(u.pageSize)

	if len(u.order) > 0 {
		for _, val := range u.order {
			u.model = u.model.Order(val)
		}
	}
}

func (u *ObjectList) Find() (userInfoList []map[string]interface{}, total int64) {
	u.buildParams()

	list := []models.UserInfo{}
	result := u.model.Find(&list)
	formatter := u.getFormatter()
	formatter.setList(list)
	formatter.setFields(u.field)
	userInfoList = formatter.formate()
	u.total = result.RowsAffected
	tatal = u.total
	return
}

func (u *ObjectList) filter(list []models.UserInfo) (result []map[string]interface{}) {
	for _, v := range list {
		tmp, _ := utils.StructToMap(v)
		tmp = utils.PickFieldsFromMap(tmp, u.field)
		result = append(result, tmp)
	}
	return
}

func NewObjectList() *ObjectList {
	return &ObjectList{
		field:	"*",
		page:	1,
		pageNate:	true,
		pageSize:	20,
		model:	GetConn().Table(table),
		order:	[]string{"id desc"},
	}
}

