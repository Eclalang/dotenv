package dotenv

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	key := "test"
	value := "value"
	SetEnv(key, value)
	expected := os.Getenv("test")
	got := os.Getenv("test")
	if expected != got {
		t.Errorf("dotenv.SetEnv(%s, %s) returned %s, expected %s", key, value, got, expected)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("test", "value")
	value := "test"
	expected := os.Getenv("test")
	got := GetEnv("test")
	if expected != got {
		t.Errorf("dotenv.GetEnv(%s) returned %s, expected %s", value, got, expected)
	}
}

func TestLoadFile(t *testing.T) {
	err := os.Setenv("Test", "othervalue")
	if err != nil {
		return
	}
	err = LoadFile("/testdotenv/test.env")
	if err != nil {
		return
	}
	expected1 := "othervalue"
	expected2 := GetEnv("Test1")
	got1 := GetEnv("Test")
	got2 := GetEnv("Test1")
	if expected1 != got1 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected1)
	}
	if expected2 != got2 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected2)
	}
}

func TestOverloadFile(t *testing.T) {
	err := os.Setenv("Test", "othervalue")
	if err != nil {
		return
	}
	err = os.Setenv("Test1", "othervalue1")
	if err != nil {
		return
	}
	err = OverloadFile("/testdotenv/test.env")
	if err != nil {
		return
	}
	expected1 := "testvalue"
	expected2 := "testvalue1"
	got1 := GetEnv("Test")
	got2 := GetEnv("Test1")
	if expected1 != got1 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected1)
	}
	if expected2 != got2 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got2, expected2)
	}
}
