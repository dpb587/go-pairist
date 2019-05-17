package denormalized_test

import (
	. "github.com/dpb587/go-pairist/denormalized"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lanes", func() {
	var subject Lanes

	BeforeEach(func() {
		subject = Lanes{
			{
				ID: "id1",
				People: []Entity{
					{
						Name: "people1",
					},
				},
				Roles: []Entity{
					{
						Name: "role1",
					},
					{
						Name: "role2",
					},
				},
				Tracks: []Entity{
					{
						Name: "track1",
					},
					{
						Name: "track2",
					},
				},
			},
			{
				ID: "id2",
				People: []Entity{
					{
						Name: "people2",
					},
				},
				Roles: []Entity{
					{
						Name: "role2",
					},
					{
						Name: "role3",
					},
				},
				Tracks: []Entity{
					{
						Name: "track2",
					},
					{
						Name: "track3",
					},
				},
			},
			{
				ID: "id3",
			},
		}
	})

	Describe("ByRole", func() {
		It("filters", func() {
			roles := subject.ByRole("role3")
			Expect(roles).To(HaveLen(1))
			Expect(roles[0].ID).To(Equal("id2"))
		})

		It("filters", func() {
			roles := subject.ByRole("role2")
			Expect(roles).To(HaveLen(2))
			Expect([]string{roles[0].ID, roles[1].ID}).To(ConsistOf("id1", "id2"))
		})
	})
})
