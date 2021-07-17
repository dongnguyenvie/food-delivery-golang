###
```json
// curl --location --request GET 'http://localhost:8080/v1/restaurants'
{
    "data": [
        {
            "id": "3mPHhHfJLUTe1c",
            "status": 1,
            "created_at": "2021-05-07T20:19:23+07:00",
            "updated_at": "2021-05-07T20:19:23+07:00",
            "name": "Test res",
            "addr": "some where",
            "type": ""
        }
    ],
    "paging": {
        "page": 1,
        "limit": 50,
        "total": 2,
        "cursor": "",
        "next_cursor": ""
    },
    "filter": {}
}
```
###
```json
// {{FOOD_DELIVERY}}/v1/restaurants/:id

```