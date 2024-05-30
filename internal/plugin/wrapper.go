package plugin

import (
	"github.com/goexl/authn"
	"github.com/pangum/session/internal/config"
)

type Wrapper struct {
	*authn.Config `validate:"required"`

	Session *config.Session `json:"session,omitempty" yaml:"session" xml:"session" toml:"session" validate:"required"`
}
