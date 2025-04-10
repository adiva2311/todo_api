CREATE DATABASE todo_list;

USE todo_list;

CREATE TABLE user
(
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
) ENGINE = InnoDB;

CREATE TABLE list (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    information VARCHAR(100),
    completed BOOLEAN DEFAULT FALSE,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES user(id)
)ENGINE = InnoDB;

SELECT * FROM user;

SELECT * FROM list;