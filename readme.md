# Aero Restaurant Reservation REST API
## Test project


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Небольшая система бронирования столиков

- БД PostgreSQL
- REST API
- Golang
- docker-compose
- yml configs && .env

## Features

- Можно зарезервировать столик (POST /restaurants)
- Можно запросить все доступные рестораны и столики в указанное время для компании людей (POST /restaurants/available)

POST: http://localhost:8000/restaurants/available
```json
{
    "people_quantity": 5,
    "time": "[2022-06-25 15:30, 2022-06-25 17:30)"
}
```
Response:
```json
{
  "restaurants": [
    {
      "name": "Молодость",
      "location": "ул Московская",
      "avg_waiting_time": 15,
      "avg_bill_amount": 1000,
      "available_tables": [
        {
          "id": 11,
          "restaurant": 2,
          "capacity": 3
        },
        {
          "id": 12,
          "restaurant": 2,
          "capacity": 3
        }
      ]
    }
  ]
}
```

POST: http://localhost:8000/restaurants
```json
{
  "restaurant": 1,
  "customer_name": "Nail",
  "customer_phone": "+7931233327",
  "table": 1,
  "time": "[2022-06-25 15:30, 2022-06-25 17:30)"
}
```
Response:
```json
{
  "customer_name": "Nail",
  "customer_phone": "+7931233327",
  "id": 3,
  "restaurant": 1,
  "table": 1,
  "time": "[2022-06-25 15:30, 2022-06-25 17:30)"
}
```

## Installation

Запуск postgres контейнера
```sh
make up_build # stops docker-compose (if running), builds all projects and starts docker compose
make up # starts all containers in the background without forcing build
make down # stop docker compose
```

Без make:
```sh
docker-compose up --build -d
```

Запуск приложения:
```sh
go mod tidy
go run ./cmd
```

## Tech

| Plugin           | README |
|------------------| ------ |
| Make             | [https://www.gnu.org/software/make/][PlDb] |
| Docker           | [https://www.docker.com/][PlGh] |
| PostgreSQL       | [https://www.postgresql.org/][PlGd] |

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

[dill]: <https://github.com/joemccann/dillinger>
[git-repo-url]: <https://github.com/joemccann/dillinger.git>
[john gruber]: <http://daringfireball.net>
[df1]: <http://daringfireball.net/projects/markdown/>
[markdown-it]: <https://github.com/markdown-it/markdown-it>
[Ace Editor]: <http://ace.ajax.org>
[node.js]: <http://nodejs.org>
[Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
[jQuery]: <http://jquery.com>
[@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
[express]: <http://expressjs.com>
[AngularJS]: <http://angularjs.org>
[Gulp]: <http://gulpjs.com>

[PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
[PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
[PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
[PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
[PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
[PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
