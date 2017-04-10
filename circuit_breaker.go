package circuitbreaker

const (
	closed = iota
	open
	halfOpen
)

type Caller interface {
	Call() (interface{}, error)
}

type CircuitBreaker struct {
	state      int
	errorCount int
	caller     Caller
}

func New(c *Caller) CircuitBreaker {
	return CircuitBreaker{state: closed, caller: c}
}

func (cb *CircuitBreaker) Call() interface{} {
	response, err := caller.Call()
	if err != nil {
		if errorCount == 0 {
			// start interval to check counter vs some threshold  and reset if fine
		}
		errorCount += 1
	}

	// if we got a response assume it should be returned regardless of whether we got an error
	// the external caller will know how to type convert this response
	return response
}
