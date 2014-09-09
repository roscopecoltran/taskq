package taskq_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	taskqPkg "github.com/zyndiecate/taskq"
)

func SeriesQueue() (*Ctx, error) {
	ctx := &Ctx{}
	q := taskqPkg.NewQueue(ctx)

	err := q.RunTasks(
		taskqPkg.InSeries(
			Task1,
			Task3,
			Task4,
			Task2,
		),
	)

	if err != nil {
		return &Ctx{}, err
	}

	return ctx, nil
}

func TestSeriesQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "series-queue")
}

var _ = Describe("series-queue", func() {
	var (
		ctx *Ctx
		err error
	)

	BeforeEach(func() {
		ctx, err = SeriesQueue()
	})

	Context("executing mixed queue", func() {
		It("should not throw error", func() {
			Expect(err).To(BeNil())
		})

		It("should create correct context value for task1", func() {
			Expect(ctx.Task1).To(Equal("task1"))
		})

		It("should create correct context value for task2", func() {
			Expect(ctx.Task2).To(Equal(2))
		})

		It("should create correct context value for task3", func() {
			Expect(ctx.Task3).To(Equal([]string{"task3"}))
		})

		It("should create correct context value for task4", func() {
			Expect(ctx.Task4).To(Equal(4.4))
		})
	})
})
