package memo

import "fmt"

type result struct {
	value interface{}
	err   error
}

type request struct {
	key       string
	done      <-chan struct{}
	responses chan<- result
}

type entry struct {
	res   result
	ready chan struct{}
}

// Func is the type of function to memorize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// Memo is the type that manages caches of the result returned by Func.
type Memo struct {
	requests chan request
}

// New returns a Memo instance of Func f.
func New(f Func) *Memo {
	memo := &Memo{make(chan request)}
	go memo.loop(f)
	return memo
}

// Get returns the result of `memo` by invoking Func with parameter `key`
func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, done, response}
	res := <-response
	return res.value, res.err
}

// Close closes the memo to stop everything.
func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) loop(f Func) {
	caches := make(map[string]*entry)

	for req := range memo.requests {
		key := req.key
		e, ok := caches[key]
		if !ok {
			e = &entry{ready: make(chan struct{})}
			caches[key] = e
			go e.call(f, key, req.done)
		} else {
			if e.res.err == ErrorDone {
				// if done error result is cached, the result need to be calculated again
				e.ready = make(chan struct{}) // ready must be a new one
				go e.call(f, key, req.done)
			}
		}
		go e.deliver(req.responses, req.done)
	}
}

// ErrorDone error
var ErrorDone error = fmt.Errorf("done")

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(responses chan<- result, done <-chan struct{}) {
	select {
	case <-done:
		responses <- result{nil, ErrorDone}
	case <-e.ready:
		responses <- e.res
	}
}
