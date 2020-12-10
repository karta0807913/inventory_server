package main

var CreateBorrower Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Create",
    Fields: []Field{
        {
            Required: true,
            Comment: "借貸人名稱",
            Name: "Name",
            Alias: "name",
            Type: "string",
        },{
            Required: true,
            Comment: "借貸人手機",
            Name: "Phone",
            Alias: "phone",
            Type: "string",
        },
    },
}