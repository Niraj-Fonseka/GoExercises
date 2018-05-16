# valet-crud-api_go
Go Implementation of the valet crud api


Database description 

Tables - 

collection - 
```
+---------------+--------------+------+-----+---------+----------------+
| Field         | Type         | Null | Key | Default | Extra          |
+---------------+--------------+------+-----+---------+----------------+
| id            | int(11)      | NO   | PRI | NULL    | auto_increment |
| name          | varchar(100) | YES  |     | NULL    |                |
| timeseries_id | int(11)      | YES  | MUL | NULL    |                |
| service_ref   | varchar(100) | YES  | MUL | NULL    |                |
+---------------+--------------+------+-----+---------+----------------+     
```           
measurement - 
```
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int(11)      | NO   | PRI | NULL    | auto_increment |
| ref   | varchar(100) | YES  | UNI | NULL    |                |
+-------+--------------+------+-----+---------+----------------+
```

metric - 
```
+-----------------+--------------+------+-----+---------+----------------+
| Field           | Type         | Null | Key | Default | Extra          |
+-----------------+--------------+------+-----+---------+----------------+
| id              | int(11)      | NO   | PRI | NULL    | auto_increment |
| value           | double       | YES  |     | NULL    |                |
| service_ref     | varchar(100) | YES  | MUL | NULL    |                |
| measurement_ref | varchar(100) | YES  | MUL | NULL    |                |
| collection_id   | int(11)      | YES  | MUL | NULL    |                |
+-----------------+--------------+------+-----+---------+----------------+
```

service  - 
```
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int(11)      | NO   | PRI | NULL    | auto_increment |
| ref   | varchar(100) | YES  | UNI | NULL    |                |
+-------+--------------+------+-----+---------+----------------+
```

timeserise - 
```
+-----------------+-------------+------+-----+---------+----------------+
| Field           | Type        | Null | Key | Default | Extra          |
+-----------------+-------------+------+-----+---------+----------------+
| id              | int(11)     | NO   | PRI | NULL    | auto_increment |
| interval_type   | varchar(15) | YES  |     | NULL    |                |
| date_start      | date        | YES  |     | NULL    |                |
| date_end        | date        | YES  |     | NULL    |                |
| datetime_start  | datetime    | YES  |     | NULL    |                |
| datetime_end    | datetime    | YES  |     | NULL    |                |
| datetime_insert | datetime    | YES  |     | NULL    |                |
+-----------------+-------------+------+-----+---------+----------------+
```

user - 
```
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(11)      | NO   | PRI | NULL    | auto_increment |
| api_key    | varchar(100) | YES  | UNI | NULL    |                |
| api_secret | varchar(100) | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```

#DB

export dbconn="root:password@/dbname?charset=utf8&parseTime=True&loc=Local"
