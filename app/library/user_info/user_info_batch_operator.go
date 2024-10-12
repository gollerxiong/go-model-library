package user_info

import (
	"gorm.io/gorm"
	"huabao/app/models"
)

type BatchOperator struct {
	field      string
	fieldValue interface{}
	model      *gorm.DB
	ids        []int
}

func NewBatchOperator() *BatchOperator {
	return &BatchOperator{
		model: GetConn(),
	}
}

func (b *BatchOperator) SetFieldValue(v interface{}) *BatchOperator {	b.fieldValue = v
	return b
}

func (b *BatchOperator) buildParams() {
	if len(b.ids) > 0 {
		b.model = b.model.Where("id in ?", b.ids)
	}
}

func (b *BatchOperator) Update() bool {
	b.buildParams()
	b.model.Model(&models.UserBan{}).Update(b.field, b.fieldValue)
	return true
}

func (b *BatchOperator) Delete() bool {
	b.buildParams()
	b.model.Delete(&models.UserBan{})
	return true
}

