// Code generated by counterfeiter. DO NOT EDIT.
package accessorfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/api/accessor"
)

type FakeActionRoleMap struct {
	RoleOfActionStub        func(string) string
	roleOfActionMutex       sync.RWMutex
	roleOfActionArgsForCall []struct {
		arg1 string
	}
	roleOfActionReturns struct {
		result1 string
	}
	roleOfActionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActionRoleMap) RoleOfAction(arg1 string) string {
	fake.roleOfActionMutex.Lock()
	ret, specificReturn := fake.roleOfActionReturnsOnCall[len(fake.roleOfActionArgsForCall)]
	fake.roleOfActionArgsForCall = append(fake.roleOfActionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RoleOfAction", []interface{}{arg1})
	fake.roleOfActionMutex.Unlock()
	if fake.RoleOfActionStub != nil {
		return fake.RoleOfActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.roleOfActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionRoleMap) RoleOfActionCallCount() int {
	fake.roleOfActionMutex.RLock()
	defer fake.roleOfActionMutex.RUnlock()
	return len(fake.roleOfActionArgsForCall)
}

func (fake *FakeActionRoleMap) RoleOfActionCalls(stub func(string) string) {
	fake.roleOfActionMutex.Lock()
	defer fake.roleOfActionMutex.Unlock()
	fake.RoleOfActionStub = stub
}

func (fake *FakeActionRoleMap) RoleOfActionArgsForCall(i int) string {
	fake.roleOfActionMutex.RLock()
	defer fake.roleOfActionMutex.RUnlock()
	argsForCall := fake.roleOfActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionRoleMap) RoleOfActionReturns(result1 string) {
	fake.roleOfActionMutex.Lock()
	defer fake.roleOfActionMutex.Unlock()
	fake.RoleOfActionStub = nil
	fake.roleOfActionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeActionRoleMap) RoleOfActionReturnsOnCall(i int, result1 string) {
	fake.roleOfActionMutex.Lock()
	defer fake.roleOfActionMutex.Unlock()
	fake.RoleOfActionStub = nil
	if fake.roleOfActionReturnsOnCall == nil {
		fake.roleOfActionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.roleOfActionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeActionRoleMap) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.roleOfActionMutex.RLock()
	defer fake.roleOfActionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeActionRoleMap) recordInvocation(key string, args []interface{}) {
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

var _ accessor.ActionRoleMap = new(FakeActionRoleMap)
