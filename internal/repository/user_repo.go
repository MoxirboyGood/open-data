package repository

import (
	"errors"
	"testDeployment/internal/delivery/dto"
	"testDeployment/internal/domain"
	"time"
)

func (r repo) Register(user domain.User) (id int, err error) {
	query := `
	insert into users (email,password,role,created_at,updated_at,deleted_at,is_email_verified) values($1,$2,$3,$4,$5,$6,$7) returning id
`
	row := r.db.QueryRow(query, user.Phone_number(), user.Password(), user.Role(), user.Created_at(), user.Updated_at(), user.Deleted_at(),true)
	if err := row.Scan(&id); err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (r repo) Exist(email string) (exist bool, err error) {

	query := `
	
		SELECT is_active
		FROM users
		WHERE email = $1 AND is_active = true
		
`
	err = r.db.QueryRow(query, email).Scan(&exist)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return false, domain.ErrCouldNotScan
	}
	
	
	return exist, nil
}
func (r repo) GetByEmail(email string) (id int, password string, err error) {
	query := `
		select id ,password from users where email=$1
`
	err = r.db.QueryRow(query, email).Scan(&id, &password)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, "", err
	}
	return id, password, nil
}
func (r repo) GetAll() (User []dto.User) {
	var user userDB
	query := `
		select * from users
`
	rows, err := r.db.Query(query)
	if err != nil {
		r.Bot.SendErrorNotification(err)
	}
	for rows.Next() {

		err := rows.Scan(&user.id, &user.phone_number, &user.role, &user.created_at, &user.updated_at, &user.deleted_at)
		if err != nil {
			r.Bot.SendErrorNotification(err)
		}

		User = append(User, r.f.ParseModelToDomain(user.id, user.phone_number, user.role, user.created_at, user.updated_at, user.deleted_at))
	}
	return User
}
func (r repo) UpdatePhoneNumber(number string) (id int, err error) {
	return 0, err
}
func (r repo) UpdateIsActive(id int, deleteAt time.Time) (err error) {
	query := `
	update users set is_active=false ,deleted_at=$1 where id=$2
	`
	_, err = r.db.Exec(query, deleteAt, id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return errors.New("could not delete")
	}
	return nil

}

func (r repo) UpdateVerified(userId interface{}) (err error) {
	query := `
		Update users set is_email_verified=$1 where id=$2 
`

	_, err = r.db.Exec(query, true, userId)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return err
	}
	return nil
}
