package main

var CreateBorrowRecord Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Create",
    Fields: []Field{
        {
            Required: true,
            Comment: "借出物品ID",
            Name: "ItemID",
            Alias: "item_id",
            Type: "uint",
        },{
            Required: true,
            Comment: "借出時間",
            Name: "BorrowDate",
            Alias: "borrow_date",
            Type: "time.Time",
        },{
            Required: false,
            Comment: "收回物品時間",
            Name: "ReplyDate",
            Alias: "reply_date",
            Type: "time.Time",
        },{
            Required: false,
            Comment: "備註",
            Name: "Note",
            Alias: "note",
            Type: "string",
        },
    },
}