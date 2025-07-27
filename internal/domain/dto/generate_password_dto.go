// Package dto (Data Transfer Objects) contains request/response structures
// for API communication between application layers.
//
// These types are used for:
// - JSON serialization/deserialization
// - Input validation
// - Interface contracts between packages
package dto

// GeneratePasswordOptions contains password generation settings
type GeneratePasswordOptions struct {
	Length           uint8 `json:"length"`
	IncludeUppercase bool  `json:"include_uppercase"`
	IncludeLowercase bool  `json:"include_lowercase"`
	IncludeDigits    bool  `json:"include_digits"`
}
