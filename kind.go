package errors

type Kind uint

const (
	ErrKindGeneral Kind = iota
	ErrKindAuthentication
	ErrKindAuthorization
	ErrKindBadRequest
	ErrKindValidation
	ErrKindNotFound
	ErrKindAlreadyExists
	ErrKindLimitExceeded
	ErrKindInconsistent
	ErrKindPersistence
	ErrKindInfrastructure
	ErrKindThirdParties
	ErrKindTimeout
)

func kindName(k Kind) string {
	var m = map[Kind]string{
		ErrKindGeneral:        "General",
		ErrKindAuthentication: "Authentication",
		ErrKindAuthorization:  "Authorization",
		ErrKindBadRequest:     "BadRequest",
		ErrKindValidation:     "Validation",
		ErrKindNotFound:       "NotFound",
		ErrKindAlreadyExists:  "AlreadyExists",
		ErrKindLimitExceeded:  "LimitExceeded",
		ErrKindInconsistent:   "Inconsistent",
		ErrKindPersistence:    "Persistence",
		ErrKindInfrastructure: "Infrastructure",
		ErrKindThirdParties:   "ThirdParties",
		ErrKindTimeout:        "Timeout",
	}

	return m[k]
}

func ParseKind(code string) Kind {
	var m = map[string]Kind{
		"General":        ErrKindGeneral,
		"Authentication": ErrKindAuthentication,
		"Authorization":  ErrKindAuthorization,
		"BadRequest":     ErrKindBadRequest,
		"Validation":     ErrKindValidation,
		"NotFound":       ErrKindNotFound,
		"AlreadyExists":  ErrKindAlreadyExists,
		"LimitExceeded":  ErrKindLimitExceeded,
		"Inconsistent":   ErrKindInconsistent,
		"Persistence":    ErrKindPersistence,
		"Infrastructure": ErrKindInfrastructure,
		"ThirdParties":   ErrKindThirdParties,
		"Timeout":        ErrKindTimeout,
	}

	if kind, ok := m[code]; ok {
		return kind
	}

	return ErrKindGeneral
}

// NewAuthenticationError returns an Authentication error.
func NewAuthenticationError(message string, args ...interface{}) Error {
	return New(ErrKindAuthentication, message, args...)
}

// NewAuthenticationFactory returns an error factory that creates Authentication user-friendly errors.
func NewAuthenticationFactory(template string) Factory {
	return NewFactory(ErrKindAuthentication, template)
}

// NewAuthorizationError returns an Authorization error.
func NewAuthorizationError(message string, args ...interface{}) Error {
	return New(ErrKindAuthorization, message, args...)
}

// NewAuthorizationFactory returns an error factory that creates Authorization user-friendly errors.
func NewAuthorizationFactory(template string) Factory {
	return NewFactory(ErrKindAuthorization, template)
}

// NewBadRequestError returns an BadRequest error.
func NewBadRequestError(message string, args ...interface{}) Error {
	return New(ErrKindBadRequest, message, args...)
}

// NewBadRequestFactory returns an error factory that creates BadRequest user-friendly errors.
func NewBadRequestFactory(template string) Factory {
	return NewFactory(ErrKindBadRequest, template)
}

// NewValidationError returns an Validation error.
func NewValidationError(message string, args ...interface{}) Error {
	return New(ErrKindValidation, message, args...)
}

// NewValidationFactory returns an error factory that creates Validation user-friendly errors.
func NewValidationFactory(template string) Factory {
	return NewFactory(ErrKindValidation, template)
}

// NewNotFoundError returns an NotFound error.
func NewNotFoundError(message string, args ...interface{}) Error {
	return New(ErrKindNotFound, message, args...)
}

// NewNotFoundFactory returns an error factory that creates NotFound user-friendly errors.
func NewNotFoundFactory(template string) Factory {
	return NewFactory(ErrKindNotFound, template)
}

// NewAlreadyExistsError returns an AlreadyExists error.
func NewAlreadyExistsError(message string, args ...interface{}) Error {
	return New(ErrKindAlreadyExists, message, args...)
}

// NewAlreadyExistsFactory returns an error factory that creates AlreadyExists user-friendly errors.
func NewAlreadyExistsFactory(template string) Factory {
	return NewFactory(ErrKindAlreadyExists, template)
}

// NewLimitExceededError returns an LimitExceeded error.
func NewLimitExceededError(message string, args ...interface{}) Error {
	return New(ErrKindLimitExceeded, message, args...)
}

// NewLimitExceededFactory returns an error factory that creates LimitExceeded user-friendly errors.
func NewLimitExceededFactory(template string) Factory {
	return NewFactory(ErrKindLimitExceeded, template)
}

// NewInconsistentError returns an Inconsistent error.
func NewInconsistentError(message string, args ...interface{}) Error {
	return New(ErrKindInconsistent, message, args...)
}

// NewInconsistentFactory returns an error factory that creates Inconsistent user-friendly errors.
func NewInconsistentFactory(template string) Factory {
	return NewFactory(ErrKindInconsistent, template)
}

// NewPersistenceError returns an Persistence error.
func NewPersistenceError(message string, args ...interface{}) Error {
	return New(ErrKindPersistence, message, args...)
}

// NewPersistenceFactory returns an error factory that creates Persistence user-friendly errors.
func NewPersistenceFactory(template string) Factory {
	return NewFactory(ErrKindPersistence, template)
}

// NewInfrastructureError returns an Infrastructure error.
func NewInfrastructureError(message string, args ...interface{}) Error {
	return New(ErrKindInfrastructure, message, args...)
}

// NewInfrastructureFactory returns an error factory that creates Infrastructure user-friendly errors.
func NewInfrastructureFactory(template string) Factory {
	return NewFactory(ErrKindInfrastructure, template)
}

// NewThirdPartiesError returns an ThirdParties error.
func NewThirdPartiesError(message string, args ...interface{}) Error {
	return New(ErrKindThirdParties, message, args...)
}

// NewThirdPartiesFactory returns an error factory that creates ThirdParties user-friendly errors.
func NewThirdPartiesFactory(template string) Factory {
	return NewFactory(ErrKindThirdParties, template)
}

// NewTimeoutError returns an Timeout error.
func NewTimeoutError(message string, args ...interface{}) Error {
	return New(ErrKindTimeout, message, args...)
}

// NewTimeoutFactory returns an error factory that creates Timeout user-friendly errors.
func NewTimeoutFactory(template string) Factory {
	return NewFactory(ErrKindTimeout, template)
}
