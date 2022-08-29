# gRPC и REST API для https://shtrafovnet.ru/

Проект предоставляет создание двух серверов: gRPC и REST (gRPC gateway).
Реализован gRPC метод для получения информации о компании по её ИНН с помощью ресурса https://www.rusprofile.ru/.

- Сервера можно запустить локально, в докер контейнере или докер композиции.
- Доступна swagger документация, находящаяся в папке docs.

## Запуск приложения

### Быстрый запуск докер композиции

Чтобы запустить композицию:
```
docker compose build
docker compose up
```

### Локальный запуск

Чтобы запустить gRPC сервер:
```
make grpc
```

Чтобы запустить REST сервер:
```
make rest
```

## Требования

- go 1.18.2
- docker & docker-compose
- утилита make (необязательно)

## Лицензия

[MIT](https://choosealicense.com/licenses/mit/)
