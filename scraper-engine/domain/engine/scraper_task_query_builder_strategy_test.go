package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQueryBuilderStrategy(t *testing.T) {
	t.Run("Valid LinkedIn query builder", func(t *testing.T) {
		builder, err := GenerateQueryBuilderStrategy(LinkedIn)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		task, err := MakeTask(3600, "Test search", "54633212", "LINKEDIN", "40", "Lahore")
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		query, err := builder.Construct(task)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		assert.Equal(t, query, "https://www.linkedin.com/jobs/search?keywords=Test%20search&location=Lahore&geoId=54633212&distance=40&f_TPR=r3600")
	})
}
