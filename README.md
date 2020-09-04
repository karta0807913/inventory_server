# common

### POST `/login`

* header
`Content-Type: application/json`

* body 

| 參數     | 型別   | 備注           |
| -------- | ------ | -------------- |
| account  | string | 帳號，最多30字 |
| password | string | 密碼           |
| name     | string | 暱稱           |

### POST `/sign_up`

* header
`Content-Type: application/json`

* body 

| 參數     | 型別   | 備注 |
| -------- | ------ | ---- |
| account  | string | 帳號 |
| password | string | 密碼 |

# api

### GET `/api/item`

* query

| 參數    | 型別   | 備注                         |
| ------- | ------ | ---------------------------- |
| item_id | string | 財產編號，例如3111401-47-3-3 |

* error code

| Status Code | 備注                   |
| ----------- | ---------------------- |
| 400         | 參數或是item not found |