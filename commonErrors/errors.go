package commonErrors

import (
	"errors"
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
)

var (
	ErrEndDateBeforeStartDate = errors.New("end date is before start date")
	ErrEmptyName              = errors.New("name is empty")
	ErrNameTooLong            = errors.New("name is too long, max 50 characters") // Deprecated: use StringTooLongError
	ErrDuplicate              = errors.New("duplicate")
	ErrNegativeValue          = errors.New("negative value")
	//ErrNotFound               = errors.New("not found %w %w") // Deprecated: use NotFoundError
	ErrPasswordMismatch         = errors.New("password mismatch")
	ErrUserNotFound             = errors.New("user not found")
	ErrUserNotActive            = errors.New("user not active")
	ErrNoTranslation            = errors.New("no translation available")
	ErrNoMessageToTranslate     = errors.New("no message to translate")
	ErrProtocolIsRequired       = errors.New("protocol is required")
	ErrInvalidProtocol          = errors.New("invalid protocol")
	ErrHostIsRequired           = errors.New("host is required")
	ErrInvalidHost              = errors.New("invalid host")
	ErrPortIsRequired           = errors.New("port is required")
	ErrInvalidPort              = errors.New("invalid port")
	ErrUsernameIsRequired       = errors.New("username is required")
	ErrInvalidUsername          = errors.New("invalid username")
	ErrPasswordIsRequired       = errors.New("password is required")
	ErrInvalidPassword          = errors.New("invalid password")
	ErrOriginIsRequired         = errors.New("no origin defined, and origin identifier is required")
	ErrInvalidOrigin            = errors.New("invalid origin")
	ErrContextCannotGetUserCode = errors.New("cannot get user from context")

	// Generic, commonly reused sentinel errors for DAO / database usage.
	// These are intended to be wrapped with context using the helper
	// functions below (e.g. WrapNotFoundError, WrapDAOInitialisationError).
	ErrNotFound          = errors.New("not found")
	ErrValidationFailed  = errors.New("validation failed")
	ErrInvalidField      = errors.New("invalid field")
	ErrInvalidType       = errors.New("invalid type")
	ErrDAONotInitialised = errors.New("dao not initialised")
	ErrDAOInitialisation = errors.New("dao initialisation failed")
	ErrDBConnect         = errors.New("database connect failed")
	ErrDBDisconnect      = errors.New("database disconnect failed")
	ErrDBQuery           = errors.New("database query failed")
	ErrNotImplemented    = errors.New("not implemented")
)

func ErrStringLengthExceededWrapper(err error, ln int) error {
	return fmt.Errorf("string too long, max %d characters error (%w)", ln, err)
}

func ErrNotFoundWrapper(table string, err error) error {
	return fmt.Errorf("%v not found (%w)", table, err)
}
func ErrReadWrapper(err error) error {
	return fmt.Errorf("read error (%w)", err)
}
func ErrWriteWrapper(err error) error {
	return fmt.Errorf("write error (%w)", err)
}
func ErrEmptyWrapper(err error) error {
	return fmt.Errorf("empty error (%w)", err)
}
func ErrClearWrapper(err error) error {
	return fmt.Errorf("clear error (%w)", err)
}
func ErrUpdateWrapper(err error) error {
	return fmt.Errorf("update error (%w)", err)
}
func ErrCreateWrapper(err error) error {
	return fmt.Errorf("create error (%w)", err)
}
func ErrDeleteWrapper(err error) error {
	return fmt.Errorf("delete error (%w)", err)
}
func ErrDropWrapper(err error) error {
	return fmt.Errorf("drop error (%w)", err)
}
func ErrValidationWrapper(err error) error {
	return fmt.Errorf("validate error (%w)", err)
}
func ErrDisconnectWrapper(err error) error {
	return fmt.Errorf("disconnect error (%w)", err)
}
func ErrConnectWrapper(err error) error {
	return fmt.Errorf("connect error (%w)", err)
}
func HandleGoValidatorError(err error) error {
	return nil
	// if err != nil {

	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.InfoLogger.Println(err)
	// 		return err
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		op := fmt.Sprintf("VALIDATION: Field[%s] Tag[%s] Kind[%s] Param[%s] Value[%s]", err.Field(), err.Tag(), err.Kind(), err.Param(), err.Value())
	// 		logger.InfoLogger.Println(op)

	// 	}

	// 	return err
	// }
	// return nil
}
func ErrEmailWrapper(err error) error {
	return fmt.Errorf("send email error (%w)", err)
}
func ErrIDGenerationWrapper(err error) error {
	return fmt.Errorf("ID generation error (%w)", err)
}

