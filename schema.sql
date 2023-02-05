CREATE DATABASE db;
USE db;
CREATE TABLE exercise (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(64),
    description TEXT,
    timeCreated TIMESTAMP,
    timeUpdated TIMESTAMP
);

CREATE TABLE user (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64),
    password VARCHAR(64),
    timeCreated TIMESTAMP,
    timeUpdated TIMESTAMP
)

CREATE TABLE schedule (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    userId INT NOT NULL,
    title VARCHAR(64) NOT NULL,
    description TEXT,
    timeCreated TIMESTAMP,
    timeUpdated TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES user(id)
)

CREATE TABLE active_exercise (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    scheduleId INT NOT NULL,
    originId INT NOT NULL,
    reps INT,
    sets INT,
    timeCreated TIMESTAMP,
    timeUpdated TIMESTAMP,
    FOREIGN KEY (scheduleId) REFERENCES schedule(id),
    FOREIGN KEY (originId) REFERENCES exercise(id)
)