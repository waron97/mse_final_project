FROM gradle:8.0.0-jdk17 AS build
WORKDIR /app
COPY . .
RUN ./gradlew build

FROM openjdk:17
RUN microdnf install findutils
WORKDIR /app
COPY --from=build /app/app/build/distributions/app.tar app.tar
RUN tar -xvf app.tar
RUN rm app.tar
CMD ./app/bin/app
