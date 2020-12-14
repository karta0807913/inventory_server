package main

var UpdateUserData Document = Document{
    Path: "/",
    Comment: "",
	Mode: "Updates",
    Fields: []Field{
        {
            Required: true,
            Comment: "使用者密碼",
            Name: "Password",
            Alias: "password",
            Type: "string",
        },{
            Required: false,
            Comment: "使用者名稱",
            Name: "Name",
            Alias: "nickname",
            Type: "string",
        },
    },
}