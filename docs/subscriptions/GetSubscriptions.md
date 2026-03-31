# GetSubscriptions

[dev.max.ru/docs-api/methods/GET/subscriptions](https://dev.max.ru/docs-api/methods/GET/subscriptions)

Возвращает список всех активных подписок (Webhook), через которые бот получает обновления.

## Запрос {#request}

### Пример запроса {#request-example}

```go
response, err := bot.Updates.GetSubscriptions(context.Background())
```

## Ответ {#response}

| Поле            | Тип       | Описание                                             |
|-----------------|-----------|------------------------------------------------------|
| `subscriptions` | [**array**](#поля-объекта-subscription-subscription) | Список активных подписок            |

### Поля объекта `Subscription` {#Subscription}

| Поле           | Тип         | Описание                                                  |
|----------------|-------------|-----------------------------------------------------------|
| `url`          | **string**  | URL-адрес вебхука, на который отправляются обновления     |
| `time`         | **integer** | Unix-время, когда была создана подписка                   |
| `update_types` | **array**   | Список типов обновлений, на которые оформлена подписка    |

### Пример тела ответа {#response-example-body}

```json
{
    "subscriptions": [
        {
            "url": "https://webhook.site/my-bot-endpoint",
            "time": 1773835382644,
            "update_types": [
                "message_created",
                "bot_added"
            ]
        }
    ]
}
```
