package vars
import "sync"
var (
	ThreadNumber = 5000
	Result    *sync.Map
)

func init() {
	Result = &sync.Map{}
}
