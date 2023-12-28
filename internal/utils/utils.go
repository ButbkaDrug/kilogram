package utils

import(
    "os"
    "bufio"
    "strings"
    "strconv"
)

func ReadStdin() string {

        var text string

        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {

            text += scanner.Text()
            text += "\n"

        }

        return text
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

