package file

import "os"

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(path string, name string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0766)
	}
	fo, err := os.Create(path + "/" + name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
