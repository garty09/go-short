![Go-short](https://img.shields.io/badge/build-go--short-brightgreen) ![Go-long](https://img.shields.io/badge/build-go--long-blue)


## **Сервис go-short**
Результат выполнения [тестового задания для backend-стажёра в команду Fintech.]()
Сервис выполняет функцию сокращения ссылок.

## **Условия**
- **Request**: POST /short {"url": "long-url-here"} =
**Response**: {"url": "short-url-here"}

- **Request**: POST /long {"url": "short-url-here"} =
**Response**: {"url": "long-url-here"}

## **Архитектура**
Взаимодействие с сервисом происходит посредством использования **REST-API**. Сервис обрабатывает входящие ссылки, и для каждого из них определяет, какой из обработчиков вызывать. Каждый из обработчиков реализует определенную бизнес-логику и использует взаимодействия с БД. По завершению работы обработчика, в ответ на запрос возвращаем результат в виде сокращённой либо длинной(полной) ссылки.

## **Требования**
> Для запуска сервиса необходимо наличие запущенного сервера БД PostgreSQL.

## **Структура проекта**
|**DIR**|**Описание**|
---| ---|
|**cmd/shortener**|код запускающий сервис|
|**cmd/migrations**|файл миграции|
|**internal/app**|обработчики сервиса|
|**internal/db**|работа с БД|

## **Документация методов API**
Все ссылки относительны http://localhost:8080

|**Method**|**HTTP-Request**|**Description**|
---| ---| ---|
|[POST]()|**POST**/app|выдача сокращенной ссылки|
|[POST]()|**POST**/app|выдача полной(длинной) ссылки|
