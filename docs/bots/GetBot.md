# GetBot 

[dev.max.ru/docs-api/methods/GET/me](https://dev.max.ru/docs-api/methods/GET/me)

Метод предназначен для получения данных бота.

## Запрос {#request}

### Пример запроса {#request-example}

```golang
response, err := bot.Bots.GetBot(context.Background())
```

## Ответ {#response}

### Поля ответа {#response-parameters}

[**объект `BotInfo`**](../objects/User.md#botinfo-botinfo)

| Поле               | Тип         | Описание                                                                                          |
|--------------------|-------------|---------------------------------------------------------------------------------------------------|
| `user_id`          | **integer** | Идентификатор пользователя или бота                                                               |
| `first_name`       | **string**  | Отображаемое имя пользователя или бота                                                            |
| `last_name`        | **string**  | Отображаемая фамилия пользователя. Для ботов это поле не возвращается                             |
| `username`         | **string**  | Никнейм бота или уникальное публичное имя пользователя                                           |
| `is_bot`           | **boolean** | `true`, если это бот                                                                              |
| <nobr>`last_activity_time`<nobr> | **integer** | Время последней активности пользователя или бота в MAX (Unix-время в миллисекундах) |
| `description`      | **string**  | Описание бота (до 16000 символов)                                                                 |
| `avatar_url`       | **string**  | URL аватара бота в уменьшенном размере                                                            |
| `full_avatar_url`  | **string**  | URL аватара бота в полном размере                                                                 |
| `commands`         | [**object**](../objects/User.md#поля-объекта-botcommand-botcommand)  | Команды, поддерживаемые ботом (до 32 элементов)  |

### Пример тела ответа {#response-example-body}

```json
{
    "user_id": 123456789,
    "first_name": "GREEN-API",
    "username": "id1234567890_1_bot",
    "is_bot": true,
    "last_activity_time": 1773814559150,
    "description": "Test bot",
    "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5HiiuX5O9Fsy5lfWoOPCwG1Y9LtsPSxKRIDWbrHrJ8JfPo",
    "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5Hi0Hf8i3P1bQrH_cLXNMIZVtLtsPSxKRIDWbrHrJ8JfPo"
}
```