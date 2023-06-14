FROM maven:3.9.2-eclipse-temurin-17 AS build
WORKDIR /home/app
COPY pom.xml .
COPY src ./src
RUN mvn clean package

FROM openjdk:17
COPY --from=build /home/app/target/*.jar /usr/local/lib/app.jar
CMD ["java", "-jar", "/usr/local/lib/app.jar"]