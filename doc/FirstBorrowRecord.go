package main

var FirstBorrowRecord Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Search",
    Fields: []Field{
        {
            Required: false,
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
        },
    },
}