package registry

import (
	"errors"
	"fmt"
	"os"
)

type Registry struct {
	file *os.File
	databases []string
}

func InitRegistry(name, dir string) (*Registry, error) {
	location := fmt.Sprintf("%s/%s", dir, name)
	file , err := os.OpenFile(location, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0644);

	if err != nil {
		return nil, err;
	}

	return  &Registry{
		file: file,
		databases: make([]string, 0),
	}, nil
}

func (r *Registry) AddEntry(database string) error {
	entry := fmt.Sprintf("%s: ./%s",database, database);

	if _, err := r.file.WriteString(entry); err != nil {
		errorString := fmt.Sprintf("Failed to add entry to registry: %s", err.Error())
		return errors.New(errorString)
	}
	
	return nil
}

func (r *Registry) Cleanup() {
	if err := r.file.Close(); err != nil { 
		panic("Could not close registry file")
	}
}