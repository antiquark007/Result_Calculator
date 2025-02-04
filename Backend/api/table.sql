CREATE TABLE User(
    User_id   SERIAL INT PRIMARY KEY ,
    Username VARCHAR(100) NOT NULL,
    Userpass VARCHAR(100) NOT NULL,
    Semester INT NOT NULL
);

INSERT INTO User (Username, Userpass, Semester)
VALUES ('user1', 'pass1', 1);