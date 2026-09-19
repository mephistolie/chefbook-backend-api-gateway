package contracts

import _ "embed"

// OpenAPI is the exact vendored specification pinned by chefbook.lock.json.
//
//go:embed chefbook.yaml
var OpenAPI []byte
