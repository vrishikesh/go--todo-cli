package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (l *List) Get(fileName string) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, l)
}
