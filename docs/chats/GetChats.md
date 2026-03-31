# GetChats

[dev.max.ru/docs-api/methods/GET/chats](https://dev.max.ru/docs-api/methods/GET/chats)

Метод предназначен для получения списка групповых чатов, в которых участвует бот, информации о каждом чате и маркера для перехода к следующей странице списка.

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                                                       |
|----------|-------------|--------------|----------------------------------------------------------------|
| `Count`  | **integer** | Нет          | Максимальное количество возвращаемых чатов (по умолчанию 50)   |
| `Marker` | **integer** | Нет          | Маркер пагинации для получения следующей страницы списка чатов |

### Пример запроса {#request-example}  

```go
response, err := bot.Chats.GetChats(
    context.Background(), 
    &models.GetChatsReq{
        Count:  20,
        Marker: 1679823456,
    }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

| Поле     | Тип        | Описание                                                    |
|----------|------------|-------------------------------------------------------------|
| `chats`  | [**array**](../objects/Chat.md) | Массив объектов с информацией о чатах  |
| `marker` | **integer** | Маркер для получения следующей страницы (если есть)        |

### Пример тела ответа {#response-example-body}

```json
{
    "chats": [
        {
            "chat_id": -72270307698082,
            "type": "chat",
            "status": "active",
            "title": "Green-API group",
            "last_event_time": 1773990351033,
            "participants_count": 4,
            "is_public": false,
            "messages_count": 4
        },
        {
            "chat_id": -72270307698082,
            "type": "chat",
            "status": "active",
            "title": "Public group",
            "last_event_time": 1773990351033,
            "participants_count": 4,
            "is_public": true,
            "messages_count": 4
        }
    ]
}
```
