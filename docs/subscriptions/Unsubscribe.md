# Unsubscribe

[dev.max.ru/docs-api/methods/DELETE/subscriptions](https://dev.max.ru/docs-api/methods/DELETE/subscriptions)

Удаляет подписку и отключает доставку обновлений бота на указанный Webhook.     

После вызова этого метода вы можете снова получать обновления методом долгого опроса [`GetUpdates`](./GetUpdates.md).

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле  | Тип        | Обязательный | Описание                                                                |
|-------|------------|--------------|-------------------------------------------------------------------------|
| `URL` | **string** | Да           | URL-адрес вебхука, который необходимо удалить из списка подписок бота.  |

### Пример запроса {#request-example}

```go
response, err := bot.Updates.Unsubscribe(
    context.Background(), 
    models.UnsubscribeReq{
        Url: "https://webhook.site/my-bot-endpoint",
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле      | Тип         | Описание                                                     |
|-----------|-------------|--------------------------------------------------------------|
| `success` | **boolean** | `true`, если запрос был успешным,`false` — в противном случае       |
| `message` | **string**  | Текстовое сообщение с подробностями выполнения или ошибкой   |

### Пример тела ответа {#response-example-body}

```json
{
    "success": true,
    "message": "Webhook has been successfully deleted"
}
```