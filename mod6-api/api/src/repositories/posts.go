package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

// func (repository users) Search(nameOrNick string) ([]models.User, error) {
// 	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

// 	lines, err := repository.db.Query(
// 		"select id, name, nick, email from users where name like ? or nick like ?",
// 		nameOrNick, nameOrNick,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer lines.Close()

// 	var users []models.User

// 	for lines.Next() {
// 		var user models.User

// 		if err = lines.Scan(
// 			&user.ID,
// 			&user.Name,
// 			&user.Nick,
// 			&user.Email,
// 		); err != nil {
// 			return nil, err
// 		}

// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// func (repository users) RetrieveUserByID(ID uint64) (models.User, error) {
// 	lines, err := repository.db.Query(
// 		"select id, name, nick, email from users where id = ?", ID,
// 	)
// 	if err != nil {
// 		return models.User{}, nil
// 	}

// 	var user models.User
// 	if lines.Next() {
// 		if err = lines.Scan(
// 			&user.ID,
// 			&user.Name,
// 			&user.Nick,
// 			&user.Email,
// 		); err != nil {
// 			return models.User{}, err
// 		}
// 	} else {
// 		return models.User{}, errors.New("user not found")
// 	}

// 	return user, nil
// }

// func (repository users) RetrieveUserByEmail(email string) (models.User, error) {
// 	lines, err := repository.db.Query(
// 		"select id, password from users where email = ?", email,
// 	)
// 	if err != nil {
// 		return models.User{}, nil
// 	}

// 	var user models.User
// 	if lines.Next() {
// 		if err = lines.Scan(
// 			&user.ID,
// 			&user.Password,
// 		); err != nil {
// 			return models.User{}, err
// 		}
// 	}

// 	return user, nil
// }

// func (repository users) RetrieveUserPasswordById(userId uint64) (string, error) {
// 	lines, err := repository.db.Query(
// 		"select id, password from users where id = ?", userId,
// 	)
// 	if err != nil {
// 		return "", nil
// 	}

// 	var user models.User
// 	if lines.Next() {
// 		if err = lines.Scan(
// 			&user.ID,
// 			&user.Password,
// 		); err != nil {
// 			return "", err
// 		}
// 	}

// 	return user.Password, nil
// }

// func (repository users) Update(ID uint64, user models.User) error {
// 	statement, err := repository.db.Prepare(
// 		"update users set name = ?, nick = ?, email = ? where id = ?",
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()

// 	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repository users) Delete(ID uint64) error {
// 	statement, err := repository.db.Prepare("delete from users where id = ?")
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()

// 	if _, err := statement.Exec(ID); err != nil {
// 		return err
// 	}

// 	return nil
// }
