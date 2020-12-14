package main

var FindBorrowRecord Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Search",
    Fields: []Field{
        {
            Required: false,
            Comment: "借出人ID",
            Name: "BorrowerID",
            Alias: "borrower_id",
            Type: "uint",
        },{
            Required: false,
            Comment: "借出物品ID",
            Name: "ItemID",
            Alias: "item_id",
            Type: "uint",
        },{
            Required: false,
            Comment: "是否歸還",
            Name: "Returned",
            Alias: "returned",
            Type: "bool",
        },
    },
}