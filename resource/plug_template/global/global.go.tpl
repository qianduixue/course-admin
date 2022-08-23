package global

{{- if .HasGlobal }}

import "github.com/opisnoeasy/course-service/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}