package datastream

import (
	"bufio"
	"net/http"
	"os"
)

type Parser[T any] interface {
	Parse(line string) (*T, error)
}

type Updater[T any] interface {
	Update(data *T)
}

func ProcessFromFile[T any](paths []string, parser Parser[T], updater Updater[T]) error {
	for _, path := range paths {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				parsedData, err := parser.Parse(line)
				if err != nil {
					continue
				}

				updater.Update(parsedData)
			}

			if err := scanner.Err(); err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessFromURL[T any](url string, parser Parser[T], updater Updater[T]) error {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		parsedData, err := parser.Parse(line)
		if err != nil {
			continue
		}

		updater.Update(parsedData)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
