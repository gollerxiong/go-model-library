package user_info

import (
	"gorm.io/gorm"
	"huabao/app/models"
	"huabao/utils"
	"strings"
)



type UserInfo struct {
	isNew	bool
	model	*models.UserInfo
	field	string
	formatter	*UserInfoFormatter
	attributes	map[string]interface{}
	oldAttributes	map[string]interface{}
}


func (o *UserInfo) SetNew(n bool) *UserInfo {
	o.isNew = n
	return o
}


func (o *UserInfo) IsNew() bool {
	return o.isNew
}


func (o *UserInfo) GetModel() *models.UserInfo {
	return o.model
}


func (o *UserInfo) SetField(f string) *UserInfo {
	o.field = f
	return o
}


func (o *UserInfo) addField(f string) {
	if !strings.Contains(o.field, "*") {
		o.field += "," + f
	}
}


func (o *UserInfo) SetAttributes(m map[string]interface{}) *UserInfo {
	o.attributes = m
	return o
}


func (o *UserInfo) SetOldAttributes(m map[string]interface{}) *UserInfo {
	o.oldAttributes = m
	return o
}


func (o *UserInfo) GetAttributes() map[string]interface{} {
	result := make(map[string]interface{})
	if !strings.Contains(o.field, "*") {
		fieldArr := strings.Split(o.field, ",")
		for _, key := range fieldArr {
			val, ok := o.attributes[key]
			if ok {
				result[key] = val
			}
		}
	} else {
		result = o.attributes
	}
	o.formatter.SetData(result)
	result = o.formatter.Formate()
	return result
}


func (o *UserInfo) Save() *UserInfo {
	if o.IsNew() {
		conn.Save(o.model)
	} else {
		diffMap := utils.DiffMapBaseFirst(o.oldAttributes, o.attributes)
		keys := utils.MapKeys(diffMap)
		conn.Model(&models.UserInfo{}).Where("id = ?", o.model.Id).Select(keys).Updates(diffMap)
	}
	o.SetNew(false)
	return o
}


func (o *UserInfo) Delete() bool {
	if o.IsNew() || o.model.ID < 1 {
		return true
	} else {
		res := conn.Delete(o.model).Error
		if res != nil {
			return false
		}
	}
	return true
	}
}


func (o *UserInfo) LoadById(id int) *UserInfo {
	connection := GetConn().Model(o.model)
	err := connection.Where("id=?", id).First(o.model).Error
	if err == gorm.ErrRecordNotFound {
		o.SetNew(true)
	} else {
		o.SetNew(false)
		attributes, _ := utils.StructToMap(*o.model)
		o.SetAttributes(attributes)
		o.SetOldAttributes(attributes)
	}
	return o
}