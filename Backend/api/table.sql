






CREATE TABLE User(
    User_id SERIAL PRIMARY KEY,
    Username VARCHAR(100) UNIQUE NOT NULL,
    Userpass VARCHAR(100) NOT NULL,
    Semester INT NOT NULL
);
