package role_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRole(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Role Suite")
}
