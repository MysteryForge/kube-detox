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
		policy   string
		severity string
	}{
		severity: "warn",
	}

	app := &cli.App{
		Name:  "compliance-auditor",
		Usage: "Audit K8s manifests for financial compliance",
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
					&cli.StringFlag{
						Name:        "severity",
						Usage:       "Minimum severity to report (block, warn)",
						Value:       cfg.severity,
						Destination: &cfg.severity,
					},
				},
				Action: func(c *cli.Context) error {
					var ruleset auditor.Ruleset

					// Load rules recursively
					err := filepath.Walk(cfg.policy, func(path string, info os.FileInfo, err error) error {
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
					})
					if err != nil {
						return err
					}

					manifestPath := c.Args().First()
					manifest, err := os.ReadFile(manifestPath)
					if err != nil {
						logger.Error("failed to read manifest", "error", err)
						return err
					}

					logger.Info("Scanning", "manifest", manifestPath, "rules", len(ruleset.Rules))
					for _, rule := range ruleset.Rules {
						// Severity Logic:
						// block: show only block (mandatory)
						// warn: show warn AND block (all)
						if cfg.severity == "block" && rule.Severity != "block" {
							continue
						}

						pass, desc := auditor.Validate(manifest, rule)
						if pass {
							if !strings.Contains(desc, "skipped") {
								logger.Info("[PASS]", "description", desc)
							}
						} else {
							logger.Info("[FAIL]", "description", desc, "severity", rule.Severity)
						}
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
