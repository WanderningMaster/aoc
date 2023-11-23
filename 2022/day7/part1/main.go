package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type NodeType string

const (
	Dir  NodeType = "DIR"
	File NodeType = "File"
)

type TreeNode struct {
	Name     string
	Size     int
	Type     NodeType
	children []*TreeNode
}

func createRoot() *TreeNode {
	return &TreeNode{
		Name:     "/",
		Size:     0,
		Type:     Dir,
		children: []*TreeNode{},
	}
}

func createDir(name string) *TreeNode {
	return &TreeNode{
		Name:     name,
		Size:     0,
		Type:     Dir,
		children: []*TreeNode{},
	}
}

func createFile(name string, size int) *TreeNode {
	return &TreeNode{
		Name:     name,
		Size:     size,
		Type:     File,
		children: []*TreeNode{},
	}
}

func (t *TreeNode) AddChild(child *TreeNode) {
	t.children = append(t.children, child)
}

func (t *TreeNode) searchDFS(name string) *TreeNode {
	// if t.Name == name {
	// 	return t
	// }
	for _, node := range t.children {
		if found := node.searchDFS(name); found != nil {
			return found
		}
	}
	return nil
}

func dirTotalSize(node *TreeNode) int {
	total := 0
	for _, subnode := range node.children {
		if subnode.Type == File {
			total += subnode.Size
		} else {
			total += dirTotalSize(subnode)
		}
	}

	return total
}

func (t *TreeNode) Print(depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "\t"
	}
	nodeDescr := ""
	if t.Type == Dir {
		nodeDescr = "(dir)"
	} else {
		nodeDescr = fmt.Sprintf("(file, size=%d)", t.Size)
	}
	fmt.Printf("%s- %s %s\n", indent, t.Name, nodeDescr)

	for _, child := range t.children {
		child.Print(depth + 1)
	}
}

func splitCommands(input string) [][]string {
	commands := [][]string{}

	lines := strings.Split(input, "\n")
	idx := 0
	commands = append(commands, []string{lines[0]})
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "$") {
			idx += 1
			commands = append(commands, []string{})
		}
		commands[idx] = append(commands[idx], line)
	}

	return commands
}

func cdCommand(comm []string, stackTrace *utils.Stack) {
	splitted := strings.Split(comm[0], " ")
	dirname := splitted[2]

	currDir, _ := stackTrace.Peek()
	if dirname == ".." {
		_, err := stackTrace.Pop()
		if err != nil {
			log.Fatal(err)
		}
	} else if dirname == "/" {
	} else {
		node := currDir.(*TreeNode).searchDFS(dirname)
		if node == nil {
			node = createDir(dirname)
			currDir.(*TreeNode).AddChild(node)
		}

		stackTrace.Push(node)
	}
}

func lsCommand(comm []string, stackTrace *utils.Stack, root *TreeNode) {
	currDir, err := stackTrace.Peek()
	if err != nil {
		log.Fatal(err)
	}
	for _, el := range comm[1:] {
		if strings.HasPrefix(el, "dir") {
			s := strings.Split(el, " ")
			node := createDir(s[1])
			currDir.(*TreeNode).AddChild(node)
		} else if len(el) != 0 {
			s := strings.Split(el, " ")
			size, _ := strconv.Atoi(s[0])
			fileName := s[1]

			fileNode := createFile(fileName, size)

			currDir.(*TreeNode).AddChild(fileNode)
		}
	}
}

func processCommands(root *TreeNode, commands [][]string) {
	stackTrace := utils.NewStack()
	stackTrace.Push(root)
	for _, command := range commands {
		if strings.HasPrefix(command[0], "$ ls") {
			lsCommand(command, stackTrace, root)
		}
		if strings.HasPrefix(command[0], "$ cd") {
			cdCommand(command, stackTrace)
		}
	}
}

func sizeYouCanDelete(node *TreeNode) int {
	AllowedMaxSize := 100000
	total := 0

	if node.Type == Dir {
		dirSize := dirTotalSize(node)
		if dirSize <= AllowedMaxSize {
			total += dirSize
		}
	}
	for _, node := range node.children {
		total += sizeYouCanDelete(node)
	}

	return total
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day7/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	commands := splitCommands(string(input))
	root := createRoot()
	processCommands(root, commands)

	root.Print(0)
	fmt.Println("Result:", sizeYouCanDelete(root))
}
