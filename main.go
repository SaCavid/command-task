package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"
)

type License struct {
	EndDate string          `json:"end_date"` //time.Time
	NotificationDate string `json:"notification_date"`//time.Time
}

func (l *License) ValidateLicense () error{

	endDate, err := time.Parse(time.RFC3339, l.EndDate + "T00:00:00Z")

	if err != nil {
		return fmt.Errorf("error with license file: end date format is not correct")
	}

	notificationDate, err := time.Parse(time.RFC3339, l.NotificationDate + "T00:00:00Z")
	if err != nil {
		return fmt.Errorf("error with license file: notification date format is not correct")
	}

	n := time.Since(notificationDate).Seconds()

	if n > 0 {
		fmt.Println("Notification date reached")
	} else {
		fmt.Println("Notification date not reached")
	}

	s := time.Since(endDate).Seconds()

	if s > 0 {
		fmt.Println("License date expired in: ", endDate.String()[:10])
		return nil
	}

	fmt.Println("License date will be expire in: ", endDate.String()[:10])

	return nil
}

func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	if len(args) > 1 {
		return errors.New("not allowed multiple checking")
	}

	str := args[0]

	serviceNames := strings.Split(str, "/")

	if len(serviceNames) != 2 {
		return errors.New("not allowed command form")
	}

	// company name
	if !IsLetter(serviceNames[0]) {
		return errors.New("only unicode characters allowed in company name")
	}

	data, err := ioutil.ReadFile("./" + serviceNames[0] + "/" + serviceNames[1] + "/license.json")
	if err != nil {
		return errors.New("not registered service")
	}

	license := License{}

	err = json.Unmarshal(data, &license)
	if err != nil {
		return fmt.Errorf("error with unmarshal: %s", err.Error())
	}

	err = license.ValidateLicense()
	if err != nil {
		return err
	}

	return nil
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	if err := Root(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
