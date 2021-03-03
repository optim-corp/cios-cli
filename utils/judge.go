package utils

func Is(flag bool) Judge {
	return Judge{flag: flag}
}
func (judge Judge) True(val interface{}) Judge {
	if judge.flag == true {
		judge.Value = val
	}
	return judge
}
func (judge Judge) False(val interface{}) Judge {
	if judge.flag == false {
		judge.Value = val
	}
	return judge
}

func (judge Judge) T(val interface{}) Judge {
	if judge.flag == true {
		judge.Value = val
	}
	return judge
}

func (judge Judge) F(val interface{}) Judge {
	if judge.flag == false {
		judge.Value = val
	}
	return judge
}
