package lineparser

import (
	"bufio"
	"net/http"
	"os"
)

type LineParser[T any] interface {
	Parse(line string) (*T, error)
}

func ReadFromFile[T any](paths []string, parser LineParser[T]) ([]*T, error) {
	var data []*T

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

				data = append(data, parsedData)
			}

			if err := scanner.Err(); err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			var zeroValue []*T
			return zeroValue, err
		}
	}

	return data, nil
}

func ReadFromURL[T any](url string, parser LineParser[T]) ([]*T, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []*T

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		parsedData, err := parser.Parse(line)
		if err != nil {
			continue
		}

		data = append(data, parsedData)
	}

	if err := scanner.Err(); err != nil {
		var zeroValue []*T
		return zeroValue, err
	}

	return data, nil
}
