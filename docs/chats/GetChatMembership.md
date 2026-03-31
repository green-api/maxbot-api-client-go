# GetChatMembership

[dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/me](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/me)

Метод предназначен для получения информации о статусе самого бота в указанном чате (роль, права).

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                        |
|----------|-------------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID чата                         |

### Пример запроса {#request-example}  


```go
response, err := bot.Chats.GetChatMembership(
    context.Background(), 
    &models.GetChatMembershipReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

[**Объект `ChatMember`**](../objects/User.md#chatmember-chatmember)

| Поле                 | Тип         | Описание                                                                             |
|----------------------|-------------|--------------------------------------------------------------------------------------|
| `user_id`            | **integer** | Идентификатор пользователя или бота                                                  |
| `first_name`         | **string**  | Отображаемое имя пользователя или бота                                               |
| `last_name`          | **string**  | Отображаемая фамилия пользователя. Для ботов это поле не возвращается                |
| `username`           | **string**  | Никнейм бота или уникальное публичное имя пользователя                               |
| `is_bot`             | **boolean** | `true`, если это бот                                                                 |
| <nobr>`last_activity_time`<nobr> | **integer** | Время последней активности пользователя или бота в MAX (Unix-время в миллисекундах)  |
| `description`        | **string**  | Описание пользователя или бота (до 16000 символов)                                   |
| `avatar_url`         | **string**  | URL аватара пользователя или бота в уменьшенном размере                              |
| `full_avatar_url`    | **string**  | URL аватара пользователя или бота в полном размере                                   |
| `last_access_time`   | **integer** | Время последней активности пользователя в чате. Может быть устаревшим для суперчатов (равно времени вступления) |
| `is_owner`           | **boolean** | Является ли пользователь владельцем чата                                             |
| `is_admin`           | **boolean** | Является ли пользователь администратором чата                                        |
| `join_time`          | **integer** | Дата присоединения к чату в формате Unix time                                        |
| `permissions`        | **object**  | Перечень прав пользователя. Возможные значения:                                      |
|                      |             | `read_all_messages` — Читать все сообщения                                           |
|                      |             | `add_remove_members` — Добавлять/удалять участников                                  |
|                      |             | `add_admins` — Добавлять администраторов                                             |
|                      |             | `change_chat_info` — Изменять информацию о чате                                      |
|                      |             | `pin_message` — Закреплять сообщения                                                 |
|                      |             | `write` — Писать сообщения                                                           |
|                      |             | `edit_link` — Изменять ссылку на чат                                                 |
| `alias`              | **string**  | Заголовок, который будет показан на клиенте. Если пользователь администратор или владелец и ему не установлено это название, то поле не передаётся, клиенты на своей стороне подменят на "владелец" или "админ" |

### Пример тела ответа {#response-example-body}

```json
{
    "last_access_time": 0,
    "is_owner": false,
    "is_admin": true,
    "join_time": 1773988397836,
    "permissions": [
        "add_remove_members",
        "change_chat_info",
        "read_all_messages",
        "can_call",
        "add_admins",
        "edit_link",
        "write",
        "pin_message"
    ],
    "user_id": 123456789,
    "first_name": "Green-API boy",
    "username": "id123456789_1_bot",
    "is_bot": true,
    "last_activity_time": 1774253997680,
    "description": "Bot created by Green-API",
    "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5HiiuX5fPo",
    "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tfPo"
}
```
