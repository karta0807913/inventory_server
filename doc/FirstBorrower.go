package main

var FirstBorrower Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Search",
    Fields: []Field{
        {
            Required: false,
            Comment: "借貸人ID",
            Name: "ID",
            Alias: "id",
            Type: "uint",
        },{
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