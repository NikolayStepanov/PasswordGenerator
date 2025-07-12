package password

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	"github.com/NikolayStepanov/PasswordGenerator/internal/domain/dto"
	"github.com/NikolayStepanov/PasswordGenerator/internal/repository/memory"
	"github.com/sethvargo/go-password/password"
)

type PasswordService struct {
	passwordStorage memory.PasswordStorage
}

func generateUniqueNumericPassword(length int) (string, error) {
	if length < 1 || length > 10 {
		return "", errors.New("password length must be between 1 and 10")
	}

	digits := []byte("0123456789")

	for i := 0; i < length; i++ {
		limit := int64(10 - i)
		j, err := rand.Int(rand.Reader, big.NewInt(limit))
		if err != nil {
			return "", err
		}
		idx := i + int(j.Int64())
		digits[i], digits[idx] = digits[idx], digits[i]
	}

	return string(digits[:length]), nil
}

func randIntMax10(length int) int {
	if length > 10 {
		length = 10
	}
	limit := big.NewInt(int64(length + 1))
	n, _ := rand.Int(rand.Reader, limit)
	return int(n.Int64())
}

func generatePassword(length uint8, inclUppercase bool, inclLowercase bool, inclDigits bool) (string, error) {
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

func (ps *PasswordService) GetNewPassword(ctx context.Context, options dto.GeneratePasswordOptions) (string, error) {
	return generatePassword(
		options.Length,
		options.IncludeUppercase,
		options.IncludeLowercase,
		options.IncludeDigits,
	)
}

func NewPasswordService(passwordStorage *memory.PasswordStorage) *PasswordService {
	return &PasswordService{
		//passwordStorage: passwordStorage,
	}
}
