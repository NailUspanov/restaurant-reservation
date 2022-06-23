# build a tiny docker image
FROM alpine:latest

RUN mkdir /app
COPY restaurantReservApp /app

CMD ["/app/restaurantReservApp"]