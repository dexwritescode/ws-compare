# ws-compare
Three, functionally similar web services, written in Java, Go, and Rust.

  The main motivation for writing these three simple web services is to compare the three languages, in this case Java, Go, and Rust.

## Endpoints

| Route                    | Request example  | Response example                  |
|:-------------------------|:-----------------|:----------------------------------|
| GET /hello               | /hello           | {"message":"Hello, World"}        |
| GET /greeting/{name}     | /greeting/Jane   | {"greeting":"Hello, Jane!"}       |
| GET /fibonacci/{number}  | /fibonacci/45    | {"input":45,"output":1134903170}  |


# Build and Run
  For each version of the program, issue the commands from the specific language folder.

## Java

### Dependencies
* [Akka](https://akka.io/)
 
### Build
    mvn package

  The compiled program should now be in the `target` folder: 
    
    webservice-1.0-allinone.jar

### Run
    java -jar target/webservice-1.0-allinone.jar

## Go

### Dependencies
* [Echo](https://echo.labstack.com/)
 
### Build
    go build -o webservice

  The compiled program should now be in the `ws-compare/go` folder: 
    
    webservice

### Run
    ./webservice

## Rust
### Dependencies
### Build
### Run

## Testing it
  For example, using curl:  

    curl http://localhost:8080/fibonacci/20

  The response should be:

    {"input":20,"output":75025}

# Benchmark
## Using [wrk](https://github.com/wg/wrk)
  Default timeout (2 seconds)

    wrk -t2 -c400 -d30s http://127.0.0.1:8080/fibonacci/25


  Extended timeout (5 seconds)

    wrk -t2 -c400 -d30s http://127.0.0.1:8080/fibonacci/25 --timeout 5

## Notes:
- The code is not suitable for production use. And is not idiomatic to each language. Also, lacks tests and documentation.
- This project I wrote to compare some metrics using different programming languages. The metrics I was looking for:
    - Compiled artifact size
    - Size in memory while running idle
    - Size in memory while running at peak requests per second
    - Requests per second, latency, and timeouts
    - Packaged, deployment ready build with the relevant runtime. (Docker image)
