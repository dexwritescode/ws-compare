# ws-compare
Three, functionally similar web services, written in Java, Go, and Rust.


## Endpoints

| Route                    | Request example  | Response example                  |
|:-------------------------|:-----------------|:----------------------------------|
| GET /hello               | /hello           | {"message":"Hello, World"}        |
| GET /greeting/{name}     | /greeting/Jane   | {"greeting":"Hello, Jane!"}       |
| GET /fibonacci/{number}  | /fibonacci/45    | {"input":45,"output":1134903170}  |


# Build and Run
## Java
### Build
In the java folder:
* Run the command: 
```cmd
mvn package
```  
* The output should be in the ```target``` folder: 
```webservice-1.0-SNAPSHOT-allinone.jar```

### Run
In the java folder issue the command: 
```cmd
java -jar target/webservice-1.0-SNAPSHOT-allinone.jar
```
### Test
* Fr example, using curl:  
```cmd
curl http://localhost:8080/fibonacci/20
```
* The response should be:
```json
{"input":20,"output":6765}
```


## Notes:
- The code is not suitable for production use. And is not idiomatic to each language. Also, lacks tests and documentation.
- This project I wrote to compare some metrics using different programming languages. The metrics I was looking for:
    - Compiled artifact size
    - Size in memory while running idle
    - Size in memory while running at peak
    - Average requests per second




