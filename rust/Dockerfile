# Build: docker build -t rust/webservice:1.0 .
# Run: docker run -p 8080:8080 rust/webservice:1.0

FROM rust:1.43.0 AS build
WORKDIR /usr/src/
RUN apt update && apt-get install -y musl-tools && rm -rf /var/lib/apt/lists/*
RUN rustup target add x86_64-unknown-linux-musl

RUN USER=root cargo new webservice
WORKDIR /usr/src/webservice
COPY Cargo.toml Cargo.lock ./
RUN cargo build --release

COPY src ./src
RUN cargo install --target x86_64-unknown-linux-musl --path .

FROM scratch
WORKDIR /app/
COPY --from=build /usr/local/cargo/bin/webservice /app/
EXPOSE 8080
USER 1000
ENTRYPOINT ["./webservice"]