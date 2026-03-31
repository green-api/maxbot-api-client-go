# DeleteAdmin

[dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/admins/-userId-](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/admins/-userId-)

Отменяет права администратора у пользователя в групповом чате, лишая его административных привилегий

## Запрос {#request}  

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                        |
|----------|-------------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID чата                         |
| `UserID` | **integer** | Да           | ID пользователя                 |

### Пример запроса {#request-example}  

```go
response, err := bot.Chats.DeleteAdmin(
    context.Background(), 
    &models.DeleteAdminReq{
        ChatID: 123456789,
        UserID: 55555,
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
