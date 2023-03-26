package userDAO

import (
    "log"
    "errors"
    "database/sql"
    "api/pkg/config"
    _"github.com/gin-gonic/gin"
)

var NullDBField *sql.NullString

type User struct {
    Id int `json:id`
    Role int `json:role_id`
    Password string `json:password`
    Email string `json:email`
    Username string `json:username`
    Firstname string `json:firstname`
    Lastname string `json:lastname`
    RegistrationDate string `json:registrationDate`
}

func GetAll() []User {

    result, err := config.DB.Query(
        `SELECT 
        U.id, U.username, R.role_name, U.email, U.first_name, last_name, registration_date
        FROM users as U
        JOIN roles as R
        ON (U.role_id=R.role_id)`);

    if err != nil {
        panic(err)
    }
    
    var resultSlice []User = []User{}

    // the result object has a method called Next,
    // which is used to iterate through all returned rows.
    for result.Next() {
        var u User

        // The result object provided Scan  method
        // to read row data, Scan returns error,
        // if any. Here we read id and name returned.
        err = result.Scan(
            &NullDBField,
            &u.Username,
            &u.Role,
            &NullDBField,
            &u.Email,
            &u.Firstname,
            &u.Lastname,
            &u.RegistrationDate,
        )

        // handle error
        if err != nil {
            log.Println(err)
        }    

        resultSlice = append(resultSlice, u) 

    }
    return resultSlice
}

func GetByUsernameAndPassword(username string, password string) User {

    var foundUser User
    var err error

    result := config.DB.QueryRow(
        `SELECT 
        *
        FROM users
        WHERE username=? and password=?;`, username, password);

    // The result object provided Scan  method
    // to read row data, Scan returns error,
    // if any. Here we read id and name returned.
    err = result.Scan(
        &NullDBField,
        &foundUser.Username,
        &foundUser.Role,
        &NullDBField,
        &foundUser.Email,
        &foundUser.Firstname,
        &foundUser.Lastname,
        &foundUser.RegistrationDate,
    )

    // handle error
    if err != nil {
        log.Println(err)
    }    

    return foundUser
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
