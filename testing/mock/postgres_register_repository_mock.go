package mock

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/novabankapp/usermanagement.data/domain/registration"
	"reflect"
)

type PostgresRegisterRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *PostgresRegisterRepositoryMockRecorder
}
type PostgresRegisterRepositoryMockRecorder struct {
	mock *PostgresRegisterRepositoryMock
}

func NewMockUserPGRepository(ctrl *gomock.Controller) *PostgresRegisterRepositoryMock {
	mock := &PostgresRegisterRepositoryMock{ctrl: ctrl}
	mock.recorder = &PostgresRegisterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *PostgresRegisterRepositoryMock) EXPECT() *PostgresRegisterRepositoryMockRecorder {
	return m.recorder
}

func (m *PostgresRegisterRepositoryMock) Create(ctx context.Context, user registration.User) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*registration.User)
	ret1, _ := ret[1].(error)
	return &ret0.ID, ret1
}

// Create indicates an expected call of Create
func (mr *PostgresRegisterRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).Create), ctx, user)
}

func (m *PostgresRegisterRepositoryMock) VerifyEmail(ctx context.Context, email string, code string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmail", ctx, email, code)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) VerifyEmail(ctx, email interface{}, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmail",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).VerifyEmail), ctx, email, code)
}

func (m *PostgresRegisterRepositoryMock) VerifyPhone(ctx context.Context, phone string, code string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPhone", ctx, phone, code)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) VerifyPhone(ctx, phone interface{}, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPhone",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).VerifyPhone), ctx, phone, code)
}

func (m *PostgresRegisterRepositoryMock) VerifyOTP(cxt context.Context, userId string, pin string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOTP", cxt, userId, pin)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) VerifyOTP(ctx, userId interface{}, pin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOTP",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).VerifyOTP), ctx, userId, pin)
}

func (m *PostgresRegisterRepositoryMock) SaveDetails(cxt context.Context, userId string, details registration.UserDetails) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDetails", cxt, userId, details)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveDetails(ctx, userId interface{}, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDetails",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveDetails), ctx, userId, details)
}

func (m *PostgresRegisterRepositoryMock) SaveResidenceDetails(cxt context.Context, userId string, details registration.ResidenceDetails) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResidenceDetails", cxt, userId, details)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveResidenceDetails(ctx, userId interface{}, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResidenceDetails",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveResidenceDetails), ctx, userId, details)
}

func (m *PostgresRegisterRepositoryMock) SaveUserIdentification(cxt context.Context, userId string, identification registration.UserIdentification) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserIdentification", cxt, userId, identification)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveUserIdentification(ctx, userId interface{}, identification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserIdentification",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveUserIdentification), ctx, userId, identification)
}

func (m *PostgresRegisterRepositoryMock) SaveUserIncome(cxt context.Context, userId string, income registration.UserIncome) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserIncome", cxt, userId, income)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveUserIncome(ctx, userId interface{}, income interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserIncome",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveUserIncome), ctx, userId, income)
}

func (m *PostgresRegisterRepositoryMock) SaveEmployment(cxt context.Context, userId string, employment registration.UserEmployment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEmployment", cxt, userId, employment)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveEmployment(ctx, userId interface{}, employment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEmployment",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveEmployment), ctx, userId, employment)
}

func (m *PostgresRegisterRepositoryMock) SaveContact(cxt context.Context, userId string, contact registration.Contact) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveContact", cxt, userId, contact)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *PostgresRegisterRepositoryMockRecorder) SaveContact(ctx, userId interface{}, contact interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveContact",
		reflect.TypeOf((*PostgresRegisterRepositoryMock)(nil).SaveContact), ctx, userId, contact)
}
