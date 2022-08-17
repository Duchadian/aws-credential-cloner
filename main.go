package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
)

type AwsCredentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
}

var (
	awsConfigFile      = fmt.Sprintf("%s/.aws/config", os.Getenv("HOME"))
	awsCredentialsFile = fmt.Sprintf("%s/.aws/credentials", os.Getenv("HOME"))
)

func main() {
	cfg, err := ini.Load(awsConfigFile)
	if err != nil {
		log.Fatalf("Error! %s", err)
	}
	credsFile, err := ini.Load(awsCredentialsFile)
	if err != nil {
		log.Fatalf("Error! %s", err)
	}

	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		log.Fatal("No profile set as default profile, exiting.")
	}

	profileSection := fmt.Sprintf("profile %s", profile)
	section, err := cfg.GetSection(profileSection)

	if err != nil {
		log.Fatalf("Section %s cannot be found in aws config, exiting.", profile)
	}

	proc, err := section.GetKey("credential_process")
	if err != nil {
		log.Fatalf("Section %s has no credential_process to run, exiting.", profile)
	}

	log.Printf("Attempting to get credentials for profile %s", profile)
	cmd := fmt.Sprintf("%s %s", proc.Value(), "--no-session-cache")
	command := exec.Command("sh", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	var creds AwsCredentials
	err = json.Unmarshal(output, &creds)
	if err != nil {
		log.Fatal("Cannot unmarshal json response, exiting")
	}

	credsFile.Section(profile).Key("aws_access_key_id").SetValue(creds.AccessKeyId)
	credsFile.Section(profile).Key("aws_secret_access_key").SetValue(creds.SecretAccessKey)
	credsFile.Section(profile).Key("aws_session_token").SetValue(creds.SessionToken)

	err = credsFile.SaveTo(awsCredentialsFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully wrote credentials to %s", awsCredentialsFile)
}
