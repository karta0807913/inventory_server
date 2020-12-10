package main

var UpdateItemTable Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
            Required: false,
            Comment: "使用年限",
            Name: "AgeLimit",
            Alias: "age_limit",
            Type: "uint",
        },{
            Required: false,
            Comment: "物品存放位置",
            Name: "Location",
            Alias: "location",
            Type: "string",
        },{
            Required: false,
            Comment: "物品現今狀態",
            Name: "State",
            Alias: "state",
            Type: "ItemState",
        },{
            Required: false,
            Comment: "備註",
            Name: "Note",
            Alias: "note",
            Type: "string",
        },
    },
}