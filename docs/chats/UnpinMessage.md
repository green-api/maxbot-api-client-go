# UnpinMessage

[dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/pin](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/pin)

Удаляет закреплённое сообщение в групповом чате

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле     | Тип       | Обязательный | Описание                        |
|----------|-----------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID чата                       |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.UnpinMessage(
    context.Background(), 
    &models.UnpinMessageReq{
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

```json
{
    "success": true
}
```