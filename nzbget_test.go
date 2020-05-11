package nzbget_test

import (
	"mylar"
	"nzbget"
	"testing"

	"github.com/SemanticallyNull/golandreporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
)

const (
	nzbgetURL = "http://localhost:8090"
	apiKey    = "testapikey"
)

func TestMylar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Mylar Suite", []Reporter{
		golandreporter.NewGolandReporter(),
	})
}

var (

)

var _ = Describe("NZBGet", func() {

	Context("#GetStatus", func() {

		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/api").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(mockIndex)
			})

			It("should return two items", func() {

			})

		})
	})
})
