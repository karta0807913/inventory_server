package main

var UpdateBorrowRecord Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
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