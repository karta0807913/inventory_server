{{ range $key, $value := .  }}
{{ $key }}
{{ if eq $value.Mode "Create" }}{{ template "Create" $value }}{{ else if eq $value.Mode "Updates" }}{{ template "Updates" $value }}{{ else if eq $value.Mode "Search" }}{{ template "Search" $value }}{{ end }}
{{ end }}
