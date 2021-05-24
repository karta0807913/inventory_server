package main

var FirstItemTable Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Search",
    Fields: []Field{
        {
            Required: false,
            Comment: "物品ID",
            Name: "ID",
            Alias: "id",
            Type: "uint",
        },{
            Required: false,
            Comment: "學校產條上的ID",
            Name: "ItemID",
            Alias: "item_id",
            Type: "string",
        },
    },
}