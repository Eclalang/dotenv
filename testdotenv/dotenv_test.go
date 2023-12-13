package dotenv

import (
	"github.com/Eclalang/dotenv"
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	key := "test"
	value := "value"
	dotenv.SetEnv(key, value)
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
	got := dotenv.GetEnv("test")
	if expected != got {
		t.Errorf("dotenv.GetEnv(%s) returned %s, expected %s", value, got, expected)
	}
}

func TestLoadFile(t *testing.T) {
	os.Setenv("Test", "Pates")
	dotenv.LoadFile("test.env")
	expected1 := "truc"
	expected2 := os.Getenv("Chose")
	got1 := dotenv.GetEnv("Test")
	got2 := dotenv.GetEnv("Chose")
	if expected1 != got1 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected1)
	}
	if expected2 != got2 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected2)
	}
}

func TestOverloadFile(t *testing.T) {
	os.Setenv("Test", "Pates")
	os.Setenv("Chose", "Pate")
	dotenv.OverloadFile("test.env")
	expected1 := "truc"
	expected2 := "pates"
	got1 := dotenv.GetEnv("Test")
	got2 := dotenv.GetEnv("Chose")
	if expected1 != got1 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got1, expected1)
	}
	if expected2 != got2 {
		t.Errorf("dotenv.GetEnv() returned %s, expected %s", got2, expected2)
	}
}
