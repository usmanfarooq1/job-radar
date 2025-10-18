package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeManager(t *testing.T) {
	t.Run("Creating a valid empty manager", func(t *testing.T) {
		manager := MakeManager()
		assert.Equal(t, 0, manager.GetManagerTasksCount())
	})
	manager := MakeManager()
	task, err := MakeTask(3600, "Test search", "54633212", "LINKEDIN", "40", "Lahore")
	if err != nil {
		t.Fatalf("Unable to create a task: %s", err.Error())
	}
	task.Execute()

	t.Run("Adds a task to the manager", func(t *testing.T) {
		manager.AddScraperTask(*task)
		assert.Equal(t, 1, manager.GetManagerTasksCount())
	})
	t.Run("Update a task in the manager", func(t *testing.T) {
		manager.UpdateScraperTask(task.id, 3600, "Updated Search", "54633212", "45", "Hamburg")
		assert.Equal(t, true, task.Equal(*task))
	})
	t.Run("Stop Execution of a task to the manager", func(t *testing.T) {
		err := manager.StopScraperTask(task.id)
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		stoppedTask := manager.getScraperTask(task.id)
		if stoppedTask == nil {
			t.Fatalf("No Task found")
		}

		assert.Equal(t, false, stoppedTask.isRunning)
	})

	t.Run("Remove a task to the manager", func(t *testing.T) {
		err := manager.RemoveScraperTask(task.id)
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		assert.Equal(t, 0, manager.GetManagerTasksCount())
	})
}
