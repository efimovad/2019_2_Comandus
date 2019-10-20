package sqlstore

import (
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (firstName, secondName, username, email, encryptPassword) VALUES ($1, $2, $3, $4, $5) RETURNING accountId",
		u.FirstName,
		u.SecondName,
		u.UserName,
		u.Email,
		u.EncryptPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT accountId, firstName, secondName, username, email, encryptPassword, avatar FROM users WHERE accountId = $1",
		id,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.SecondName,
		&u.UserName,
		&u.Email,
		&u.EncryptPassword,
		&u.Avatar,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT accountId, firstName, secondName, username, email, encryptPassword, avatar FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.SecondName,
		&u.UserName,
		&u.Email,
		&u.EncryptPassword,
		&u.Avatar,
	); err != nil {
		return nil, err
	}
	return u, nil
}

//TODO: validate user
func (r *UserRepository) Edit(u * model.User) error {
	return r.store.db.QueryRow("UPDATE users SET firstName = $1, secondName = $2, userName = $3" +
		"encryptPassword = $4, avatar = $5 WHERE id = $6",
		u.FirstName,
		u.SecondName,
		u.UserName,
		u.EncryptPassword,
		u.Avatar,
		u.ID,
	).Scan(&u.ID)
}