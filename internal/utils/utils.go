package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadStdin() ([]string, error) {
    var data []string
    var err error

    file, err := os.Stdin.Stat()

    if err != nil {
        return data, err
    }


    if !(file.Mode() & os.ModeCharDevice == 0) { return data, nil }

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        data = append(data, scanner.Text())
    }


    return data, nil
}

func ReadFile(fp string) (string, error){

        var text string

        bytes, err := os.ReadFile(fp)

        if err != nil {
            return text, err
        }

        text = string(bytes)


    return text, nil
}

func ArgsToIds(args []string) ([]int64, error) {
    var ids []int64
        for _, arg := range args {
            arg = strings.Trim(arg, "\n\t")
            if arg == "" { continue }
            id, err := strconv.Atoi(arg)

            if err != nil {
                return ids, err
            }

            ids = append(ids, int64(id))
        }
    return ids, nil
}

