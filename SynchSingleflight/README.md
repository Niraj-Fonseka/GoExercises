# sync/singleflight

single flight is a package that lets you handle multiple http requests with limited calls to the resources


### To run the locust
```
locust --host=http://localhost:8080
```

End Goal : 

Workflow : 
    request comes in : check the db for the last timestamp 
    
    if the requested timestamp is younger than the timestamp that exist in the databse, 
    -> go to the external datasource to fetch data , write to the database and then return the response

    else 
    -> gran from the database and then return the response