func ErrOSWrapper(err error) error {
	return fmt.Errorf("OS error (%w)", err)
}

func ErrMockingWrapper(err error) error {
	return fmt.Errorf("mocking error (%w)", err)
}

func ErrNotificationWrapper(err error) error {
	return fmt.Errorf("notification error (%w)", err)
}

func ErrFunctionalWrapper(err error, f string) error {
	return fmt.Errorf("functional error - %v (%w)", f, err)
}

func ErrWrapper(err error) error {
	logHandler.WarningLogger.Println("It is not advised to wrap errors without a specific error message")
	return fmt.Errorf("error (%w)", err)
}

func ErrInvalidFilterWrapper(err error, f string) error {
	return fmt.Errorf("invalid filter [%v] (%w)", f, err)
}

func ErrInvalidHttpReturnStatusWrapper(s string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v]", s)
}

func ErrInvalidHttpReturnStatusWithMessageWrapper(status, message string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v] (%v)", status, message)
}

func ErrInvalidFieldWrapper(f string) error {
	return fmt.Errorf("invalid field %v", f)
}

func ErrInvalidTypeWrapper(f, d, s string) error {
	return fmt.Errorf("invalid type for field %v (%v != %v)", f, d, s)
}

func ErrRecordNotFoundWrapper(table, field, id string) error {
	return fmt.Errorf("%v not found where (%v=%v)", table, field, id)
}

func ErrDAOUpdateAuditWrapper(table string, id any, auditErr error) error {
	return fmt.Errorf("updating %v audit failed (ID=%v) %e", table, id, auditErr)
}

func ErrDAOCreateWrapper(table string, id any, createErr error) error {
	return fmt.Errorf("creating %v failed (ID=%v) %e", table, id, createErr)
}

func ErrDAOInitialisationWrapper(table string, initErr error) error {
	return fmt.Errorf("initialising %v failed %e", table, initErr)
}

func ErrDAOCaclulationWrapper(table string, calcErr error) error {
	return fmt.Errorf("calculating %v failed %e", table, calcErr)
}

func ErrDAOValidationWrapper(table string, valErr error) error {
	return fmt.Errorf("validating %v failed %e", table, valErr)
}

func ErrDAOUpdateWrapper(table string, updateErr error) error {
	return fmt.Errorf("updating %v failed %e", table, updateErr)
}

func ErrDAODeleteWrapper(table, field string, value any, deleteErr error) error {
	return fmt.Errorf("deleting %v failed (%v=%v) %e", table, field, value, deleteErr)
}

func ErrGetWrapper(table, field string, value any, readErr error) error {
	return fmt.Errorf("reading %v failed (%v=%v) %e", table, field, value, readErr)
}

func ErrDAOAssertWrapper(table, field string, value any, assetErr error) error {
	return fmt.Errorf("asserting %v failed (%v=%v) %e", table, field, value, assetErr)
}

func ErrDAOLookupWrapper(table, field string, value any, lookupErr error) error {
	return fmt.Errorf("builing looking up for %v failed (key=%v,value=%v) %e", table, field, value, lookupErr)
}

func ErrDAONotInitialisedWrapper(table, action string) error {
	return fmt.Errorf("%v DAO not initialised (Action=%v)", table, action)
}
