package main

import (
	"github.com/kyfk/log"
	"github.com/kyfk/log/format"
	"github.com/kyfk/log/level"
	"github.com/pkg/errors"
)

func main() {
	logger := log.New(
		log.MinLevel(level.Info),
		log.Format(format.JSONPretty),
		log.Metadata(map[string]interface{}{
			// "service":    "book",
			"uesr_id":    "86f32b8b-ec0d-479f-aed1-1070aa54cecf",
			"request_id": "943ad105-7543-11e6-a9ac-65e093327849",
			// "path":       "/operator/hello",
			// more other metadatas
		}),
		// log.FlattenMetadata(true),
	)

	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error(errors.New("error"))
}
