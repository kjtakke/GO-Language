package utils

import (
    "bufio"
    "errors"
    "io/fs"
    "os"
    "path/filepath"
)

// ReadFile reads the contents of a file and returns it as a string.
func ReadFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// ReadLines reads a file line-by-line and returns a slice of strings.
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil
}

// WriteFile writes the given content to a file, overwriting if it exists.
func WriteFile(path string, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

// AppendToFile appends content to the end of a file or creates it if it doesn't exist.
func AppendToFile(path string, content string) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(content)
    return err
}

// FileExists checks if a file or directory exists at the given path.
func FileExists(path string) bool {
    _, err := os.Stat(path)
    return !errors.Is(err, os.ErrNotExist)
}

// ListFiles returns a list of file names (not directories) in the given folder.
func ListFiles(dir string) ([]string, error) {
    var files []string

    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    for _, entry := range entries {
        if !entry.IsDir() {
            files = append(files, entry.Name())
        }
    }

    return files, nil
}

// FindFilesByExtension recursively finds files with a given extension.
func FindFilesByExtension(root, ext string) ([]string, error) {
    var matched []string

    err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && filepath.Ext(path) == ext {
            matched = append(matched, path)
        }
        return nil
    })

    return matched, err
}
