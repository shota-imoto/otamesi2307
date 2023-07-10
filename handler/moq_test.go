// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/shota-imoto/otamesi2307/entity"
	"sync"
)

// Ensure, that ListTasksServiceMock does implement ListTasksService.
// If this is not the case, regenerate this file with moq.
var _ ListTasksService = &ListTasksServiceMock{}

// ListTasksServiceMock is a mock implementation of ListTasksService.
//
//	func TestSomethingThatUsesListTasksService(t *testing.T) {
//
//		// make and configure a mocked ListTasksService
//		mockedListTasksService := &ListTasksServiceMock{
//			ListTasksFunc: func(ctx context.Context) (entity.Tasks, error) {
//				panic("mock out the ListTasks method")
//			},
//		}
//
//		// use mockedListTasksService in code that requires ListTasksService
//		// and then make assertions.
//
//	}
type ListTasksServiceMock struct {
	// ListTasksFunc mocks the ListTasks method.
	ListTasksFunc func(ctx context.Context) (entity.Tasks, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListTasks holds details about calls to the ListTasks method.
		ListTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockListTasks sync.RWMutex
}

// ListTasks calls ListTasksFunc.
func (mock *ListTasksServiceMock) ListTasks(ctx context.Context) (entity.Tasks, error) {
	if mock.ListTasksFunc == nil {
		panic("ListTasksServiceMock.ListTasksFunc: method is nil but ListTasksService.ListTasks was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListTasks.Lock()
	mock.calls.ListTasks = append(mock.calls.ListTasks, callInfo)
	mock.lockListTasks.Unlock()
	return mock.ListTasksFunc(ctx)
}

// ListTasksCalls gets all the calls that were made to ListTasks.
// Check the length with:
//
//	len(mockedListTasksService.ListTasksCalls())
func (mock *ListTasksServiceMock) ListTasksCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListTasks.RLock()
	calls = mock.calls.ListTasks
	mock.lockListTasks.RUnlock()
	return calls
}

// Ensure, that AddTaskServiceMock does implement AddTaskService.
// If this is not the case, regenerate this file with moq.
var _ AddTaskService = &AddTaskServiceMock{}

// AddTaskServiceMock is a mock implementation of AddTaskService.
//
//	func TestSomethingThatUsesAddTaskService(t *testing.T) {
//
//		// make and configure a mocked AddTaskService
//		mockedAddTaskService := &AddTaskServiceMock{
//			AddTaskFunc: func(ctx context.Context, title string) (*entity.Task, error) {
//				panic("mock out the AddTask method")
//			},
//		}
//
//		// use mockedAddTaskService in code that requires AddTaskService
//		// and then make assertions.
//
//	}
type AddTaskServiceMock struct {
	// AddTaskFunc mocks the AddTask method.
	AddTaskFunc func(ctx context.Context, title string) (*entity.Task, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddTask holds details about calls to the AddTask method.
		AddTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Title is the title argument value.
			Title string
		}
	}
	lockAddTask sync.RWMutex
}

// AddTask calls AddTaskFunc.
func (mock *AddTaskServiceMock) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	if mock.AddTaskFunc == nil {
		panic("AddTaskServiceMock.AddTaskFunc: method is nil but AddTaskService.AddTask was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Title string
	}{
		Ctx:   ctx,
		Title: title,
	}
	mock.lockAddTask.Lock()
	mock.calls.AddTask = append(mock.calls.AddTask, callInfo)
	mock.lockAddTask.Unlock()
	return mock.AddTaskFunc(ctx, title)
}

// AddTaskCalls gets all the calls that were made to AddTask.
// Check the length with:
//
//	len(mockedAddTaskService.AddTaskCalls())
func (mock *AddTaskServiceMock) AddTaskCalls() []struct {
	Ctx   context.Context
	Title string
} {
	var calls []struct {
		Ctx   context.Context
		Title string
	}
	mock.lockAddTask.RLock()
	calls = mock.calls.AddTask
	mock.lockAddTask.RUnlock()
	return calls
}

// Ensure, that RegisterUserServiceMock does implement RegisterUserService.
// If this is not the case, regenerate this file with moq.
var _ RegisterUserService = &RegisterUserServiceMock{}

// RegisterUserServiceMock is a mock implementation of RegisterUserService.
//
//	func TestSomethingThatUsesRegisterUserService(t *testing.T) {
//
//		// make and configure a mocked RegisterUserService
//		mockedRegisterUserService := &RegisterUserServiceMock{
//			RegisterUserFunc: func(contextMoqParam context.Context, name string, password string, role string) (*entity.User, error) {
//				panic("mock out the RegisterUser method")
//			},
//		}
//
//		// use mockedRegisterUserService in code that requires RegisterUserService
//		// and then make assertions.
//
//	}
type RegisterUserServiceMock struct {
	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(contextMoqParam context.Context, name string, password string, role string) (*entity.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Name is the name argument value.
			Name string
			// Password is the password argument value.
			Password string
			// Role is the role argument value.
			Role string
		}
	}
	lockRegisterUser sync.RWMutex
}

// RegisterUser calls RegisterUserFunc.
func (mock *RegisterUserServiceMock) RegisterUser(contextMoqParam context.Context, name string, password string, role string) (*entity.User, error) {
	if mock.RegisterUserFunc == nil {
		panic("RegisterUserServiceMock.RegisterUserFunc: method is nil but RegisterUserService.RegisterUser was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Name            string
		Password        string
		Role            string
	}{
		ContextMoqParam: contextMoqParam,
		Name:            name,
		Password:        password,
		Role:            role,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(contextMoqParam, name, password, role)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedRegisterUserService.RegisterUserCalls())
func (mock *RegisterUserServiceMock) RegisterUserCalls() []struct {
	ContextMoqParam context.Context
	Name            string
	Password        string
	Role            string
} {
	var calls []struct {
		ContextMoqParam context.Context
		Name            string
		Password        string
		Role            string
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}
