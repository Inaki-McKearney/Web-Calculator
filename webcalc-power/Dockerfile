FROM maven:3.6.2-jdk-11-slim AS builder

COPY ./src /app/src

COPY pom.xml /app/pom.xml 

RUN mvn -f /app/pom.xml clean package


FROM openjdk:11.0.5-jre-stretch

COPY --from=builder /app/target/power-1.jar ./power-1.jar

EXPOSE 80  

ENTRYPOINT ["java"]

CMD ["-jar","power-1.jar","--server.port=80"]


# FROM openjdk:11.0.5-jre-stretch

# COPY ./target/power-1.jar /app/

# WORKDIR /app/

# EXPOSE 80

# ENTRYPOINT ["java"]

# CMD ["-jar","power-1.jar","--server.port=80"]