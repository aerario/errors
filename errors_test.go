package errors

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorsSuite struct {
	suite.Suite
}

func TestErrorsSuite(t *testing.T) {
	suite.Run(t, new(ErrorsSuite))
}

func (suite *ErrorsSuite) TestFrom() {
	var implErr = NewAlreadyExistsError("kek bek")
	var tests = []struct {
		name       string
		err        error
		assertFunc func(error)
	}{
		{
			name: "nil err",
			assertFunc: func(err error) {
				suite.Require().NoError(err)
			},
		},
		{
			name: "no implementation err",
			err:  errors.New("kek: bek"),
			assertFunc: func(err error) {
				suite.Require().Error(err)
				suite.Require().IsType(&implementation{}, err)
				suite.Require().False(IsUserFriendly(err))
				suite.Require().Equal(ErrKindGeneral, KindOf(err))
				suite.Require().Equal(DefaultUserFriendlyError, err.Error())
				suite.Require().Equal("kek: bek", Raw(err).Error())
			},
		},
		{
			name: "implementation err",
			err:  implErr,
			assertFunc: func(err error) {
				suite.Require().Equal(implErr, err)
			},
		},
	}

	for _, t := range tests {
		suite.Run(t.name, func() {
			var err = From(t.err)
			t.assertFunc(err)
		})
	}
}

func (suite *ErrorsSuite) TestWrap() {
	var (
		first  = NewAuthenticationError("invalid login/password")
		second = NewAuthorizationError("access denied")
	)

	suite.Require().NotNil(first)
	suite.Require().NotNil(second)

	var res = first.Wrap(second)

	suite.Require().True(Is(res, first))
	suite.Require().True(Is(res, second))
	suite.Require().Equal(DefaultUserFriendlyError, res.Error())
	suite.Require().Equal(ErrKindAuthentication, KindOf(res))
}

