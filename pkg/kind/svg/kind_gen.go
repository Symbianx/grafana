// This will be autogenerated

package svg

import (
	"embed"

	"github.com/grafana/grafana/pkg/framework/kind"
)

//go:embed kind.cue
var cueFS embed.FS

type Kind struct {
	meta kind.RawMeta
}

var _ kind.Raw = &Kind{}

// TODO standard generated docs
func NewKind() (*Kind, error) {
	kdef, err := kind.ParseKindFS(cueFS, nil)
	if err != nil {
		return nil, err
	}

	k := &Kind{
		meta: kdef.Meta.(kind.RawMeta),
	}
	return k, nil
}

// TODO standard generated docs
func (k *Kind) Name() string {
	return "svg"
}

// TODO standard generated docs
func (k *Kind) Maturity() kind.Maturity {
	return k.meta.Maturity
}

// TODO standard generated docs
func (k *Kind) Meta() kind.RawMeta {
	return k.meta
}
