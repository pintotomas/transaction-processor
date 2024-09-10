package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		name         string
		email        *Email
		valid        bool
		expectedErrs map[string]string
	}{
		{
			name: "ValidEmail",
			email: &Email{
				From:        "sender@example.com",
				Credentials: "username:password",
				To:          "recipient@example.com",
				Message:     "Hello, this is a test email",
			},
			valid:        true,
			expectedErrs: nil,
		},
		{
			name: "MissingFrom",
			email: &Email{
				From:    "sender@example.com",
				To:      "recipient@example.com",
				Message: "Hello, this is a test email",
			},
			valid: false,
			expectedErrs: map[string]string{
				"Credentials": "Key: 'Email.Credentials' Error:Field validation for 'Credentials' failed on the 'required' tag",
			},
		},
		{
			name: "MissingCredentials",
			email: &Email{
				Credentials: "username:password",
				To:          "recipient@example.com",
				Message:     "Hello, this is a test email",
			},
			valid: false,
			expectedErrs: map[string]string{
				"From": "Key: 'Email.From' Error:Field validation for 'From' failed on the 'required' tag",
			},
		},
		{
			name: "EmptyMessage",
			email: &Email{
				From:        "sender@example.com",
				Credentials: "username:password",
				To:          "recipient@example.com",
				Message:     "",
			},
			valid: false,
			expectedErrs: map[string]string{
				"Message": "Key: 'Email.Message' Error:Field validation for 'Message' failed on the 'required' tag",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.email.Validate()

			if len(tc.expectedErrs) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)

				for _, expectedErr := range tc.expectedErrs {
					assert.Contains(t, err.Error(), expectedErr)
				}
			}
		})
	}
}
