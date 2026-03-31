# DeleteMember

[dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members)

Удаляет участника из группового чата. Для этого могут потребоваться дополнительные права

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                                              |
|----------|-------------|--------------|-------------------------------------------------------|
| `ChatID` | **integer** | Да           | ID чата                                               |
| `UserID` | **integer** | Да           | ID удаляемого пользователя                            |
| `Block`  | **boolean** | Нет          | Если `true`, пользователь будет заблокирован в чате <br>Применяется только для чатов с публичной или приватной ссылкой <br>Игнорируется в остальных случаях |


### Пример запроса {#request-example}  

```go
response, err := bot.Chats.DeleteMember(
    context.Background(), 
    &models.DeleteMemberReq{
        ChatID: 123456789,
        UserID: 77777,
        Block: true,
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