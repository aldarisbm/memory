package memory

const (
	BoltDB     = "memoryinternal"
	BucketName = "memories"
)

type storer interface {
	saveMemoryToStore(name string, mem *Memory) error
	getMemoryFromStore(name string) (*Memory, error)
}
