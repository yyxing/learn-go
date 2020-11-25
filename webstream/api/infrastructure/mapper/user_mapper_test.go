package mapper

import (
	"testing"
	"time"
	"webstream/api/entity"
)

func clearTable() {
	Db.Exec("truncate sys_user")
}
func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("insert", testCreateUser)
	t.Run("update", testUpdateUser)
}

func testUpdateUser(t *testing.T) {
	user := entity.User{
		Id:       1,
		Username: "admin123",
		Password: "admin123",
	}
	result, err := UpdateUser(&user)
	if err != nil {
		panic(err)
	}
	credential, err := GetUserCredential("admin123", "admin123")
	if !credential || !result {
		t.Errorf("update user error fail: %s", err)
	}
}

func testCreateUser(t *testing.T) {
	user := entity.User{
		Username:   "admin",
		Password:   "admin",
		CreateTime: time.Now(),
	}
	result, err := CreateUser(&user)
	if err != nil {
		panic(err)
	}
	credential, err := GetUserCredential("admin", "admin")
	if !credential || result == nil {
		t.Errorf("update user error fail: %s", err)
	}
}
