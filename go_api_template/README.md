export DBCON="root:password@/gormtest?charset=utf8&parseTime=True&loc=Local"



### Login post body
```
{
	"username": "admin",
	"password": "admin"

}
```


### Setting the timeout for the token in the auth.go file 
```
Timeout:       20 * time.Second
```