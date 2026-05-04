package xmdw

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platform/xenums"
	"github.com/jindasoft/jinda-platform/xlogger"
	"github.com/stoewer/go-strcase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	Env string
}
type CustomValidator struct {
	Validator *validator.Validate
	Cfg       Config
}

type Validation struct {
	Tag     string
	Handler validator.Func
}

type Opt struct {
	Validations []Validation
}

func NewCustomValidator(cfg Config, opts ...Opt) *CustomValidator {
	v := validator.New()

	// Register all custom validation functions here
	_ = v.RegisterValidation("uuid_v7", validateUuidV7)
	_ = v.RegisterValidation("objectid", validatePrimitiveObjectID)
	_ = v.RegisterValidation("is_now", validateDateTimeNow)
	_ = v.RegisterValidation("regexp", validationRegexp)
	_ = v.RegisterValidation("is_floating_point", validateIsFloatingPoint)
	_ = v.RegisterValidation("string_between", validateStringWithBetweenLength)
	_ = v.RegisterValidation("locale", validateLocale)

	var opt Opt
	if len(opts) > 0 {
		opt = opts[0]
	}

	if len(opt.Validations) > 0 {
		for _, validateFunc := range opt.Validations {
			_ = v.RegisterValidation(validateFunc.Tag, validateFunc.Handler)
		}
	}

	return &CustomValidator{
		Validator: v,
		Cfg:       cfg,
	}
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i any) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			msg := translateValidationErrors(validationErrs)
			c := strings.Join(msg, "; ")

			return fmt.Errorf("%s", c)
		}
		return err
	}
	return nil
}

func translateValidationErrors(validationErrs validator.ValidationErrors) []string {
	var messages []string
	for _, e := range validationErrs {
		fieldName := strcase.SnakeCase(e.Field())
		tag := e.Tag()
		param := e.Param()
		if len(param) > 0 {
			tag = fmt.Sprintf("%s=%s", tag, param)
		}
		value := e.Value()
		message := fmt.Sprintf(
			"invalid input on field '%s'; expected '%s', got '%v'",
			fieldName, tag, value,
		)
		messages = append(messages, message)
	}

	return messages
}

// ValidatePrimitiveObjectID : Custom validation function for PrimitiveObjectID
func validatePrimitiveObjectID(fl validator.FieldLevel) bool {
	switch fl.Field().Kind() {
	case reflect.String:
		id := fl.Field().String()
		_, err := primitive.ObjectIDFromHex(id)
		return err == nil
	case reflect.Slice, reflect.Array:
		slice := reflect.ValueOf(fl.Field().Interface())
		for i := 0; i < slice.Len(); i++ {
			id := slice.Index(i).String()
			_, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// ValidateDateTimeNow : Custom validation function for DateTimeNow can't be in the past
func validateDateTimeNow(fl validator.FieldLevel) bool {
	now := time.Now().Add(1000 * time.Millisecond) // Add delay for latency between server and client
	switch fl.Field().Kind() {
	case reflect.String:
		date := fl.Field().String()
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			return false
		}
		return t.After(now)
	case reflect.Slice, reflect.Array:
		slice := reflect.ValueOf(fl.Field().Interface())
		for i := 0; i < slice.Len(); i++ {
			date := slice.Index(i).String()
			t, err := time.Parse(time.RFC3339, date)
			if err != nil {
				return false
			}
			if !t.After(now) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func validationRegexp(fl validator.FieldLevel) bool {
	// Get the regexp pattern from the struct tag
	pattern := fl.Param()

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		// If there's an error compiling the regexp, return false
		return false
	}

	// Check if the field matches the regular expression
	return re.MatchString(fl.Field().String())
}

func validateIsFloatingPoint(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	switch fl.Field().Kind() {
	case reflect.String:
		_, err := strconv.ParseFloat(v, 64)
		return err == nil
	case reflect.Slice, reflect.Array:
		slice := reflect.ValueOf(fl.Field().Interface())
		for i := 0; i < slice.Len(); i++ {
			_, err := strconv.ParseFloat(slice.Index(i).String(), 64)
			if err != nil {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func validateUuidV7(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	switch fl.Field().Kind() {
	case reflect.String:
		uid, err := uuid.Parse(v)
		return err == nil && uid.Version() == uuid.Version(7)
	case reflect.Slice, reflect.Array:
		slice := reflect.ValueOf(fl.Field().Interface())
		for i := 0; i < slice.Len(); i++ {
			uid, err := uuid.Parse(slice.Index(i).String())
			if err != nil || uid.Version() != uuid.Version(7) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// Validate string with length , min , max optional and check if found space must trim space
//
//	eg. validate:"is_string=3 500" 3 is min length and 500 is max length
func validateStringWithBetweenLength(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	v = strings.TrimSpace(v)
	param := fl.Param()
	params := strings.Split(param, " ")
	if len(params) < 2 {
		xlogger.SysErrorf("Invalid validate tag for string_between: %s", param)
		return false
	}
	min := params[0]
	max := params[1]
	if len(min) > 0 {
		minInt, err := strconv.Atoi(min)
		if err != nil {
			return false
		}
		if len(v) < minInt {
			return false
		}
	}
	if len(max) > 0 {
		maxInt, err := strconv.Atoi(max)
		if err != nil {
			return false
		}
		if len(v) > maxInt {
			return false
		}
	}
	return true
}

func validateLocale(fl validator.FieldLevel) bool {
	lang := strings.TrimSpace(fl.Field().String())

	locale, err := xenums.FromLocaleCode(lang)
	if err != nil {
		xlogger.SysErrorf("Invalid validate tag for locale: %s", err)
		return false
	}

	return locale.IsValid()
}
