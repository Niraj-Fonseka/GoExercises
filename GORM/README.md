SQL - Structured Query Language ( :B )
ORM - Object Relational Mapper
DDL - Database Definition Language 

-- Introduction 

* CRUD Operations
* Schema Management
* Transactions 
    - Series of database transactions happen ( atomically ( all or nothing ))
* Migrations 
    - Add and delete fields from a database 
* Event Hooks
    - Functions that get called before and after
* Chainable API 
* Logger 
    - Inspect what happens while the database works with the database


-- Supported Databases

* PostgreSQL
* MySQL
* SQLite
* Foundation
    - Similar to Postgre

-- Documentation 

http://gorm.io/docs/


-- Go Get 

- Gorm
go get -u github.com/jinzhu/gorm

- PostgreSQL Driver
go get -u github.com/lib/pq

