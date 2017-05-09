# Заготовка микросервиса с json-rpc

## Запуск локально
```bash
$ gvt restore
$ go build ./main.go
$ ./main -c config-local.toml
```
-c - флаг для указания конфигурационного файла. По-умолчанию ищем config.toml

## Ручки

### Диагностика
Запрос:
```bash
$ curl -X GET "http://127.0.0.1:3042/health-check"
```

### Тестовая
Запрос:
```bash
$ curl -X POST \
    http://127.0.0.1:3042/api/v1 \
    -H 'content-type: application/json' \
    -d '{
          "jsonrpc": "2.0",
          "method": "test.test",
          "params": {
          	"param": "test"
          },
          "id": 1
        }'
```

Ответ:
```json
{
  "jsonrpc": "2.0",
  "result": {
    "param": "test"
  },
  "id": 1
}
```