FROM ubuntu:20.04

WORKDIR /app

RUN apt update && apt upgrade -y
RUN apt install openjdk-11-jdk -y
RUN apt install maven -y

COPY . /app

RUN mvn clean install

EXPOSE 8080

CMD ["java", "-jar", "target/jersey-1.0.jar"]
