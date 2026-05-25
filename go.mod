module github.com/mysteryforge/kube-detox

go 1.26

require (
	github.com/tidwall/gjson v1.19.0
	github.com/urfave/cli/v2 v2.27.7
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/yaml v1.6.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
)

replace github.com/mysteryforge/kube-detox/auditor => ./auditor

replace github.com/mysteryforge/kube-detox/auditor/logging => ./auditor/logging
