# GetChatAdmins

[dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/admins](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/admins)

Возвращает список всех администраторов группового чата. Бот должен быть администратором в запрашиваемом чате

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание   |
|----------|-------------|--------------|------------|
| `ChatID` | **integer** | Да           | ID чата    |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.GetChatAdmins(
    context.Background(), 
    &models.GetChatAdminsReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

| Поле      | Тип       | Описание                                                       |
|-----------|-----------|----------------------------------------------------------------|
| `members` | [**array**](../objects/User.md#chatmember-chatmember) | Массив администраторов группы |
| `marker`  | **integer** | Маркер для получения следующей страницы (если есть)          |

### Пример тела ответа {#response-example-body}

```json
{
    "members": [
        {
            "last_access_time": 0,
            "is_owner": false,
            "is_admin": true,
            "join_time": 1773988397836,
            "permissions": [
                "read_all_messages",
                "pin_message",
                "add_admins",
                "change_chat_info",
                "can_call",
                "write",
                "add_remove_members",
                "edit_link"
            ],
            "user_id": 123456789,
            "first_name": "Green-API bot",
            "username": "id1234567890_1_bot",
            "is_bot": true,
            "last_activity_time": 1774253351985,
            "description": "Bot created by Green-API service",
            "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8YrfPo",
            "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5H"
        },
        {
            "last_access_time": 1774253340181,
            "is_owner": true,
            "is_admin": true,
            "join_time": 1773988397836,
            "permissions": [
                "read_all_messages",
                "edit",
                "pin_message",
                "add_admins",
                "change_chat_info",
                "can_call",
                "write",
                "add_remove_members",
                "edit_link",
                "delete",
                "view_stats"
            ],
            "user_id": 987654321,
            "first_name": "Jane",
            "last_name": "",
            "is_bot": false,
            "last_activity_time": 1774253338000,
            "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OrJ8JfPo",
            "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1brHrJ8JfPo"
        }
    ]
}
```