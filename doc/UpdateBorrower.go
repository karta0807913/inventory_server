package main

var UpdateBorrower Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
            Required: false,
            Comment: "借貸人名稱",
            Name: "Name",
            Alias: "name",
            Type: "string",
        },{
            Required: false,
            Comment: "借貸人手機",
            Name: "Phone",
            Alias: "phone",
            Type: "string",
        },
    },
}