package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("07/test.txt")

	sizes := getFileSystem(b)

	var disksize, required int64 = 70000000, 30000000
	var used int64
	for _, v := range sizes {
		used += v
	}

	log.Println("Part 1 answer:", getTotalBigDirs(sizes, 100000))
	log.Println("Part 2 answer:", getSmallestDirBiggerThan(sizes, required-(disksize-used)))
}

type CommandResponse struct {
	Command     string
	CommandLine string
	Response    string
}

func getFileSystem(b []byte) map[string]int64 {
	pwd := "/"
	commands := getCommands(b)

	sizes := make(map[string]int64)

	for _, v := range commands {
		switch v.Command {
		case "cd":
			pwd = changeDirectory(pwd, v.CommandLine)
		case "ls":
			sizes[pwd] = getSize(v.Response)
		default:
			panic(fmt.Sprintf("unrecognised command: %s", v.Command))
		}
	}

	return sizes
}

func getTotalBigDirs(sizes map[string]int64, bigDirSize int64) int64 {
	var total int64
	for dir := range sizes {
		totalDirSize := getTotalDirSize(sizes, dir)

		if totalDirSize < bigDirSize {
			total += totalDirSize
		}
	}

	return total
}

func getSmallestDirBiggerThan(sizes map[string]int64, size int64) int64 {
	var candidates []int64
	for dir := range sizes {
		totalDirSize := getTotalDirSize(sizes, dir)
		if totalDirSize > size {
			candidates = append(candidates, totalDirSize)
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	return candidates[0]
}

func getTotalDirSize(sizes map[string]int64, parent string) int64 {
	var total int64
	for dir, size := range sizes {
		if strings.HasPrefix(dir, parent) {
			total += size
		}
	}

	return total
}

func getSize(response string) int64 {
	var total int64
	lines := strings.FieldsFunc(response, func(r rune) bool {
		return r == '\n'
	})

	for _, v := range lines {
		chunks := strings.Fields(v)
		if chunks[0] == "dir" {
			continue
		}

		size, err := strconv.ParseInt(chunks[0], 10, 64)
		if err != nil {
			panic(err)
		}

		total += size
	}

	return total
}

func getCommands(b []byte) []CommandResponse {
	var c []CommandResponse

	chunks := strings.Split(string(b), `$`)
	for _, v := range chunks {
		if strings.TrimSpace(v) == "" {
			continue
		}

		cmdLineAndResponse := strings.SplitN(v, "\n", 2)
		cmd := strings.Fields(cmdLineAndResponse[0])

		c = append(c, CommandResponse{
			Command:     cmd[0],
			CommandLine: cmdLineAndResponse[0],
			Response:    cmdLineAndResponse[1],
		})
	}

	return c
}

func changeDirectory(pwd string, cmd string) string {
	parts := strings.Fields(cmd)
	if len(parts) != 2 {
		panic("bad cd command; expected one argument")
	}

	// absolute path
	if strings.HasPrefix(parts[1], "/") {
		return parts[1]
	}

	dirs := strings.FieldsFunc(pwd[1:], func(r rune) bool {
		return r == '/'
	})

	// move up relatively
	if parts[1] == ".." {
		dirs = dirs[:len(dirs)-1]
		return "/" + strings.Join(dirs, "/")
	}

	dirs = append(dirs, parts[1])

	return "/" + strings.Join(dirs, "/")
}
