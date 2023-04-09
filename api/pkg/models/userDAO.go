package userDAO

import (
	"api/pkg/config"
	"database/sql"
	"errors"
	_ "github.com/gin-gonic/gin"
	_ "log"
)

// Used as a null placeholder for db row scans.
var NullDBField *sql.NullString

type User struct {
	Id               int    `json:id`
	Role             string `json:role`
	Password         string `json:password`
	Email            string `json:email`
	Username         string `json:username`
	Firstname        string `json:firstname`
	Lastname         string `json:lastname`
	RegistrationDate string `json:registrationDate`
}

var generalSelectQuery string = `SELECT
                                    U.id, U.role_name, U.password,
                                    U.email, U.username, U.first_name,
                                    U.last_name, U.registration_date
                                FROM users as U
                                JOIN roles as R
                                ON (U.role_id=R.role_id)`

// Returns a slice of User objects scanned from *sql.Rows.
func userSliceFromResult(rows *sql.Rows) ([]User, error) {
	var err error
	var resultSlice []User = []User{}

	for rows.Next() {
		var user User

		err = rows.Scan(
			&user.Id,
			&user.Role,
			&user.Password,
			&user.Email,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.RegistrationDate,
		)

		// handle error
		if err != nil {
			return resultSlice, err
		}

		resultSlice = append(resultSlice, user)
	}

	return resultSlice, nil
}

// Returns a single User object from *sql.Row.
func userFromResult(row *sql.Row) (User, error) {
	var err error
	var user User
	err = row.Scan(
		&user.Id,
		&user.Role,
		&user.Password,
		&user.Email,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.RegistrationDate,
	)

	// handle error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Return a slice with all the users in the database.
func GetAll() ([]User, error) {
	var err error
	var users []User

	result, err := config.DB.Query(generalSelectQuery)

	if err != nil {
		return []User{}, err
	}

	users, err = userSliceFromResult(result)

	if err != nil {
		return users, err
	}

	return users, err

}

// Get the user corresponding to the given ID.
func GetByID(ID int) (User, error) {
	var err error
	var foundUser User

	result := config.DB.QueryRow(generalSelectQuery+`WHERE U.id=?;`, ID)

	foundUser, err = userFromResult(result)

	// handle error
	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return foundUser, nil
}

// Get the user corresponding to the given username.
func GetByUsername(username string) (User, error) {
	var err error
	var foundUser User

	result := config.DB.QueryRow(generalSelectQuery+`WHERE U.username=?;`, username)

	foundUser, err = userFromResult(result)

	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	return foundUser, nil
}

func InsertNewUser(newUser User) error {
	var err error
	var res sql.Result
	var affectedRowNum int64

	res, err = config.DB.Exec(
		`INSERT INTO users (username,password,email,first_name,last_name) 
         VALUES (?,?,?,?,?)`,
		newUser.Username, newUser.Password, newUser.Email, newUser.Firstname, newUser.Lastname,
	)

	if err != nil {
		return err
	}

	affectedRowNum, err = res.RowsAffected()

	if err != nil {
		return err
	} else if affectedRowNum < 1 {
		return errors.New("No rows have been affected!")
	}

	return nil
}
