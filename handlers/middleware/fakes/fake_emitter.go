// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/handlers/middleware"
)

type FakeEmitter struct {
	IncrementRequestCounterStub        func(delta int)
	incrementRequestCounterMutex       sync.RWMutex
	incrementRequestCounterArgsForCall []struct {
		delta int
	}
	UpdateLatencyStub        func(latency time.Duration)
	updateLatencyMutex       sync.RWMutex
	updateLatencyArgsForCall []struct {
		latency time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) IncrementRequestCounter(delta int) {
	fake.incrementRequestCounterMutex.Lock()
	fake.incrementRequestCounterArgsForCall = append(fake.incrementRequestCounterArgsForCall, struct {
		delta int
	}{delta})
	fake.recordInvocation("IncrementRequestCounter", []interface{}{delta})
	fake.incrementRequestCounterMutex.Unlock()
	if fake.IncrementRequestCounterStub != nil {
		fake.IncrementRequestCounterStub(delta)
	}
}

func (fake *FakeEmitter) IncrementRequestCounterCallCount() int {
	fake.incrementRequestCounterMutex.RLock()
	defer fake.incrementRequestCounterMutex.RUnlock()
	return len(fake.incrementRequestCounterArgsForCall)
}

func (fake *FakeEmitter) IncrementRequestCounterArgsForCall(i int) int {
	fake.incrementRequestCounterMutex.RLock()
	defer fake.incrementRequestCounterMutex.RUnlock()
	return fake.incrementRequestCounterArgsForCall[i].delta
}

func (fake *FakeEmitter) UpdateLatency(latency time.Duration) {
	fake.updateLatencyMutex.Lock()
	fake.updateLatencyArgsForCall = append(fake.updateLatencyArgsForCall, struct {
		latency time.Duration
	}{latency})
	fake.recordInvocation("UpdateLatency", []interface{}{latency})
	fake.updateLatencyMutex.Unlock()
	if fake.UpdateLatencyStub != nil {
		fake.UpdateLatencyStub(latency)
	}
}

func (fake *FakeEmitter) UpdateLatencyCallCount() int {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return len(fake.updateLatencyArgsForCall)
}

func (fake *FakeEmitter) UpdateLatencyArgsForCall(i int) time.Duration {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return fake.updateLatencyArgsForCall[i].latency
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementRequestCounterMutex.RLock()
	defer fake.incrementRequestCounterMutex.RUnlock()
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ middleware.Emitter = new(FakeEmitter)
