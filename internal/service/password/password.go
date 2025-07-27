package password

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
	"unicode"

	"github.com/NikolayStepanov/PasswordGenerator/internal/domain/dto"
	"github.com/NikolayStepanov/PasswordGenerator/internal/repository/memory"
	"github.com/sethvargo/go-password/password"
)

const (
	defaultGeneratePasswordAttempts = 1000 // Maximum password generation attempts
	digits                          = "0123456789"
)

var (
	// ErrInvalidCharacter indicates that a password contains disallowed characters
	ErrInvalidCharacter = errors.New("invalid character in password")

	// ErrPasswordGeneration indicates that password generation failed after maximum attempts
	ErrPasswordGeneration = errors.New("unable to generate password")

	// ErrInvalidNumberLength indicates invalid numeric string length
	ErrInvalidNumberLength = errors.New("number length must be between 1 and 10 digits")
)

// Service implements password generation and management business logic
type Service struct {
	passwordStorage memory.PasswordStorage
}

func passwordRequirementsMask(inclUppercase, inclLowercase, inclDigits bool) byte {
	var mask byte
	mask = byte(0)
	if inclLowercase {
		mask |= 1 << 0
	}
	if inclUppercase {
		mask |= 1 << 1
	}
	if inclDigits {
		mask |= 1 << 2
	}
	return mask
}

func checkPasswordRequirements(pass string, requirementsMask byte) (bool, error) {
	var (
		currentMask byte
		marching    bool
	)

	currentMask = byte(0)

	for _, val := range pass {
		switch {
		case unicode.IsLower(val):
			currentMask |= 1 << 0
		case unicode.IsUpper(val):
			currentMask |= 1 << 1
		case unicode.IsDigit(val):
			currentMask |= 1 << 2
		default:
			return false, ErrInvalidCharacter
		}

		if requirementsMask == currentMask {
			marching = true
			break
		}
	}
	return marching, nil
}

func generateUniqueNumericPassword(length int) (string, error) {
	if length < 1 || length > 10 {
		return "", ErrInvalidNumberLength
	}

	digitsVal := []byte(digits)

	for i := 0; i < length; i++ {
		limit := int64(10 - i)
		j, err := rand.Int(rand.Reader, big.NewInt(limit))
		if err != nil {
			return "", err
		}
		idx := i + int(j.Int64())
		digitsVal[i], digitsVal[idx] = digitsVal[idx], digitsVal[i]
	}

	return string(digitsVal[:length]), nil
}

func randIntMax10(length int) int {
	if length > 10 {
		length = 10
	}
	limit := big.NewInt(int64(length + 1))
	n, _ := rand.Int(rand.Reader, limit)
	return int(n.Int64())
}

func generatePassword(length uint8, inclUppercase, inclLowercase, inclDigits bool) (string, error) {
	var (
		numDigits int
		noUpper   bool
		pass      string
		err       error
	)

	if !inclUppercase && !inclLowercase {
		return generateUniqueNumericPassword(int(length))
	}

	if !inclDigits {
		numDigits = 0
	} else {
		numDigits = randIntMax10(int(length))
	}

	if !inclUppercase {
		noUpper = true
	}

	if !inclLowercase {
		pass, err = password.Generate(int(length), numDigits, 0, true, false)
		if err != nil {
			return "", err
		}
		pass = strings.ToUpper(pass)
	} else {
		pass, err = password.Generate(int(length), numDigits, 0, noUpper, false)
		if err != nil {
			return "", err
		}
	}
	return pass, nil
}

// GetNewPassword generates a new password using the provided options
func (ps *Service) GetNewPassword(ctx context.Context, options dto.GeneratePasswordOptions) (string, error) {
	passwordReqMask := passwordRequirementsMask(options.IncludeUppercase, options.IncludeLowercase, options.IncludeDigits)

	for i := 0; i < defaultGeneratePasswordAttempts; i++ {
		pass, err := generatePassword(
			options.Length,
			options.IncludeUppercase,
			options.IncludeLowercase,
			options.IncludeDigits,
		)
		if err != nil {
			return "", err
		}
		if ok, errCheck := checkPasswordRequirements(pass, passwordReqMask); ok {
			if errCheck != nil {
				return "", errCheck
			}
			return pass, nil
		}
	}
	return "", ErrPasswordGeneration
}

func NewPasswordService(passwordStorage *memory.PasswordStorage) *Service {
	return &Service{
		//passwordStorage: passwordStorage,
	}
}
