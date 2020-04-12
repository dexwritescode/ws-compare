# Build: docker build -t java/webservice:1.0 .
# Run: docker run -p 8080:8080 java/webservice:1.0

FROM maven:3.6.3-openjdk-8 AS build  
COPY src /usr/src/app/src  
COPY pom.xml /usr/src/app  
RUN mvn -f /usr/src/app/pom.xml clean package

FROM openjdk:8-jre-alpine
COPY --from=build /usr/src/app/target/webservice.jar /usr/app/webservice.jar  
EXPOSE 8080  
ENTRYPOINT ["java","-jar","/usr/app/webservice.jar"]  