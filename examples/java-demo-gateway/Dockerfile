FROM anapsix/alpine-java

MAINTAINER demo

RUN mkdir -p /workspace &&  \
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY . .

RUN ./gradlew build -v
EXPOSE 8080

ADD ./build/libs/demo-0.0.1-SNAPSHOT.jar ./app.jar

ENTRYPOINT ["sh","-c","java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTS -jar app.jar"]
CMD ["--spring.profiles.active=prod"]