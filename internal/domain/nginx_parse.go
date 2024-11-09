package domain

import (
	"regexp"
	"strconv"
	"time"
)

type NGINXParser struct{}

func (p *NGINXParser) Parse(logString string) (*NGINX, error) {
	regexPattern := regexp.MustCompile(`^(\S+) (\S+) (\S+) \[([^\]]+)\] "([^"]+)" (\d{3}) (\d+) "([^"]*)" "([^"]*)"$`)
	matches := regexPattern.FindStringSubmatch(logString)

	if len(matches) != 10 {
		return nil, &ParseNGINXStringError{Message: "incorrect logString format"}
	}

	if matches[2] != "-" {
		return nil, &ParseNGINXStringError{Message: "incorrect NGINX format"}
	}

	timeLocal, err := time.Parse("02/Jan/2006:15:04:05 +0000", matches[4])
	if err != nil {
		return nil, &ParseNGINXStringError{Message: "incorrect timeLocal format"}
	}

	status, err := strconv.Atoi(matches[6])
	if err != nil {
		return nil, &ParseNGINXStringError{Message: "incorrect status format"}
	}

	bodyBytesSent, err := strconv.Atoi(matches[7])
	if err != nil {
		return nil, &ParseNGINXStringError{Message: "incorrect bodyBytesSent format"}
	}

	return NewNGINX(
		matches[1],
		matches[3],
		timeLocal,
		matches[5],
		status,
		bodyBytesSent,
		matches[8],
		matches[9],
	), nil
}
