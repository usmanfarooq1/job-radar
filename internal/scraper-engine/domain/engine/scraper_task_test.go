package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeScraperTask(t *testing.T) {
	t.Run("Valid creation of the scraper task", func(t *testing.T) {
		t.Parallel()
		_, err := MakeTask(3600, "Test search", "54633212", "LINKEDIN", "40", "Lahore")
		if err != nil {
			t.Errorf("%s", err.Error())
		}

	})
	t.Run("Out of range radius distance", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid distance has been passed! distance should be in between 25 and 100"
		_, err := MakeTask(3600, "Test search", "54633212", "LINKEDIN", "404", "Lahore")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
	t.Run("Invalid radius distance type", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid distance type! Not an integer"
		_, err := MakeTask(3600, "Test search", "54633212", "LINKEDIN", "a1s", "Lahore")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
	t.Run("Out bound delay for the scraper task run", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid delay! less than 30 minutes"
		_, err := MakeTask(1700, "Test search", "54633212", "LINKEDIN", "45", "Lahore")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})

	t.Run("Invaid Search Keywords empty string", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid seach keywords! Empty search"
		_, err := MakeTask(1900, "", "54633212", "LINKEDIN", "45", "Lahore")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
	t.Run("Invaid task location empty string", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid location! Empty task location"
		_, err := MakeTask(1900, "Software Engineer", "54633212", "LINKEDIN", "45", "")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
	t.Run("Invaid task location Id empty string", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid location id! Empty task location id"
		_, err := MakeTask(1900, "Software Engineer", "", "LINKEDIN", "45", "Kiel")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
	t.Run("Invaid task location Id! Contains non-digit characters", func(t *testing.T) {
		t.Parallel()
		expectedErrorMsg := "invalid location id! Location Id must be numeric"
		_, err := MakeTask(1900, "Software Engineer", "a5463aa3212", "LINKEDIN", "45", "Kiel")
		assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
	})
}
