package pool

// PolicyName is a named policy
type PolicyName string

const (
	// SendTx is the name of the policy that governs that an address may send transactions to pool
	SendTx PolicyName = "send_tx"
	// Deploy is the name of the policy that governs that an address may deploy a contract
	Deploy PolicyName = "deploy"
)
