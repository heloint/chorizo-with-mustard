DROP DATABASE IF EXISTS testdb;

CREATE DATABASE IF NOT EXISTS testdb;

USE testdb;
-- #####################################################################

-- Delete previous user and create a new one.
-- #####################################################################

DROP USER IF EXISTS testusr;
CREATE USER IF NOT EXISTS 'testusr'@'localhost' IDENTIFIED BY 'testpass';
GRANT ALL ON testdb.* TO testusr@localhost;

-- #####################################################################


-- Create sequences
-- #####################################################################
CREATE SEQUENCE user_id START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE roles_id START WITH 1 INCREMENT BY 1;
-- #####################################################################

-- Create and fetch roles table.
-- #####################################################################
CREATE TABLE IF NOT EXISTS roles (
                                role_id INT(10) PRIMARY KEY,
                                role_name VARCHAR(25) NOT NULL
);

INSERT INTO roles VALUES
    (NEXT VALUE FOR roles_id, "admin"),
    (NEXT VALUE FOR roles_id, "user")
;
-- #####################################################################

-- Create and fetch users table.
-- #####################################################################
CREATE TABLE IF NOT EXISTS users (
                                id INT(10) PRIMARY KEY DEFAULT NEXT VALUE FOR user_id,
                                username VARCHAR(25) NOT NULL CHECK(username <> ''),
                                role_id INT(10) NOT NULL  DEFAULT 2 CHECK(role_id <> 0),
                                password VARCHAR(150) NOT NULL CHECK(password <> ''),
                                email VARCHAR(30) NOT NULL CHECK(email <> ''),
                                first_name VARCHAR(50) NOT NULL CHECK(first_name <> ''),
                                last_name VARCHAR(60) NOT NULL CHECK(last_name <> ''),
                                registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                FOREIGN KEY (role_id) REFERENCES roles(role_id),
                                CONSTRAINT user_name_unique UNIQUE (username),
                                CONSTRAINT user_email_unique UNIQUE (email)
);

INSERT INTO users (username, role_id, password, email, first_name, last_name) VALUES
    ("admin", 1, "admin", "admin@gmail.com", "Fafa", "Nana"),
    ("user01", 2, "pass01", "lili@gmail.com", "Lili", "Lala"),
    ("user02", 2, "pass02", "didi@gmail.com", "Didi", "Dada")
;
