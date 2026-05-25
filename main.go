package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/mysteryforge/kube-detox/pkg/auditor"
	"github.com/mysteryforge/kube-detox/pkg/logging"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	logger := slog.New(&logging.ColorHandler{})

	cfg := struct {
		policy string
	}{}

	app := &cli.App{
		Name:  "kube-detox",
		Usage: "Audit K8s manifests",
		Commands: []*cli.Command{
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "Scan a directory of manifests",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "policy",
						Usage:       "Path to rules directory",
						Destination: &cfg.policy,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					var ruleset auditor.Ruleset

					// Load rules recursively
					if err := filepath.Walk(cfg.policy, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if !info.IsDir() && (strings.HasSuffix(info.Name(), ".yaml") || strings.HasSuffix(info.Name(), ".yml")) {
							data, _ := os.ReadFile(path)
							var rs auditor.Ruleset
							if err := yaml.Unmarshal(data, &rs); err != nil {
								logger.Error("failed to parse rule", "path", path, "error", err)
							}
							ruleset.Rules = append(ruleset.Rules, rs.Rules...)
						}
						return nil
					}); err != nil {
						return err
					}

					manifestPath := c.Args().First()
					manifest, err := os.ReadFile(manifestPath)
					if err != nil {
						logger.Error("failed to read manifest", "error", err)
						return err
					}

					logger.Info("Scanning", "manifest", manifestPath, "rules", len(ruleset.Rules))
					stats := struct {
						pass, fail, warn int
					}{}
					for _, rule := range ruleset.Rules {
						pass, desc := auditor.Validate(manifest, rule)
						if pass {
							stats.pass++
							if !strings.Contains(desc, "skipped") {
								logger.Info("[PASS]", "rule", desc)
							}
						} else {
							if rule.Severity == "warn" {
								stats.warn++
								logger.Info("[WARN]", "rule", desc)
							} else {
								stats.fail++
								logger.Info("[FAIL]", "rule", desc)
							}
						}
					}
					logger.Info("Summary", "passed", stats.pass, "failed", stats.fail, "warned", stats.warn)
					if stats.fail > 0 {
						os.Exit(1)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error("app failed", "error", err)
		os.Exit(1)
	}
}
