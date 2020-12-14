package main

var UpdateBorrowRecord Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
            Required: true,
            Comment: "借貸紀錄ID",
            Name: "ID",
            Alias: "id",
            Type: "uint",
        },{
            Required: false,
            Comment: "借出人ID",
            Name: "BorrowerID",
            Alias: "borrower_id",
            Type: "uint",
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
        },{
            Required: false,
            Comment: "是否歸還",
            Name: "Returned",
            Alias: "returned",
            Type: "bool",
        },
    },
}