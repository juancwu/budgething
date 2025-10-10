package assets

import "embed"

//go:embed js/* css/*
var AssetsFS embed.FS
