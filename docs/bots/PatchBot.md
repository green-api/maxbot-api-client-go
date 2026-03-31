# PatchBot 

Метод предназначен для изменения данных бота.

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле                       | Тип                                          | Обязательный | Описание                                        |
|----------------------------|----------------------------------------------|--------------|-------------------------------------------------|
| `Name`                     | **string**                                   | Нет          | Отображаемое имя бота                           |
| `Username`                 | **string**                                   | Нет          | Никнейм бота                                    |
| <nobr>`Description`</nobr> | **string**                                   | Нет          | Описание бота (до 16000 символов)               |
| `Commands`                 | [**object**](../objects/User.md#поля-объекта-botcommand-botcommand)  | Нет | Команды, поддерживаемые ботом (до 32 элементов) |

### Пример запроса {#request-example}

```go
response, err := bot.Bots.PatchBot(
    context.Background(), 
    models.BotPatch{
        Name: "New name"
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле                 | Тип         |  Описание                                                                            |
|----------------------|-------------|--------------------------------------------------------------------------------------|
| `user_id`            | **integer** | Идентификатор бота                                                                   |
| `first_name`         | **string**  | Отображаемое имя бота                                                                |
| `username`           | **string**  | Никнейм бота                                                                         |
| `is_bot`             | **boolean** | Для ботов всегда `true`                                                              |
| <nobr>`last_activity_time`</nobr>  | **integer** | Время последней активности бота в MAX (Unix-время в миллисекундах)     |
| `description`        | **string**  | Описание бота (до 16000 символов)                                                    |
| `avatar_url`         | **string**  | URL аватара бота в уменьшенном размере                                               |
| `full_avatar_url`    | **string**  | URL аватара бота в полном размере                                                    |
| `commands`           | [**object**](../objects/User.md#поля-объекта-botcommand-botcommand)  | Команды, поддерживаемые ботом (до 32 элементов)     |


### Пример тела ответа {#response-example-body}

```json
{
    "user_id": 123456789,
    "first_name": "New name",
    "username": "id1234567890_1_bot",
    "is_bot": true,
    "last_activity_time": 1773814559150,
    "description": "Test bot",
    "avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5HiiuX5O9Fsy5lfWoOPCwG1Y9LtsPSxKRIDWbrHrJ8JfPo",
    "full_avatar_url": "https://i.oneme.ru/i?r=BTFjO43w8Yr1OSJ4tcurq5Hi0Hf8i3P1bQrH_cLXNMIZVtLtsPSxKRIDWbrHrJ8JfPo"
}
```