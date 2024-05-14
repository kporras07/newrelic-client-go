//go:build integration
// +build integration

package entities

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type decodeEntityGuidTestCase struct {
	name           string
	encodedGuid    string
	expectedEntity DecodedEntity
	expectedError  error
}

func TestDecodeEntityGuid(t *testing.T) {
	var testCases []decodeEntityGuidTestCase

	// Valid case
	testCases = append(testCases, decodeEntityGuidTestCase{
		name:           "Valid entity GUID",
		encodedGuid:    base64.StdEncoding.EncodeToString([]byte("12345|test_domain|user|abc123")),
		expectedEntity: DecodedEntity{12345, "test_domain", "user", "abc123"},
		expectedError:  nil,
	})

	// Missing delimiter
	testCases = append(testCases, decodeEntityGuidTestCase{
		name:           "Missing delimiter",
		encodedGuid:    "invalidentityguid",
		expectedEntity: DecodedEntity{},
		expectedError:  EntityGUIDValidationErrorTypes.INVALID_ENTITY_GUID_ERROR,
	})

	// Less than 4 parts
	testCases = append(testCases, decodeEntityGuidTestCase{
		name:           "Less than 4 parts",
		encodedGuid:    base64.StdEncoding.EncodeToString([]byte("account|domain")),
		expectedEntity: DecodedEntity{},
		expectedError:  errors.New(fmt.Sprintf("invalid entity GUID format: expected at least 4 parts delimited by '%s': %s", DELIMITER, base64.StdEncoding.EncodeToString([]byte("account|domain")))),
	})

	// Empty entity type
	testCases = append(testCases, decodeEntityGuidTestCase{
		name:           "Empty entity type",
		encodedGuid:    base64.StdEncoding.EncodeToString([]byte("12345|domain||domainId")),
		expectedEntity: DecodedEntity{},
		expectedError:  errors.New(fmt.Sprintf("%v", EntityGUIDValidationErrorTypes.EMPTY_ENTITY_TYPE_ERROR)),
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			decodedEntity, err := DecodeEntityGuid(tc.encodedGuid)

			if err != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.Equal(t, tc.expectedEntity, *decodedEntity)
			}

			if tc.expectedError == nil && (decodedEntity.AccountId != tc.expectedEntity.AccountId ||
				decodedEntity.Domain != tc.expectedEntity.Domain ||
				decodedEntity.EntityType != tc.expectedEntity.EntityType ||
				decodedEntity.DomainId != tc.expectedEntity.DomainId) {
				t.Errorf("TestCase: %s - Decoded entity does not match original entity", tc.name)
			}
		})
	}
}
