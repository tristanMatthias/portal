package lib

import "os"

func ListDirs(path string) ([]string, error) {
    var dirs []string

    // Open the directory
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    // Read the directory entries
    entries, err := f.Readdir(-1)
    if err != nil {
        return nil, err
    }

    // Extract the directory names
    for _, entry := range entries {
        if entry.IsDir() {
            dirs = append(dirs, entry.Name())
        }
    }

    return dirs, nil
}
