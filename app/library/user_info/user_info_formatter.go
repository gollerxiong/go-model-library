package user_info

type ModelFormatter struct {
	 data map[string]interface{}
}

func (u *ModelFormatter) SetData(data map[string]interface{}) {
	u.data = data
}
func (u *ModelFormatter) Formate() map[string]intermace{} {
	var result = make(map[string]interface{})
	for key, val := range u.data {
		tmp := columnFormate(key, val)
		result[key] = tmp
	}

	u.data = result
	return u.data
}


