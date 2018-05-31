
install mysql and create local database called = gormtest

username : root
password : password

.env >> export dbconn="root:password@/gormtest?charset=utf8&parseTime=True&loc=Local"


### Addresses 
`
+---------+--------------+------+-----+---------+-------+
| Field   | Type         | Null | Key | Default | Extra |
+---------+--------------+------+-----+---------+-------+
| street  | varchar(255) | YES  |     | NULL    |       |
| number  | varchar(255) | YES  |     | NULL    |       |
| city    | varchar(255) | YES  |     | NULL    |       |
| api_key | varchar(255) | YES  | MUL | NULL    |       |
+---------+--------------+------+-----+---------+-------+
`

### User 
`
+------------+--------------+------+-----+---------+-------+
| Field      | Type         | Null | Key | Default | Extra |
+------------+--------------+------+-----+---------+-------+
| api_key    | varchar(100) | YES  | UNI | NULL    |       |
| api_secret | varchar(100) | YES  |     | NULL    |       |
+------------+--------------+------+-----+---------+-------+
`

### POST for Address 
`
{
	"street": "Street Name",
	"number": "01",
	"city" :  "Austin",
	"api_key" : "username1"
}
`

### POST for User 
`
{
	"api_key": "username6",
	"api_secret": "secret1"
}
`