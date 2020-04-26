# ws-compare
Three, functionally similar web services, written in Java, Go, and Rust.

  The main motivation for writing these three simple web services, is to compare the three languages, Java, Go, and Rust.

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

  The compiled program (executable jar) should now be in the `ws-compare/java/target` folder: 
    
    webservice.jar

### Run
    java -jar target/webservice.jar

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
* [Actix](https://github.com/actix/actix)
* [Actix web](https://github.com/actix/actix-web)
* [Serde](https://serde.rs/)
### Build
    cargo build --release

  The compiled program should be in `ws-compare/rust/target/release` folder:
    webservice
### Run
    ./webservice

# Testing it
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

## Using [vegeta](https://github.com/tsenart/vegeta)
  Run for 30 seconds, 2 workers, no rate limiting

    echo "GET http://localhost:8080/fibonacci/25" | vegeta attack -duration=30s -rate=0 -max-workers=2 | tee results.bin | vegeta report

  Generate graph
  
    vegeta plot > plot.html

## Docker
### Build
  Example: while in the java directory, run the command:

    docker build -t java/webservice:1.0 .

### Run 

    docker run -p 8080:8080 java/webservice:1.0

# Notes:
- The code is not suitable for production use. And is not idiomatic to each language. Also, lacks tests and documentation.
- This project was written to compare some metrics using different programming languages. The metrics I was looking for:
    - Compiled artifact size
    - Size in memory while running idle
    - Size in memory while running at peak requests per second
    - Requests per second, latency, and timeouts
    - Packaged, deployment ready build with the relevant runtime. (Docker image)
    - Compile/build time
