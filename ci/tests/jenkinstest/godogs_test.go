package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var url string
var body string

func page(page string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body_bytes), nil
}

func thereIsAJenkinsInstall() error {
	url = os.Getenv("JENKINS_HOST")
	return nil
}

func iAccessTheLoginScreen() error {
	page_body, err := page("/login")
	if err != nil {
		return err
	}
	body = page_body
	return nil
}

func jenkinsShouldBeUnlocked() error {
	if strings.Contains(body, "Unlock Jenkins") {
		return fmt.Errorf("expected %s not to contain 'Unlock Jenkins'", body)
	}
	return nil
}

func anAnonymousUserTriesToCreateAJob() error {
	page_body, err := page("/newJob")
	if err != nil {
		return err
	}
	body = page_body
	return nil
}

func theUserIsPromptedToLogin() error {
	if !strings.Contains(body, "<a href=\"/login?from=/\">Log in</a> to create new jobs.") {
		return fmt.Errorf("expected %s to contain '<a href=\"/login?from=/\">Log in</a> to create new jobs.'", body)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a jenkins install$`, thereIsAJenkinsInstall)
	s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^jenkins should be unlocked$`, jenkinsShouldBeUnlocked)
	s.Step(`^an anonymous user tries to create a job$`, anAnonymousUserTriesToCreateAJob)
	s.Step(`^the user is prompted to login$`, theUserIsPromptedToLogin)
}
