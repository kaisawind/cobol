package configuration

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/internal/environment/configuration/object"
	"github.com/kaisawind/cobol/asg/internal/environment/configuration/source"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ConfigurationSection struct {
	element.CobolDivisionElement
	ctx                     *cobol85.ConfigurationSectionContext
	objectComputerParagraph object.ObjectComputerParagraph
	sourceComputerParagraph source.SourceComputerParagraph
}

func NewConfigurationSection(ctx *cobol85.ConfigurationSectionContext, programUnit element.ProgramUnit) configuration.ConfigurationSection {
	return &ConfigurationSection{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}

func (e *ConfigurationSection) SetObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext) (ret object.ObjectComputerParagraph) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(object.ObjectComputerParagraph)
	} else {

	}
	return
}
