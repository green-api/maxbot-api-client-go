# PinMessage

[dev.max.ru/docs-api/methods/PUT/chats/-chatId-/pin](https://dev.max.ru/docs-api/methods/PUT/chats/-chatId-/pin)

Закрепляет сообщение в групповом чате

## Запрос {#request}

### Поля запроса{#request-parameters}

| Поле        | Тип         | Обязательный | Описание                                                 |
|-------------|-------------|--------------|----------------------------------------------------------|
| `ChatID`    | **integer** | Да           | ID чата                                                  |
| `MessageID` | **string**  | Да           | ID сообщения, которое нужно закрепить                    |
| `Notify`    | **boolean** | Нет          | Если `true`, участники получат уведомление с системным сообщением о закреплении (по умолчанию `true`)|

### Пример запроса {#request-example}  

```go
response, err := bot.Chats.PinMessage(
    context.Background(), 
    &models.PinMessageReq{
        ChatID:    123456789,
        MessageID: "mid.000000000782a4e0019d00d3ef744e91",
        Notify:    true,
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
