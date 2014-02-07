package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoblog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goblog Suite")
}
