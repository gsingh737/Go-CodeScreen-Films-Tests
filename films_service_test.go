package films

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFilms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Films Service Test Suite")
}

var _ = Describe("Films API", func() {

	Context("Films API Service Tests", func() {

		It("Best rated film for director Ridley Scott", func() {
			Expect(BestRatedFilm("Ridley Scott")).Should(Equal("Gladiator"))
		})

		It("Director with the most films", func() {
			Expect(DirectorWithMostFilms()).Should(Equal("Ridley Scott"))
		})

		It("Average rating for films directed by Ridley Scott", func() {
			Expect(AverageRating("Ridley Scott")).Should(Equal(6.9))
		})

		It("Shortest number of days between film releases by Ridley Scott", func() {
			Expect(ShortestFilmReleaseGap("Ridley Scott")).Should(Equal(29))
		})

		It("Best rated film for director Christopher Nolan", func() {
			Expect(BestRatedFilm("Christopher Nolan")).Should(Equal("The Dark Knight"))
		})

		It("Average rating for films directed by Christopher Nolan", func() {
			Expect(AverageRating("Christopher Nolan")).Should(Equal(8.3))
		})

		It("Shortest number of days between film releases by Christopher Nolan", func() {
			Expect(ShortestFilmReleaseGap("Christopher Nolan")).Should(Equal(147))
		})

		It("Best rated film for unknown director", func() {
			Expect(BestRatedFilm("Unknown Director")).Should(Equal(""))
		})

		It("Average rating for films directed by unknown director", func() {
			Expect(AverageRating("Unknown Director")).Should(Equal(0.0))
		})

		It("Shortest number of days between film releases by unknown director", func() {
			Expect(ShortestFilmReleaseGap("Unknown Director")).Should(Equal(0))
		})

	})

})
