package main

import (
	"bufio"
	"flag"
	"fmt"

	"io"
	"os"

	"strings"

	"code/gobro/todo"
)

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}

const todoFileName = ".todo.json"

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "sometihng"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myLists arrayFlags

func main() {
	add := flag.Bool("add", false, "Add task to the TOdo list")
	list := flag.Bool("list", false, "List all tasks")
	listwo := flag.Bool("listwo", false, "List tasks that are not completed")
	complete := flag.Int("complete", 0, "Item to be completed")
	fromFile := flag.String("file", "", "name of file to pull from")
	fromFiles := flag.Var(&myLists, "files", "name of multiple files")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *listwo:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *fromFile != "":
		taskList, err := parseFile(*fromFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		for _, task := range taskList {
			fmt.Println(task)
			l.Add(task)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		//	case *fromFiles != "":
		//		filesList := flag.Args()
		//		fmt.Println(filesList)

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func parseFile(fName string) ([]string, error) {
	file, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var taskList []string

	for scanner.Scan() {
		taskList = append(taskList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return taskList, nil
}
