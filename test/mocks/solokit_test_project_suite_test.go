package mocks

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolokitTestProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolokitTestProject Suite")
}





