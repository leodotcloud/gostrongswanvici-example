package main

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/bronze1man/goStrongswanVici"
)

func TestInitiate(t *testing.T) {
	logrus.Info("testing initiate")

	c, err := goStrongswanVici.NewClientConnFromDefaultSocket()
	if err != nil {
		logrus.Errorf("error: %v", err)
		t.Fail()
	}
	err = c.Initiate("child-192.168.236.202", "")
	if err != nil {
		logrus.Errorf("error: %v", err)
		t.Fail()
	}
}
