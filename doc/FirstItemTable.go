package main

var FirstItemTable Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Search",
    Fields: []Field{
        {
            Required: true,
            Comment: "學校產條上的ID",
            Name: "ItemID",
            Alias: "item_id",
            Type: "string",
        },{
            Required: false,
            Comment: "物品名稱",
            Name: "Name",
            Alias: "name",
            Type: "string",
        },
    },
}