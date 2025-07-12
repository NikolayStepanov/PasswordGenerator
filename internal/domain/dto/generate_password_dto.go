package dto

type GeneratePasswordOptions struct {
	Length           uint8 `json:"length"`
	IncludeUppercase bool  `json:"include_uppercase"`
	IncludeLowercase bool  `json:"include_lowercase"`
	IncludeDigits    bool  `json:"include_digits"`
}
