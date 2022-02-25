# Booksy WIP

A simple REST API that creates, updates, lists, gets and deletes a book. 

# Tools

- Golang
- Ginkgo for testing
- mysql for database
- gorilla/mux to handle requests


# Installation

```
git clone <this or your fork>
cd booksy
go build
./booksy
```

# Usage

Database `books` should exist locally.
A table called books should exist locally as well
```sql
CREATE TABLE books (
    id INT unsigned NOT NULL AUTO_INCREMENT, 
    title VARCHAR(150) NOT NULL, 
    author VARCHAR(150) NOT NULL, 
    content VARCHAR(150) NOT NULL, 
    PRIMARY KEY     (id)  
    );
```

To **create** a book 

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Northern lights", "author": "Philip Pullman", "content": "Awesome"}' http://localhost:10000/book
```

To **update** a book

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"title": "Northern lights", "author": ""Philip Pullman", "content": "double awesome"}' http://localhost:10000/book/<id>
```

To **list** books

```bash
curl http://localhost:10000/books
```

To **get** a single book

```bash
curl http://localhost:10000/book/<id>
```

To **delete** a book

```bash
curl -X "DELETE" http://localhost:10000/book/<id>
```

# Test 

Run `ginkgo -r` 

# Improvements

- Add test for sql

-  Refactor Specs

- Add test_database

- Create a database

- Add config for database info

- Refactor Handler