func (suite *ErrorsSuite) TestIs() {
	var (
		goError   = errors.New("access denied")
		authError = NewAuthorizationError("access denied")
		// type (*implementation) is comparable, so golang errors.Is compares pointers before call 'Is(target error) bool'
		authErrorClone   = NewAuthorizationError("access denied")
		authErrTemplated = NewAuthorizationError("access %s", "denied")
		thirdPartyErr    = NewThirdPartiesError("access denied")
		wrappedAuthErr   = fmt.Errorf("internal error: %w", authError)
	)

	tests := []struct {
		name   string
		source error
		target error
		want   bool
	}{
		{name: "error IS error clone", source: authError, target: authErrorClone, want: true},
		{name: "error IS NOT error templated", source: authError, target: authErrTemplated, want: false},
		{name: "go-wrapped error IS error", source: wrappedAuthErr, target: authError, want: true},
		{name: "go-wrapped error IS error clone", source: wrappedAuthErr, target: authErrorClone, want: true},
		{name: "go-wrapped error IS NOT error templated", source: wrappedAuthErr, target: authErrTemplated, want: false},
		{name: "error IS go-error", source: authError, target: goError, want: true},
		{name: "error different kind IS go-error", source: thirdPartyErr, target: goError, want: true},
		{name: "error IS NOT error different kind", source: authError, target: thirdPartyErr, want: false},
		{name: "go-wrapped error IS NOT error different kind", source: wrappedAuthErr, target: thirdPartyErr, want: false},
		{name: "go-wrapped error IS NOT error", source: nil, target: thirdPartyErr, want: false},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if got := Is(tt.source, tt.target); got != tt.want {
				suite.Fail("test failed", "Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *ErrorsSuite) TestIn() {
	var (
		already   = NewAlreadyExistsError("already")
		badReq    = NewBadRequestError("bad request").Wrap(already)
		authError = NewAuthorizationError("access denied").Wrap(badReq)
	)

	tests := []struct {
		name   string
		source error
		target []any
		want   bool
	}{
		{name: "target error IN error", source: authError, target: []any{already}, want: true},
		{name: "target error Not IN error", source: authError, target: []any{NewInconsistentError("no")}, want: false},
		{name: "target some errors IN error", source: authError, target: []any{NewInconsistentError("no"), badReq}, want: true},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if got := In(tt.source, tt.target...); got != tt.want {
				suite.Fail("test failed", "In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *ErrorsSuite) TestHumanFriendly() {
	var (
		first  = NewAuthenticationFactory("invalid login/password").New()
		second = NewAuthorizationError("access denied")
		third  = errors.New("kek bek")
	)

	suite.Require().NotNil(first)
	suite.Require().NotNil(second)

	var res = first.Wrap(second.Wrap(third))

	suite.Require().True(Is(res, first))
	suite.Require().True(Is(res, second))
	suite.Require().True(IsUserFriendly(res))
	suite.Require().False(IsUserFriendly(errors.New("something")))
	suite.Require().Equal("invalid login/password", res.Error())
	suite.Require().Equal(ErrKindAuthentication, KindOf(res))
}

func (suite *ErrorsSuite) TestRaw() {
	var err = NewAuthorizationError("access denied")
	suite.Require().NotNil(err)

	suite.Require().False(IsUserFriendly(err))
	suite.Require().Equal(DefaultUserFriendlyError, err.Error())
	suite.Require().Equal(ErrKindAuthorization, KindOf(err))

	suite.Require().Equal("access denied", Raw(err).Error())
}

func (suite *ErrorsSuite) TestStackTrace() {
	var err = &implementation{
		kind:    ErrKindAuthentication,
		message: "invalid login/password",
		labels:  LabelList{LabelUserFriendly},
		previous: &implementation{
			kind:     ErrKindAuthorization,
			message:  "access denied",
			previous: nil,
			location: location{
				file: "bek.go",
				line: 500100,
			},
		},
		location: location{
			file: "kek.go",
			line: 100500,
		},
	}

	suite.Require().True(IsUserFriendly(err))
	suite.Require().Equal("invalid login/password", err.Error())
	suite.Require().Equal(ErrKindAuthentication, KindOf(err))

	var (
		expected = []stackTraceFrame{
			{
				Kind:           "Authentication",
				IsUserFriendly: true,
				Error:          "invalid login/password",
				Location:       "kek.go:100500",
			},
			{
				Kind:     "Authorization",
				Error:    "access denied",
				Location: "bek.go:500100",
			},
		}
		resulted []stackTraceFrame
	)

	suite.Require().NoError(json.Unmarshal(err.StackTrace(), &resulted))
	suite.Require().Equal(expected, resulted)
}

func (suite *ErrorsSuite) TestRawStackTrace() {
	var err = &implementation{
		kind:    ErrKindAuthentication,
		message: "invalid login/password",
		labels:  LabelList{LabelUserFriendly},
		previous: &implementation{
			kind:     ErrKindAuthorization,
			message:  "access denied",
			previous: nil,
			location: location{
				file: "bek.go",
				line: 500100,
			},
		},
		location: location{
			file: "kek.go",
			line: 100500,
		},
	}

	suite.Require().True(IsUserFriendly(err))
	suite.Require().Equal("invalid login/password", err.Error())
	suite.Require().Equal(ErrKindAuthentication, KindOf(err))

	var (
		expected = []stackTraceFrame{
			{
				Kind:           "Authentication",
				IsUserFriendly: true,
				Error:          "invalid login/password",
				Location:       "kek.go:100500",
			},
			{
				Kind:     "Authorization",
				Error:    "access denied",
				Location: "bek.go:500100",
			},
		}
		resulted []stackTraceFrame
	)

	suite.Require().NoError(json.Unmarshal(err.StackTrace(), &resulted))
	suite.Require().Equal(expected, resulted)

	rawErr := Raw(err).(Stacker)

	suite.Require().True(reflect.DeepEqual(err.StackTrace(), rawErr.StackTrace()))
	suite.Require().Equal(err.Location(), rawErr.Location())
}

func (suite *ErrorsSuite) TestErrorCode() {
	var i = &implementation{}
	suite.Require().Equal("General", i.ErrorCode())

	i.kind = ErrKindAuthentication
	suite.Require().Equal("Authentication", i.ErrorCode())
}

func (suite *ErrorsSuite) TestMarshalJSON() {
	var fac = implementation{
		message: "",
	}

	var actual, err = fac.MarshalJSON()
	suite.Require().NoError(err)
	suite.Require().Equal(actual, []byte(`"something went wrong"`))
}

func (suite *ErrorsSuite) TestAnnotate() {
	var err = NewBadRequestFactory("something").New().Annotate("test %s", "err")
	suite.Require().Error(err)
	suite.Require().Equal(err.Error(), "something: test err")
}

func (suite *ErrorsSuite) TestDetails() {
	var err = NewBadRequestFactory("something").New().Wrap(&implementation{
		details: map[string]string{
			"field1": "test err",
		},
	})
	suite.Require().Error(err)
	suite.Require().Equal(err.Error(), "something")
	suite.Require().EqualValues(err.Details(), map[string]string{
		"field1": "test err",
	})
}
