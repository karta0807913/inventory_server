
CreateBorrowRecord

### POST `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| item_id | uint |  Y  | 借出物品ID | 
| borrow_date | time.Time |  Y  | 借出時間 | 
| reply_date | time.Time |  N  | 收回物品時間 | 
| note | string |  N  | 備註 | 


CreateBorrower

### POST `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| name | string |  Y  | 借貸人名稱 | 
| phone | string |  Y  | 借貸人手機 | 


CreateItemTable

### POST `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| item_id | string |  Y  | 學校產條上的ID | 
| name | string |  Y  | 物品名稱 | 
| date | string |  Y  | 構入日期 | 
| age_limit | uint |  Y  | 使用年限 | 
| cost | uint |  Y  | 物品價值 | 
| location | string |  Y  | 物品存放位置 | 
| state | ItemState |  Y  | 物品現今狀態 | 
| note | string |  Y  | 備註 | 


FindBorrowRecord

### GET `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| item_id | uint |  N  | 借出物品ID | 
| returned | bool |  N  | 是否歸還 | 


FindBorrower

### GET `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| name | string |  N  | 借貸人名稱 | 
| phone | string |  N  | 借貸人手機 | 


FirstBorrowRecord

### GET `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| id | uint |  N  | 借貸紀錄ID | 


FirstBorrower

### GET `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| id | uint |  N  | 借貸人ID | 
| name | string |  N  | 借貸人名稱 | 
| phone | string |  N  | 借貸人手機 | 


FirstItemTable

### GET `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| item_id | string |  Y  | 學校產條上的ID | 
| name | string |  N  | 物品名稱 | 


UpdateBorrowRecord

### PUT `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| reply_date | time.Time |  N  | 收回物品時間 | 
| note | string |  N  | 備註 | 
| returned | bool |  N  | 是否歸還 | 


UpdateBorrower

### PUT `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| name | string |  N  | 借貸人名稱 | 
| phone | string |  N  | 借貸人手機 | 


UpdateItemTable

### PUT `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| age_limit | uint |  N  | 使用年限 | 
| location | string |  N  | 物品存放位置 | 
| state | ItemState |  N  | 物品現今狀態 | 
| note | string |  N  | 備註 | 


UpdateUserData

### PUT `/`

* `//TODO: comment`

* Data Params

| parameters | type | required | note |
| ---------- | ---- | -------- | ---- |
| nickname | string |  N  | 使用者名稱 | 


