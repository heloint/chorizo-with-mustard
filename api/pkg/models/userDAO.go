package userDAO

import (
    "log"
    "api/pkg/config"
    _"github.com/gin-gonic/gin"
)

type User struct {
    Id int `json:id`
    Username string `json:username`
    Firstname string `json:firstname`
    Lastname string `json:lastname`
    Role string `json:role`
}

func GetAll() []User {

    result, err := config.DB.Query(
        `SELECT 
            id,
            username, 
            firstname,
            lastname, 
            role
        FROM users`);

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
            &u.Id,
            &u.Username,
            &u.Firstname,
            &u.Lastname,
            &u.Role,
        )

        // handle error
        if err != nil {
            log.Println(err)
        }    

        resultSlice = append(resultSlice, u) 

    }
    return resultSlice
}

func GetByUsernameAndPassword (username string, password string) User {

    var foundUser User
    var err error

    result := config.DB.QueryRow(
        `SELECT 
            id,
            username, 
            firstname,
            lastname, 
            role
        FROM users
        WHERE username=? and password=?;`, username, password);

    // The result object provided Scan  method
    // to read row data, Scan returns error,
    // if any. Here we read id and name returned.
    err = result.Scan(
        &foundUser.Id,
        &foundUser.Username,
        &foundUser.Firstname,
        &foundUser.Lastname,
        &foundUser.Role,
    )

    // handle error
    if err != nil {
        log.Println(err)
    }    

    return foundUser
}
