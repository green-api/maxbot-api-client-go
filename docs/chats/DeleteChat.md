# DeleteChat

[dev.max.ru/docs-api/methods/DELETE/chats/-chatId-](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-)

Удаляет групповой чат для всех участников

## Запрос {#request}  

### Поля запроса  {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                    |
|----------|-------------|--------------|-----------------------------|
| `ChatID` | **integer** | Да           | ID удаляемого чата          |

### Пример запроса {#request-example}  

```go
response, err := bot.Chats.DeleteChat(
    context.Background(),
        &models.DeleteChatReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

| Поле      | Тип         | Описание                                                 |
|-----------|-------------|----------------------------------------------------------|
| `success` | **boolean** | `true`, если запрос был успешным,`false` — в противном случае                       |
| `message` | **string**  | Объяснительное сообщение, если результат не был успешным |

### Пример тела ответа {#response-example-body}

**Успех:**

```json
{
    "success": true
}
```

**Ошибка:**

```json
{
    "success": false,
    "message": "error.chat.not_found"
}
```