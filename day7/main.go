package main

import (
	"strings"
	"github.com/cfagiani/aoc2018/util"
	"sort"
	"fmt"
)

type Task struct {
	Id           string
	Dependencies []string
}

type Assignment struct {
	TaskId   string
	Deadline int
}

func main() {
	inputString := util.ReadFileAsString("input/day7.input")
	lines := strings.Split(inputString, "\n")
	tasks := buildDependencies(lines)
	part1(tasks)
	part2(tasks)

}
func part1(tasks map[string]Task) {
	var tasksRun []string
	for len(tasksRun) < len(tasks) {
		canRun := getEligibleTasks(tasks, tasksRun)
		sort.Strings(canRun)
		tasksRun = append(tasksRun, canRun[0])
	}
	fmt.Println("Tasks run:\n")
	for _, v := range tasksRun {
		fmt.Printf("%s", v)
	}
}

func part2(tasks map[string]Task) {

	minStepDuration := 60
	workerCount := 5
	time := 0
	workers := make(map[int]Assignment)
	for i := 0; i < workerCount; i++ {
		workers[i] = Assignment{TaskId: "", Deadline: 0}
	}
	var tasksRun []string
	var tasksComplete []string
	for len(tasksComplete) < len(tasks) {
		tasksComplete, workers = updateCompletion(time, workers, tasksComplete)
		canRun := getEligibleTasks(tasks, tasksComplete)
		canRun = util.FilterArray(canRun, tasksRun)
		sort.Strings(canRun)
		for id, a := range workers {
			if a.TaskId == "" && len(canRun) > 0 {
				a.TaskId = canRun[0]
				a.Deadline = computeDeadline(time, canRun[0], minStepDuration)
				tasksRun = append(tasksRun, canRun[0])
				canRun = canRun[1:]
				workers[id] = a
			}
		}
		time++
	}
	fmt.Printf("\nTotal time to complete: %d\n", time-1)
}

func computeDeadline(time int, taskId string, minDuration int) int {
	return time + minDuration + int(taskId[0]) - 64
}

func updateCompletion(time int, workers map[int]Assignment, tasksComplete []string) ([]string, map[int]Assignment) {
	for id, a := range workers {
		if a.TaskId != "" {
			if time >= a.Deadline {
				tasksComplete = append(tasksComplete, a.TaskId)
				a.TaskId = ""
				a.Deadline = 0
				workers[id] = a
			}
		}
	}
	return tasksComplete, workers
}

func getEligibleTasks(tasks map[string]Task, tasksRun []string) []string {
	var canRun []string
	for _, task := range tasks {
		if !util.IsStringInSlice(task.Id, tasksRun) {
			allDepsMet := true
			for _, d := range task.Dependencies {
				if !util.IsStringInSlice(d, tasksRun) {
					allDepsMet = false
					break
				}
			}
			if allDepsMet {
				canRun = append(canRun, task.Id)
			}
		}
	}
	return canRun
}

func buildDependencies(lines []string) map[string]Task {
	tasks := make(map[string]Task)
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		taskId := parts[7]
		dependency := parts[1]
		t, ok := tasks[taskId]
		if ok {
			t.Dependencies = append(t.Dependencies, dependency)
		} else {
			t = Task{Id: taskId, Dependencies: []string{dependency}}
		}
		tasks[taskId] = t
		//also ensure the dependent task is created (only way to ensure tasks with no deps of their own show up)
		t, ok = tasks[dependency]
		if !ok {
			t = Task{Id: dependency, Dependencies: []string{}}
		}
		tasks[dependency] = t
	}
	return tasks
}
