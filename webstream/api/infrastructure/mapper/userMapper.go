package mapper

import (
	"log"
	"webstream/api/entity"
)

func CreateUser(user *entity.User) (*entity.User, error) {
	stmt, err := Db.Prepare("insert into sys_user(username,password) values (?,?)")
	if err != nil {
		return nil, err
	}
	exec, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return nil, err
	}
	userId, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	user.Id = userId
	return user, nil
}
func UpdateUser(user *entity.User) (bool, error) {
	//var result int64
	stmt, err := Db.Prepare("update sys_user set username = ?,password=? where id = ?")
	if err != nil {
		return false, err
	}
	result, err := stmt.Exec(user.Username, user.Password, user.Id)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	defer stmt.Close()
	return affected != 0, nil
}

func DeleteUser() {

}

func GetUserCredential(username string, password string) (bool, error) {
	stmt, err := Db.Prepare("select count(1) from sys_user where username = ? and password = ?")
	if err != nil {
		log.Println(err)
		return false, err
	}
	var result int64
	query := stmt.QueryRow(username, password)
	err = query.Scan(&result)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer stmt.Close()
	return result == 1, nil
}
func FindUserByUserName(username string) (*entity.User, error) {
	var user entity.User
	stmt, err := Db.Prepare("select * from sys_user where username = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Password,
		&user.CreateTime, &user.UpdateTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, err
}
func FindUserNameIsExist(username string) (bool, error) {
	var result int
	stmt, err := Db.Prepare("select count(1) from sys_user where username = ?")
	if err != nil {
		log.Println(err)
		return false, err
	}
	err = stmt.QueryRow(username).Scan(&result)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return result == 1, err
}
