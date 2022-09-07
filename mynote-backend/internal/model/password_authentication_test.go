package model_test

import (
	"MyNote/internal/model"
	"MyNote/pkg/database/database_test"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PasswordAuthentication", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Context("When password correct", func() {
		var userId uint

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId = user.ID
			model.CreatePasswordAuthentication("password", userId)
		})

		It("return true.", func() {
			Expect(model.CorrectPassword("password", int(userId))).To(BeTrue())
		})
	})

	Context("When password incorrect", func() {
		var userId uint

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId = user.ID
			model.CreatePasswordAuthentication("password", userId)
		})

		It("return false.", func() {
			Expect(model.CorrectPassword("IncorrectPassword", int(userId))).To(BeFalse())
		})
	})
})
