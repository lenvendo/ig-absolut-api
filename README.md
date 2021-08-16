# Сервис остатков

## Кофигурация

Сервис может быть сконфигурирован тремя способами:

- через флаги
- через переменные окружения
- через файл конфига в формате toml

Возможные параметры конфигурации:

| Параметр                                  |   Переменная окружения                    | Значение по умолчанию  | Описание                              |
| -------------                             | :-------------                            | :-----                 |:-------------                         |
| server.grpc.port                          | API_SERVER_GRPC_PORT                     | 9090                   | grpc port server                      |
| server.grpc.timeout_sec                   | API_SERVER_GRPC_TIMEOUT_SEC              | 86400                  | server grpc connection timeout        |
| server.grpc.tls.enabled                   | API_SERVER_GRPC_TLS_ENABLED              | false                  | Enable or disable TLS 
| server.http.port                          | API_SERVER_HTTP_PORT                     | 8080                   | http port server                      |
| server.http.timeout_sec                   | API_SERVER_HTTP_TIMEOUT_SEC              | 86400                  | server http connection timeout    |
| sentry.enabled                            | API_SENTRY_ENABLED                       | false                  | Enables or disables sentry                            |
| sentry.dsn                                | API_SENTRY_DSN                           | https://7e67a2b5fd034e9dbb7cdc7d4cd1bccd@sentry.api.ru//11 |Sentry addres |
| sentry.environment                        | API_SENTRY_ENVIRONMENT                   | dev                    | The environment to be sent with events |
| tracer.enabled                            | API_TRACER_ENABLED                       | false                  | флаг, если указан, то в opentracing будут отправляться трассировки путей запросов (если передан через флаги, то любое значение будет соотвествоать true)     | 
| tracer.host                               | API_TRACER_HOST                          | 127.0.0.1              | хост трасировщика                                     | 
| tracer.port                               | API_TRACER_PORT                          | 5775                   | порт трасировщика                                     |
| tracer.name                               | API_TRACER_NAME                          | API                   | название трасировщика                                     | 
| metrics.enabled                           | API_METRICS_ENABLED                      | false                  | Enables or disables metric                            |
| metrics.port                              | API_METRICS_PORT                         | 9153                   | metrics server http port                              |
| logger.level                              | API_LOGGER_LEVEL                         | emerg                  | log level ([syslog](https://en.wikipedia.org/wiki/Syslog#Severity_level))              |
| logger.time.format                        | API_LOGGER_TIME_FORMST                   | 2006-01-02T15:04:05.999999999Z07:00 |[time format for logger](https://golang.org/src/time/format.go)                |

Флаги имеют наивысший приоритет, файл конфига - наинизший

### Конфигурация через флаги

```bash
$ ./api --server.grpc.port=9090 --server.http.port=8080
```````````````````

### Конфигурация через переменные окружения

Имена переменных окружения должны начинаться с префикса API, точки заменяются знаком
подчеркивания

```
API_SERVER_GRPC_PORT=9090
```

### Конфигурация через файл конфига

При запуске сервиса можно указать путь до файла с конфигурацией

```bash
$ ./api --config=./path/to/config.toml
```

Пример файла конфигурации см. в папке `/configs`: config.toml.dist

## База данных

В качестве базы данных используется Postgres.

Для миграций используется пакет [migrate](https://github.com/golang-migrate/migrate)

`/migrations` - миграции

Создать миграцию

```
$ migrate create -ext sql -dir ./migrations create_table    
```

Накатить/откатить миграции up/down

```
$ migrate -source file://migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up/down
```
