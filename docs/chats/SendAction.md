# SendAction

[dev.max.ru/docs-api/methods/POST/chats/-chatId-/actions](https://dev.max.ru/docs-api/methods/POST/chats/-chatId-/actions)

Позволяет отправлять в групповой чат такие действия бота, как например: «набор текста» или «отправка фото»

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле      | Тип         | Обязательный | Описание                                                    |
|-----------|-------------|--------------|-------------------------------------------------------------|
| `ChatID`  | **integer** | Да           | ID чата, куда отправляется действие                         |
| `Action`  | **string**  | Да           | Действие, отправляемое участникам чата. Возможные значения: |
|           |             |              | `typing_on` — Бот набирает сообщение                        |
|           |             |              | `sending_photo` —  Бот отправляет фото                      |
|           |             |              | `sending_video` —  Бот отправляет видео                     |
|           |             |              | `sending_audio` —  Бот отправляет аудиофайл                 |
|           |             |              | `sending_file` — Бот отправляет файл                        |
|           |             |              | `mark_seen` — Бот помечает сообщения как прочитанные        |
 
### Пример запроса {#request-example}  


```go
response, err := bot.Chats.SendAction(
    context.Background(), 
    &models.SendActionReq{
        ChatID: 123456789,
        Action: "typing_on",
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
