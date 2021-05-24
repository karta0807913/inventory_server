# common

### POST `/login`

* 登入
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

* 註冊
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

* 取得財產
* query
* note: 如沒有參數則回復全部

| 參數          | 型別    | 必須   | 備註                         |
| -------       | ------  | ------ | ---------------------------- |
| id            | int     | 否     | 財產資料庫編號，例如1        |
| item_id       | string  | 否     | 財產編號，例如3111401-47-3-3 |
| state         | Object  | 否     | 財產狀態                     |
| state.correct | boolean | -      | 符合                         |
| state.discard | boolean | -      | 報廢                         |
| state.fixing  | boolean | -      | 送修                         |
| state.unlabel | boolean | -      | 標籤未貼                     |
| name          | strings | 否     | 財產名稱                     |
| limit         | number  | 否     | 需要幾筆，預設20筆           |
| offset        | number  | 否     | offset                       |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |

* success reply
* note: 在定義`item_id`或`id`的時候為object，其他情況為array

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

* 更新財產詳細資料
* body 
* header
`Content-Type: application/json`
* note: 可選必須至少一項

| 參數          | 型別    | 必須  | 備註             |
| ------------- | ------- | ----- | ---------------- |
| item_id       | string  | 是    | 財產編號         |
| location      | string  | 否    | 擺放位置（可選） |
| note          | string  | 否    | 備註（可選）     |
| state         | object  | 否    | 自盤結果（可選） |
| state.correct | boolean | -     | 符合             |
| state.discard | boolean | -     | 報修             |
| state.fixing  | boolean | -     | 送修             |
| state.unlabel | boolean | -     | 標籤未貼         |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |
| 502         | 資料庫更新失敗         |

* success reply

```
{}
```

### POST `/api/borrower`

* 新增借出人
* body 
* header
`Content-Type: application/json`


| 參數 | 型別 | 必須 | 備註 |
| ---------- | ---- | -------- | ---- |
| name | string |  是  | 借貸人名稱 |
| phone | string |  是  | 借貸人手機 |



* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 502         | 資料庫更新失敗         |


* success reply

```
{
  "id": 1, 
  "name": "User1", 
  "phone": "09122345678"
}
```

### GET `/api/borrower`

* 取得借出人資料
* query
* note: 如沒有參數則回復全部

| 參數    | 型別   | 必須   | 備註                         |
| ------- | ------ | ------ | ---------------------------- |
| id      | number | 否     | 借出人ID                     |
| phone   | string | 否     | 借出人手機                   |
| name    | string | 否     | 借出人名稱                   |
| limit   | number | 否     | 需要幾筆，預設20筆           |
| offset  | number | 否     | offset                       |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |

* success reply
* note: 在定義`id`的時候為object，其他情況為array

```
{
  "id": 1, 
  "name": "User1", 
  "phone": "09122345678"
}
```

### PUT `/api/borrower`

* 更新借出人詳細資料
* body 
* header
`Content-Type: application/json`
* note: 可選必須至少一項


| 參數 | 型別 | 必須 | 備註 |
| ---------- | ---- | -------- | ---- |
| id | uint |  是  | 借貸人ID |
| name | string |  否  | 借貸人名稱 |
| phone | string |  否  | 借貸人手機 |


* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |
| 502         | 資料庫更新失敗         |

* success reply

```
{}
```

### GET `/api/borrow_record`

* 取得借出紀錄資料
* query
* note: 如沒有參數則回復全部

| 參數        | 型別   | 必須   | 備註                         |
| -------     | ------ | ------ | ---------------------------- |
| id          | number | 否     | 借出紀錄id                   |
| borrower_id | number | 否     | 借出人id                     |
| phone       | string | 否     | 借出人手機電話               |
| name        | string | 否     | 借出人名稱                   |
| limit       | number | 否     | 需要幾筆，預設20筆           |
| offset      | number | 否     | offset                       |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |

* success reply
* note: 在定義`id`的時候為object，其他情況為array

```
{
  "borrow_date": "2020-09-29T15:49:54.236129179+08:00", 
  "borrower_id": 1, 
  "id": 1, 
  "note": "", 
  "reply_date": "0001-01-01T00:00:00Z"
}
```

```
[
  {
    "borrow_date": "2020-09-29T19:12:12.177418328+08:00", 
    "borrower_id": 1, 
    "id": 1, 
    "note": "", 
    "reply_date": "0001-01-01T00:00:00Z", 
  }
]
```

### PUT `/api/borrow_record`

* 更新借出紀錄資料
* body 
* header
`Content-Type: application/json`
* note: 可選必須至少一項
* note: 如果想要清除reply_date的話就把returned設為true

| 參數        | 型別      | 必須     | 備註         |
| ----------  | ----      | -------- | ----         |
| id          | uint      | 是       | 借貸紀錄ID   |
| borrower_id | uint      | 否       | 借出人ID     |
| reply_date  | time.Time | 否       | 收回物品時間 |
| note        | string    | 否       | 備註         |
| returned    | bool      | 否       | 是否歸還     |

* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |
| 502         | 資料庫更新失敗         |

* success reply

```
{}
```

### POST `/api/borrow_record`

* 新增借貸紀錄
* body 
* header
`Content-Type: application/json`


| 參數 | 型別 | 必須 | 備註 |
| ---------- | ---- | -------- | ---- |
| borrower_id | uint |  是  | 借出人ID |
| item_id | uint |  是  | 借出物品ID |
| borrow_date | time.Time |  是  | 借出時間 |
| reply_date | *time.Time |  否  | 收回物品時間 |
| note | string |  否  | 備註 |


* error code

| Status Code | 備註                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |
| 502         | 資料庫更新失敗         |

* success reply

```
{
  "borrow_date": "2020-09-29T15:49:54.236129179+08:00", 
  "borrower_id": 1, 
  "id": 1, 
  "note": "", 
  "reply_date": "0001-01-01T00:00:00Z"
}
```

### GET `/api/borrower_fuzzy`

* 模糊搜尋使用者
* header
`Content-Type: application/json`

| 參數    | 型別   | 必須   | 備註         |
| ------- | ------ | ------ | ------------ |
| name    | string | 否     | 使用者名稱   |
| phone   | string | 否     | 使用者電話   |

* success reply

```
[{
  "id": 1,
  "name": "User1",
  "phone": "09122345678"
},{
  "id": 2,
  "name": "User2",
  "phone": "09122345678"
}]
```

* 備註

名稱與電話只能擇一送出，如兩者都有，則以電話為準。