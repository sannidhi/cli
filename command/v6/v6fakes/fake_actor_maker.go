// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/command"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeActorMaker struct {
	NewActorStub        func(command.Config, command.UI, bool) (v6.LoginActor, error)
	newActorMutex       sync.RWMutex
	newActorArgsForCall []struct {
		arg1 command.Config
		arg2 command.UI
		arg3 bool
	}
	newActorReturns struct {
		result1 v6.LoginActor
		result2 error
	}
	newActorReturnsOnCall map[int]struct {
		result1 v6.LoginActor
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActorMaker) NewActor(arg1 command.Config, arg2 command.UI, arg3 bool) (v6.LoginActor, error) {
	fake.newActorMutex.Lock()
	ret, specificReturn := fake.newActorReturnsOnCall[len(fake.newActorArgsForCall)]
	fake.newActorArgsForCall = append(fake.newActorArgsForCall, struct {
		arg1 command.Config
		arg2 command.UI
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewActor", []interface{}{arg1, arg2, arg3})
	fake.newActorMutex.Unlock()
	if fake.NewActorStub != nil {
		return fake.NewActorStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newActorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeActorMaker) NewActorCallCount() int {
	fake.newActorMutex.RLock()
	defer fake.newActorMutex.RUnlock()
	return len(fake.newActorArgsForCall)
}

func (fake *FakeActorMaker) NewActorCalls(stub func(command.Config, command.UI, bool) (v6.LoginActor, error)) {
	fake.newActorMutex.Lock()
	defer fake.newActorMutex.Unlock()
	fake.NewActorStub = stub
}

func (fake *FakeActorMaker) NewActorArgsForCall(i int) (command.Config, command.UI, bool) {
	fake.newActorMutex.RLock()
	defer fake.newActorMutex.RUnlock()
	argsForCall := fake.newActorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeActorMaker) NewActorReturns(result1 v6.LoginActor, result2 error) {
	fake.newActorMutex.Lock()
	defer fake.newActorMutex.Unlock()
	fake.NewActorStub = nil
	fake.newActorReturns = struct {
		result1 v6.LoginActor
		result2 error
	}{result1, result2}
}

func (fake *FakeActorMaker) NewActorReturnsOnCall(i int, result1 v6.LoginActor, result2 error) {
	fake.newActorMutex.Lock()
	defer fake.newActorMutex.Unlock()
	fake.NewActorStub = nil
	if fake.newActorReturnsOnCall == nil {
		fake.newActorReturnsOnCall = make(map[int]struct {
			result1 v6.LoginActor
			result2 error
		})
	}
	fake.newActorReturnsOnCall[i] = struct {
		result1 v6.LoginActor
		result2 error
	}{result1, result2}
}

func (fake *FakeActorMaker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newActorMutex.RLock()
	defer fake.newActorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeActorMaker) recordInvocation(key string, args []interface{}) {
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

var _ v6.ActorMaker = new(FakeActorMaker)
