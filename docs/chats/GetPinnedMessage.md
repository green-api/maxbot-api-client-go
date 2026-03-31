# GetPinnedMessage

[dev.max.ru/docs-api/methods/GET/chats/-chatId-/pin](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/pin)

Возвращает закреплённое сообщение в групповом чате

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                        |
|----------|-------------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID чата                         |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.GetPinnedMessage(
    context.Background(), 
    &models.GetPinnedMessageReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

[**объект `Message`**](../objects/Message.md)

| Поле            | Тип                                                    | Описание                                                                       |
|-----------------|--------------------------------------------------------|--------------------------------------------------------------------------------|
| `sender`        | [**object**](../objects/User.md)                       | Пользователь, отправивший сообщение                                            |
| `Recipient`     | [**object**](../objects/Message.md#поля-объекта-recipient-recipient)          | Получатель сообщения. Может быть пользователем или чатом                       |
| `timestamp`     | **integer**                                            | Время создания сообщения в формате Unix-time                                   |
| <nobr>`linked_message`</nobr> | [**object**](../objects/Message.md#поля-объекта-linkedmessage-linkedmessage) | Пересланное или ответное сообщение                                     |
| `body`          | [**object**](../objects/Message.md#поля-объекта-messagebody-messagebody)        | Содержимое сообщения                                                           |
| `stat`          | [**object**](../objects/Message.md#поля-объекта-messagestat--messagestat)        | Статистика сообщения. Возвращается только для постов в каналах                 |
| `url`           | **string**                                             | Публичная ссылка на пост на канале. Отсутствует для диалогов и групповых чатов  |

### Пример тела ответа {#response-example-body}

**Успех:**

```json
{
    "message": {
        "recipient": {
            "chat_id": -72277586598082,
            "chat_type": "chat"
        },
        "timestamp": 1773988408621,
        "body": {
            "mid": "mid.ffffbe4541fb5a5e019d09f2ed2d69f1",
            "seq": 116260104347412977,
            "text": "Hello"
        },
        "sender": {
            "user_id": 123456789,
            "first_name": "Jane",
            "last_name": "",
            "is_bot": false,
            "last_activity_time": 1774253775000
        }
    }
}
```

**Закреплённое сообщение отсутствует:**

```json
{}
```