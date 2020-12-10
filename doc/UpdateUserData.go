package main

var UpdateUserData Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
            Required: false,
            Comment: "使用者名稱",
            Name: "Name",
            Alias: "nickname",
            Type: "string",
        },
    },
}