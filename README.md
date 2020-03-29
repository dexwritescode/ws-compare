# ws-compare
Three, functionally similar web services, written in Java, Go, and Rust.


## Routes

| Route                    | Request          | Response                          |
|:-------------------------|:-----------------|:----------------------------------|
| GET /hello               | /hello           | {"message":"Hello, World"}        |
| GET /greeting/{name}     | /greeting/Jane   | {"greeting":"Hello, Jane!"}       |
| GET /fibonacci/{number}  | /fibonacci/45    | {"input":45,"output":1134903170}  |
|       |        |         |


## Notes:
- The code is not suitable for production use. And is not idiomatic to each language. Also, lacks tests and documentation.
- This project I wrote to compare some metrics using different programming languages. The metrics I was looking for:
    - Compiled artifact size
    - Size in memory while running idle
    - Size in memory while running at peak
    - Average requests per second




