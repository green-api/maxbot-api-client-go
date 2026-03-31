# LeaveChat

[dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/me](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/me)

Удаляет бота из участников группового чата

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                        |
|----------|-------------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID чата, который нужно покинуть |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.LeaveChat(
    context.Background(), 
    &models.LeaveChatReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле       | Тип         | Описание                                                          |
|------------|-------------|-------------------------------------------------------------------|
| `success`  | **boolean** | `true`, если запрос был успешным, `false` — в противном случае    |
| `message`  | **string**  | Объяснительное сообщение, если результат не был успешным          |

### Пример тела ответа {#response-example-body}

```json
{
    "success": true
}
```