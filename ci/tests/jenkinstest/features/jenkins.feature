Feature: jenkins
  In order have a valid jenkins installation
  As an administrator
  I need to be able to login to jenkins

  Scenario: Can access login screen
    Given there is a jenkins install
    When I access the login screen
    Then jenkins should be unlocked

  Scenario: Anonymous user cannot create jobs
    Given there is a jenkins install
	When an anonymous user tries to create a job
	Then the user is prompted to login
