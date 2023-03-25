package user

import (
    "api/pkg/config"
)

type User struct {
    id int
    username string
    password string
    firstname string
    lastname string
    role string
}

func GetAllUsers() []User {

    result, err := config.DB.Query("Select * from users;");

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
            &u.id,
            &u.username,
            &u.password,
            &u.firstname,
            &u.lastname,
            &u.role,
        )

        // handle error
        if err != nil {
            panic(err)
        }    

        resultSlice = append(resultSlice, u) 

    }
    return resultSlice
}
