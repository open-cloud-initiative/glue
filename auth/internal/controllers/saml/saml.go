package saml

import (
	"encoding/xml"

	"github.com/open-cloud-initiative/glue/auth/internal/ports"

	"github.com/gofiber/fiber/v3"
	"github.com/katallaxie/pkg/dbx"
)

// MetadataController handles SAML metadata operations.
type MetadataController struct {
	store dbx.Database[ports.ReadTx, ports.WriteTx]
}

// NewMetadataController creates a new MetadataController.
func NewMetadataController(store dbx.Database[ports.ReadTx, ports.WriteTx]) *MetadataController {
	return &MetadataController{store: store}
}

// GetMetadata retrieves SAML metadata.
func (mc *MetadataController) GetMetadata(ctx fiber.Ctx) error {
	metadataXML, err := xml.Marshal(nil)
	if err != nil {
		return err
	}

	ctx.Set("Content-Type", "application/application/xml")
	ctx.Set("Cache-Control", "public, max-age=600")

	if ctx.FormValue("download") == "true" {
		ctx.Set("Content-Disposition", `attachment; filename="metadata.xml"`)
	}

	return ctx.Send(metadataXML)
}
