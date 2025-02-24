package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

// main struct to store task data
type task struct {
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

// node for the linked list
type Node struct {
	data *task
	next *Node
	prev *Node
}

// doubly linked list to store tasks
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

var footer, header string
var maxDescLen int

// main program loop
func main() {
	var tasksList DoublyLinkedList

	loadTasks(&tasksList)

	if err := keyboard.Open(); err != nil {
		fmt.Println("error opening keyboard:", err)
		return
	}
	defer keyboard.Close()

	for {
		maxDescLen = tasksList.getMaxDescLength()

		header = generateHeader(maxDescLen)
		footer = generateFooter(maxDescLen)

		clearScreen()

		fmt.Print(header)
		tasksList.printList(maxDescLen)
		fmt.Print(footer)

		fmt.Println("\n'a' to add / 'r' to remove / '0' to exit ")

		char, _, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("error reading key:", err)
			return
		}

		if char == 'a' {
			handleAddMode(&tasksList)
		} else if char == 'r' {
			handleRemoveMode(&tasksList)
		} else if char == '0' {
			clearScreen()
			return
		} else if char >= '1' {
			index := int(char - '0')
			changeStatus(&tasksList, index)
		}
	}
}

// add task mode
func handleAddMode(list *DoublyLinkedList) {
	keyboard.Close()

	for {
		clearScreen()
		maxDescLen = list.getMaxDescLength()
		header = generateHeader(maxDescLen)
		footer = generateFooter(maxDescLen)

		fmt.Print(header)
		list.printList(maxDescLen)
		fmt.Print(footer + "\n")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("add > ")

		desc, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			return
		}

		desc = strings.TrimRight(desc, "\r\n")

		if desc == "0" {
			if err := keyboard.Open(); err != nil {
				fmt.Println("error reopening keyboard:", err)
				os.Exit(1)
			}
			return
		} else if desc != "" {
			newTask := &task{Desc: desc}
			list.addAtEnd(newTask)
			saveTasks(list)
		}
	}
}

// remove task mode
func handleRemoveMode(list *DoublyLinkedList) {
	for {
		clearScreen()

		fmt.Print(header)
		list.printList(maxDescLen)
		fmt.Print(footer)

		fmt.Println("\n press the index task to remove / '0' to exit")

		char, _, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("error reading key:", err)
			return
		}

		if char == '0' {
			return
		} else if char >= '1' {
			index := int(char - '0')
			removeTask(list, index)
		}
	}
}

// removes a task by index
func removeTask(list *DoublyLinkedList, index int) {
	currentNode := list.head
	for i := 1; currentNode != nil; i++ {
		if i == index {
			if currentNode.prev != nil {
				currentNode.prev.next = currentNode.next
			} else {
				list.head = currentNode.next
			}

			if currentNode.next != nil {
				currentNode.next.prev = currentNode.prev
			} else {
				list.tail = currentNode.prev
			}

			fmt.Printf("Task '%s' removed\n", currentNode.data.Desc)
			break
		}
		currentNode = currentNode.next
	}
	saveTasks(list)
}

// toggles task status between complete/incomplete
func changeStatus(list *DoublyLinkedList, index int) {
	currentNode := list.head
	for i := 1; currentNode != nil; i++ {
		if i == index {
			currentNode.data.Completed = !currentNode.data.Completed
			break
		}
		currentNode = currentNode.next
	}
	saveTasks(list)
}

// adds a new task at the end of the list
func (list *DoublyLinkedList) addAtEnd(task *task) {
	newNode := &Node{data: task}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

// displays all tasks on screen
func (list *DoublyLinkedList) printList(maxDescLen int) {
	currentNode := list.head
	index := 1

	minWidth := len("cli to-do list in go :)")
	totalWidth := maxDescLen + 9
	if totalWidth < minWidth {
		totalWidth = minWidth
	}

	for currentNode != nil {
		desc := currentNode.data.Desc

		if utf8.RuneCountInString(desc) > maxDescLen {
			desc = truncateString(desc, maxDescLen)
		}

		status := "   "
		if currentNode.data.Completed {
			status = " X "
		}

		padding := totalWidth - utf8.RuneCountInString(desc) - 8
		line := fmt.Sprintf("| %d. %s%s|%s|",
			index,
			desc,
			strings.Repeat(" ", padding),
			status)

		fmt.Println(line)
		currentNode = currentNode.next
		index++
	}
}

// cuts text if it's too long
func truncateString(s string, maxLen int) string {
	var result []rune
	for len(result) < maxLen {
		r, size := utf8.DecodeRuneInString(s)
		if size == 0 {
			break
		}
		result = append(result, r)
		s = s[size:]
	}
	return string(result)
}

// gets the length of the longest description
func (list *DoublyLinkedList) getMaxDescLength() int {
	maxLen := 0
	currentNode := list.head
	for currentNode != nil {
		length := utf8.RuneCountInString(currentNode.data.Desc)
		if length > maxLen {
			maxLen = length
		}
		currentNode = currentNode.next
	}
	return maxLen
}

// creates the interface footer
func generateFooter(maxDescLen int) string {
	minWidth := len("cli to-do list in go :)")
	totalWidth := maxDescLen + 9
	if totalWidth < minWidth {
		totalWidth = minWidth
	}

	return fmt.Sprintf("+%s+", strings.Repeat("-", totalWidth))
}

// creates the interface header
func generateHeader(maxDescLen int) string {
	title := "cli to-do list in go :)"
	minWidth := len(title)
	totalWidth := maxDescLen + 9
	if totalWidth < minWidth {
		totalWidth = minWidth
	}

	border := "+" + strings.Repeat("-", totalWidth) + "+\n"

	padding := totalWidth - len(title)
	leftPadding := padding / 2
	rightPadding := padding - leftPadding

	return border +
		fmt.Sprintf("|%s%s%s|\n",
			strings.Repeat(" ", leftPadding),
			title,
			strings.Repeat(" ", rightPadding)) +
		border
}

// clears terminal screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// saves tasks to json file so they're not lost when closing
func saveTasks(list *DoublyLinkedList) {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	var tasks []task
	current := list.head
	for current != nil {
		tasks = append(tasks, *current.data)
		current = current.next
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("error saving tasks:", err)
	}
}

// loads tasks from json file when program starts
func loadTasks(list *DoublyLinkedList) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return
	}
	defer file.Close()

	var tasks []task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		fmt.Println("error loading tasks:", err)
		return
	}

	for _, t := range tasks {
		taskCopy := t
		list.addAtEnd(&taskCopy)
	}
}
