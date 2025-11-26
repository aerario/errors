package errors

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorKindSuite struct {
	suite.Suite
}

func TestErrorKindSuite(t *testing.T) {
	suite.Run(t, new(ErrorKindSuite))
}

func (suite *ErrorKindSuite) TestAuthenticationError() {
	var err = NewAuthenticationError("something %v", "Authentication")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Authentication")
	suite.Require().Equal(ErrKindAuthentication, KindOf(err))
}

func (suite *ErrorKindSuite) TestAuthenticationFactory() {
	var fac = NewAuthenticationFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Authentication")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Authentication")
	suite.Require().Equal(t.String(), "something Authentication")
	suite.Require().Equal(ErrKindAuthentication, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestAuthorizationError() {
	var err = NewAuthorizationError("something %v", "Authorization")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Authorization")
	suite.Require().Equal(ErrKindAuthorization, KindOf(err))
}

func (suite *ErrorKindSuite) TestAuthorizationFactory() {
	var fac = NewAuthorizationFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Authorization")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Authorization")
	suite.Require().Equal(t.String(), "something Authorization")
	suite.Require().Equal(ErrKindAuthorization, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestBadRequestError() {
	var err = NewBadRequestError("something %v", "BadRequest")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something BadRequest")
	suite.Require().Equal(ErrKindBadRequest, KindOf(err))
}

func (suite *ErrorKindSuite) TestBadRequestFactory() {
	var fac = NewBadRequestFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("BadRequest")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something BadRequest")
	suite.Require().Equal(t.String(), "something BadRequest")
	suite.Require().Equal(ErrKindBadRequest, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestValidationError() {
	var err = NewValidationError("something %v", "Validation")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Validation")
	suite.Require().Equal(ErrKindValidation, KindOf(err))
}

func (suite *ErrorKindSuite) TestValidationFactory() {
	var fac = NewValidationFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Validation")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Validation")
	suite.Require().Equal(t.String(), "something Validation")
	suite.Require().Equal(ErrKindValidation, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestNotFoundError() {
	var err = NewNotFoundError("something %v", "NotFound")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something NotFound")
	suite.Require().Equal(ErrKindNotFound, KindOf(err))
}

func (suite *ErrorKindSuite) TestNotFoundFactory() {
	var fac = NewNotFoundFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("NotFound")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something NotFound")
	suite.Require().Equal(t.String(), "something NotFound")
	suite.Require().Equal(ErrKindNotFound, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestAlreadyExistsError() {
	var err = NewAlreadyExistsError("something %v", "AlreadyExists")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something AlreadyExists")
	suite.Require().Equal(ErrKindAlreadyExists, KindOf(err))
}

func (suite *ErrorKindSuite) TestAlreadyExistsFactory() {
	var fac = NewAlreadyExistsFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("AlreadyExists")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something AlreadyExists")
	suite.Require().Equal(t.String(), "something AlreadyExists")
	suite.Require().Equal(ErrKindAlreadyExists, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestLimitExceededError() {
	var err = NewLimitExceededError("something %v", "LimitExceeded")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something LimitExceeded")
	suite.Require().Equal(ErrKindLimitExceeded, KindOf(err))
}

func (suite *ErrorKindSuite) TestLimitExceededFactory() {
	var fac = NewLimitExceededFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("LimitExceeded")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something LimitExceeded")
	suite.Require().Equal(t.String(), "something LimitExceeded")
	suite.Require().Equal(ErrKindLimitExceeded, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestInconsistentError() {
	var err = NewInconsistentError("something %v", "Inconsistent")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Inconsistent")
	suite.Require().Equal(ErrKindInconsistent, KindOf(err))
}

func (suite *ErrorKindSuite) TestInconsistentFactory() {
	var fac = NewInconsistentFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Inconsistent")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Inconsistent")
	suite.Require().Equal(t.String(), "something Inconsistent")
	suite.Require().Equal(ErrKindInconsistent, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestPersistenceError() {
	var err = NewPersistenceError("something %v", "Persistence")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Persistence")
	suite.Require().Equal(ErrKindPersistence, KindOf(err))
}

func (suite *ErrorKindSuite) TestPersistenceFactory() {
	var fac = NewPersistenceFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Persistence")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Persistence")
	suite.Require().Equal(t.String(), "something Persistence")
	suite.Require().Equal(ErrKindPersistence, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestInfrastructureError() {
	var err = NewInfrastructureError("something %v", "Infrastructure")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Infrastructure")
	suite.Require().Equal(ErrKindInfrastructure, KindOf(err))
}

func (suite *ErrorKindSuite) TestInfrastructureFactory() {
	var fac = NewInfrastructureFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Infrastructure")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Infrastructure")
	suite.Require().Equal(t.String(), "something Infrastructure")
	suite.Require().Equal(ErrKindInfrastructure, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestThirdPartiesError() {
	var err = NewThirdPartiesError("something %v", "ThirdParties")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something ThirdParties")
	suite.Require().Equal(ErrKindThirdParties, KindOf(err))
}

func (suite *ErrorKindSuite) TestThirdPartiesFactory() {
	var fac = NewThirdPartiesFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("ThirdParties")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something ThirdParties")
	suite.Require().Equal(t.String(), "something ThirdParties")
	suite.Require().Equal(ErrKindThirdParties, KindOf(err))
	suite.Require().True(Is(err, fac))
}

func (suite *ErrorKindSuite) TestTimeoutError() {
	var err = NewTimeoutError("something %v", "Timeout")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), DefaultUserFriendlyError)
	suite.Require().Equal(t.String(), "something Timeout")
	suite.Require().Equal(ErrKindTimeout, KindOf(err))
}

func (suite *ErrorKindSuite) TestTimeoutFactory() {
	var fac = NewTimeoutFactory("something %v")
	suite.Require().NotNil(fac)

	var err = fac.New("Timeout")
	suite.Require().NotNil(err)

	var t, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().Equal(err.Error(), "something Timeout")
	suite.Require().Equal(t.String(), "something Timeout")
	suite.Require().Equal(ErrKindTimeout, KindOf(err))
	suite.Require().True(Is(err, fac))
}
