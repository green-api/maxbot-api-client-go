# GetChatMembers

[dev.max.ru/docs-api/methods/GET/chats/-chatId-/members](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members)

Возвращает список участников группового чата

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле      | Тип         | Обязательный | Описание                                               |
|-----------|-------------|--------------|--------------------------------------------------------|
| `ChatID`  | **integer** | Да           | ID чата                                                |
| `UserIDs` | **array**   | Нет          | Список ID пользователей, чье членство нужно получить. <br>*Когда этот параметр передан, параметры `count` и `marker` игнорируются*|
| `Marker`  | **integer** | Нет          | Указатель на следующую страницу данных                 |
| `Count`   | **integer** | Нет          | Количество возвращаемых участников (по умолчанию 20)   |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.GetChatMembers(
    context.Background(), 
    &models.GetChatMembersReq{
        ChatID: 123456789,
        Count:  50,
    }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

| Поле      | Тип       | Описание                                                       |
|-----------|-----------|----------------------------------------------------------------|
| `members` | [**array**](../objects/User.md#chatmember-chatmemberr) | Массив участников группы      |
| `marker`  | **integer** | Маркер для получения следующей страницы (если есть)          |


### Пример тела ответа {#response-example-body}

```json
{
    "members": [
        {
            "last_access_time": 0,
            "is_owner": false,
            "is_admin": false,
            "join_time": 1773988397836,
            "user_id": 123456789,
            "first_name": "Green-API bot",
            "username": "id1234567890_1_bot",
            "is_bot": true,
            "last_activity_time": 1774253218499,
            "description": "Test bot",
            "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5HiiuX5O9Fsy5lfWoOPCwG1Y9LtsPSxKRIDWbrHrJ8JfPo",
            "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5Hi0Hf8i3P1bQrH_cLXNMIZVtLtsPSxKRIDWbrHrJ8JfPo"
        },
        {
            "last_access_time": 1773990351033,
            "is_owner": true,
            "is_admin": true,
            "join_time": 1773988397836,
            "permissions": [],
            "user_id": 987654321,
            "first_name": "Jane",
            "last_name": "",
            "is_bot": false,
            "last_activity_time": 1774253210000,
            "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tc",
            "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1"
        },
        {
            "last_access_time": 0,
            "is_owner": false,
            "is_admin": false,
            "join_time": 1774253213004,
            "user_id": 1234567890,
            "first_name": "John",
            "last_name": "Doe",
            "is_bot": false,
            "last_activity_time": 1774253214000,
            "avatar_url": "https://i.oneme.ru/i?r=BUFglOvkF6bn--g5U-BFgIkJFcMOCyj",
            "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcur"
        }
    ]
}
```
