package user_info

import (
	"huabao/app/models"
	"huabao/utils"
	"strings"
	"sync"
)

func newObjectListFormatter() (of *objectListFormatter) {
	of = &objectListFormatter{}
	return
}

type objectListFormatter struct {
	list	[]map[string]interface{}
	fields	string
	wg	sync.WaitGroup
	result	[]map[string]interface{}
}

func (u *objectListFormatter) setList(list []models.UserInfo) {
	var end = len(list)
	var tmp = make([]map[string]interface{}, end)

	u.result = tmp

	for i := 0; i < end; i++ {
		tmpmap, _ := utils.StructToMap(list[i])
		tmp[i] = tmpmap
	}
	u.list = tmp
}

func (u *objectListFormatter) setFields(f string) {
	u.fields = f
}

func (u *objectListFormatter) formate() []map[string]interface{} {
	var length = len(u.list)

	u.wg.Add(length)
	for i := 0; i < length; i++ {
		go u.do(i, u.list[i])
	}

	u.wg.Wait()
	return u.result
}

func (u *objectListFormatter) do(index int, item map[string]interface{}) {
	defer u.wg.Done()
	result := make(map[string]interface{})

	if !strings.Contains(u.fields, "*") {
		fieldArr := strings.Split(u.fields, ",")
		for _, key := range fieldArr {
			val, ok := item[key]
			if ok {
				result[key] = columnFormate(key, val)
			}
		}
		u.result[index] = result
	} else {
		u.result[index] = item
	}
}


