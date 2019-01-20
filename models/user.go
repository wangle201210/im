package models


type User struct {
	Id       int64     `json:"id" orm:"column(id);pk;"`
	Name     string    `json:"name" orm:"column(name);size(100)"`
	Password string    `json:"password" orm:"column(password);size(100)"`
	Role     string    `json:"role" orm:"column(role);size(100)"`
}


var InitData = []User{
	{1,"admin","password","admin"},
	{2,"user1","password","user"},
	{3,"user2","password","user"},
	{4,"user3","password","user"},
	{5,"user4","password","user"},
}

func Find(m User) (r User) {
	for _,v := range InitData {
		if v.Name == m.Name && v.Password == m.Password {
			r = v
		}
	}
	return
}