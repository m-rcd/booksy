package database

const CreateDatabase = `CREATE DATABASE IF NOT EXISTS bookshop;`

const CreateBookTable = `
CREATE TABLE if not exists books (
    id INT unsigned NOT NULL AUTO_INCREMENT, 
    title VARCHAR(150) NOT NULL, 
    author VARCHAR(150) NOT NULL, 
    content VARCHAR(150) NOT NULL, 
    PRIMARY KEY     (id)  
    );`
