package vm
import (
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms"

	"github.com/ava-labs/hypersdk/examples/tokenvm/controller"
)

var _ vms.Factory = &Factory{}

type Factory struct{}

func (*Factory) New(logging.Logger) (interface{}, error) {
	return controller.New(), nil
}
