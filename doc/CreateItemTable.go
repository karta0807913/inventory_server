package main

var CreateItemTable Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Create",
    Fields: []Field{
        {
            Required: true,
            Comment: "學校產條上的ID",
            Name: "ItemID",
            Alias: "item_id",
            Type: "string",
        },{
            Required: true,
            Comment: "物品名稱",
            Name: "Name",
            Alias: "name",
            Type: "string",
        },{
            Required: true,
            Comment: "構入日期",
            Name: "Date",
            Alias: "date",
            Type: "string",
        },{
            Required: true,
            Comment: "使用年限",
            Name: "AgeLimit",
            Alias: "age_limit",
            Type: "uint",
        },{
            Required: true,
            Comment: "物品價值",
            Name: "Cost",
            Alias: "cost",
            Type: "uint",
        },{
            Required: true,
            Comment: "物品存放位置",
            Name: "Location",
            Alias: "location",
            Type: "string",
        },{
            Required: true,
            Comment: "物品現今狀態",
            Name: "State",
            Alias: "state",
            Type: "ItemState",
        },{
            Required: true,
            Comment: "備註",
            Name: "Note",
            Alias: "note",
            Type: "string",
        },
    },
}