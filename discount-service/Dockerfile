FROM openjdk:14-jdk-slim AS build-env
ADD build/libs/discount-service.jar discount-service.jar
CMD ["java", "-jar", "/discount-service.jar"]
