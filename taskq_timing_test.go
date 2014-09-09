package taskq_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	taskqPkg "github.com/zyndiecate/taskq"
)

func TrackTiming(cb func()) int64 {
	start := time.Now().UnixNano()
	cb()
	end := time.Now().UnixNano()

	return end - start
}

func SeriesTimingQueue() (*Ctx, error) {
	ctx := &Ctx{}
	q := taskqPkg.NewQueue(ctx)

	err := q.RunTasks(
		taskqPkg.InSeries(
			Task1Sleep,
			Task3Sleep,
			Task4Sleep,
			Task2Sleep,
		),
	)

	if err != nil {
		return &Ctx{}, err
	}

	return ctx, nil
}

func ParallelTimingQueue() (*Ctx, error) {
	ctx := &Ctx{}
	q := taskqPkg.NewQueue(ctx)

	err := q.RunTasks(
		taskqPkg.InParallel(
			Task1Sleep,
			Task3Sleep,
			Task4Sleep,
			Task2Sleep,
		),
	)

	if err != nil {
		return &Ctx{}, err
	}

	return ctx, nil
}

func TestTimingQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "timing-queue")
}

var _ = Describe("timing-queue", func() {
	var (
		seriesDuration   int64
		parallelDuration int64
	)

	BeforeEach(func() {
		seriesDuration = TrackTiming(func() {
			SeriesTimingQueue()
		})

		parallelDuration = TrackTiming(func() {
			ParallelTimingQueue()
		})
	})

	Context("exectuting parallel queue", func() {
		It("should be faster than executing series queue", func() {
			Expect(parallelDuration).To(BeNumerically("<", seriesDuration))
		})
	})
})
