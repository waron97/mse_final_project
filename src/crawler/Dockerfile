FROM maven:3.9.2-eclipse-temurin-17 as build
WORKDIR /app
COPY . .
RUN mvn clean package

FROM openjdk:17
WORKDIR /app
COPY --from=build /app/target/crawler-1.0-jar-with-dependencies.jar /app/app.jar
CMD ["java", "-jar", "app.jar"]