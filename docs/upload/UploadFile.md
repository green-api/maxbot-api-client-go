# UploadFile

[dev.max.ru/docs-api/methods/POST/uploads](https://dev.max.ru/docs-api/methods/POST/uploads)

Загружает файл на сервер MAX.   

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле         | Тип        | Обязательный | Описание                                                                                                          |
|--------------|------------|--------------|-------------------------------------------------------------------------------------------------------------------|
| `Type`       | **string** | Да           | Тип вложения. Допустимые значения: `image`, `video`, `audio`, `file`.                                             |
| `FilePath`  | **string** | Да           | Абсолютный или относительный путь к файлу на локальном диске.                                                     |
| <nobr>`UploadUrl`<nobr> | **string** | Нет | Используется только для прямой загрузки, если URL уже получен. При вызове `UploadFile` заполняется автоматически. |

### Пример запроса {#request-example}


```go
response, err := bot.Uploads.UploadFile(
  context.Background(), 
  models.UploadFileReq{
      Type:     m.UploadImage,
      FilePath: "./photos/cat.png",
  }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

| Поле      | Тип         | Описание                                                                                                                        |
|-----------|-------------|---------------------------------------------------------------------------------------------------------------------------------|
| `fileId`  | **integer** | Уникальный числовой идентификатор файла на сервере. <br>*Возвращается для типов `file`, `video` и `audio`*                      |
| `token`   | **string**  | Строковый токен доступа, необходимый для прикрепления файла к сообщению. <br>*Возвращается для типов `file`, `video` и `audio`* |
| `photos`  | [**array**](#поля-объекта-photodata-photodata) | Словарь версий изображений. Ключ — уникальный хеш-идентификатор версии. <br>*Возвращается для типа `image`*         |

### Поля объекта `PhotoData` {#PhotoData}

| Поле      | Тип         | Описание                                       |
|-----------|-------------|------------------------------------------------|
| `token`   | **string**  | Токен загруженного изображения для отправки    |

### Пример тела ответа {#response-example-body}

**Пример ответа на загрузку изображения:**

```json
{
  "photos": {
    "3JJvlQgoPpDgwtDEnxfLhI9pXye9WNfuMYfhNsaldssWmWt+o4h5tQ==": {
      "token": "9+eyCStTAIyomEW0h0PzC3/2FMn6jnei4sbeAHAHzEAQPw5R+QRmK6fI2i8m2g385JIOot9+dE50Zpu2C8rWpdL3lqBMBwKZg1LgN6LAi+L7RMzvuMlRQBGMRVsHIM3qCYjJVz9rZ3bJC2BKsTpwGFT2zeZD65ubBd8t+paTYZFhHQTAHjP3XMxDQF5b201h1MnTVAJwIbBOQi3emv+C4/zX+iKFArgxeFJdUr/8UEo="
    }
  }
}
```
**Пример ответа на загрузку файла:**

```json
{
  "fileId": 2947318223,
  "token": "f9LHodD0cOKN7LWyX9hfc2icprCO5CThP_QcGNvMN5wvpD7BvhUOjNkvsdVTR5qxMvgJwgwwvqLLqXcnEGSE"
}
```