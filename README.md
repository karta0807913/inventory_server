# common

### POST `/login`

* header
`Content-Type: application/json`

* body 

| 參數     | 型別   | 備註 |
| -------- | ------ | ---- |
| account  | string | 帳號 |
| password | string | 密碼 |

* success reply

```
{
    "nickname": "nickname"
}
```

### POST `/sign_up`

* header
`Content-Type: application/json`

* body 

| 參數     | 型別   | 備註           |
| -------- | ------ | -------------- |
| account  | string | 帳號，最多30字 |
| password | string | 密碼           |
| name     | string | 暱稱           |

* success reply

```
{}
```

# api

### GET `/api/item`

* query

| 參數    | 型別   | 備註                         |
| ------- | ------ | ---------------------------- |
| item_id | string | 財產編號，例如3111401-47-3-3 |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |

* success reply

```
{
  "age_limit": 3, // 年限
  "cost": 10022121212, // 總價
  "date": "123asc", // 取得日期
  "id": 1, // 編號
  "item_id": "abc", // 財產編號
  "location": "e124", // 存置地點
  "name": "HI", // 財產名稱
  "note": "none", // 備註
  "state": { // 自盤結果
    "correct": false, // 符合
    "discard": false, // 報廢
    "fixing": false, // 送修
    "unlabel": true // 標籤未貼
  }
}
```

### PUT `/api/item`

* body 

| 參數          | 型別    | 備註             |
| ------------- | ------- | ---------------- |
| item_id       | string  | 財產編號         |
| location      | string  | 擺放位置（可選） |
| note          | string  | 備註（可選）     |
| state         | object  | 自盤結果（可選） |
| state.correct | boolean | 符合             |
| state.discard | boolean | 報修             |
| state.fixing  | boolean | 送修             |
| state.unlabel | boolean | 標籤未貼         |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |
| 502         | 資料庫更新失敗         |

* success reply

```
{}
```