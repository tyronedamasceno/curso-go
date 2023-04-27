package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick, email from users where name like ? or nick like ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) RetrieveUserByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email from users where id = ?", ID,
	)
	if err != nil {
		return models.User{}, nil
	}

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (repository users) RetrieveUserByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, password from users where email = ?", email,
	)
	if err != nil {
		return models.User{}, nil
	}

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository users) Follow(followerId, followedId uint64) error {
	statement, err := repository.db.Prepare(
		"insert into followers (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followedId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) Unfollow(followerId, followedId uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followedId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchFollowers(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email
		from users  	u
		join followers as f on u.id=f.follower_id
		where f.user_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}

func (repository users) SearchFollowing(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email
		from users  	u
		join followers as f on u.id=f.user_id
		where f.follower_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}
