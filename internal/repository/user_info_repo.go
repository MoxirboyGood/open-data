package repository

import (
	"errors"
	"testDeployment/internal/domain"
	"time"
)

func (r repo) UpdateUserInfoDeleted(id int, deleteAt time.Time) (err error) {
	query := `
	update user_info set deleted_at=$1 where user_id=$2
	`
	_, err = r.db.Exec(query, deleteAt, id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return errors.New("could not delete")
	}
	return nil
}

func (r repo) ExistUserInfo(userId int) (exist bool, err error) {
	query := `
	Select Exists (
		SELECT true
		FROM user_info
		WHERE user_id = $1)
	`
	err = r.db.QueryRow(query, userId).Scan(&exist)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return false, err
	}
	return exist, nil
}
func (r repo) CreateInfo(user domain.UserInfo) (id int, err error) {
	query := `
	insert into  user_info (user_id,name,weigh,height,age,waist,created_at,gender) values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id
`
	row := r.db.QueryRow(query, user.Id, user.Name, user.Weigh, user.Height, user.Age, user.Waist, user.UpdatedAt, user.Gender)
	if err = row.Scan(&id); err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (r repo) GetUserInfo(userId int) (user domain.UserInfo, err error) {
	query := `
	select id, name,weigh,height,age,waist,gender from user_info where user_id=$1
	`
	err = r.db.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Weigh,
		&user.Height,
		&user.Age,
		&user.Waist,
		&user.Gender,
	)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return user, domain.ErrCouldNotScan
	}
	return user, nil
}

func (r repo) UpdateInfo(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set name=$2,weigh=$3,height=$4,age=$5,waist=$6,updated_at=$7,gender=$8 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Name, user.Weigh, user.Height, user.Age, user.Waist, user.UpdatedAt, user.Gender).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) UpdateName(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set name=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Name, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) UpdateWeigh(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set weigh=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Weigh, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) UpdateHeight(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set height=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Height, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil

}
func (r repo) UpdateAge(user domain.UserInfo) (id int, err error) {

	query := `
	update user_info set age=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Age, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) UpdateWaist(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set waist=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Waist, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}

func (r repo) UpdateGender(user domain.UserInfo) (id int, err error) {
	query := `
	update user_info set gender=$2,updated_at=$3 where user_id=$1 returning id
	`
	err = r.db.QueryRow(query, user.Id, user.Gender, user.UpdatedAt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
