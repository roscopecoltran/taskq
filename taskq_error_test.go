package taskq_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	taskqPkg "github.com/zyndiecate/taskq"
)

func TaskErr1(ctx interface{}) error {
	return fmt.Errorf("TaskErr1")
}

func TaskErr2(ctx interface{}) error {
	return fmt.Errorf("TaskErr2")
}

func ErrorQueue() (*Ctx, error) {
	ctx := &Ctx{}
	q := taskqPkg.NewQueue(ctx)

	err := q.RunTasks(
		taskqPkg.InSeries(
			Task1,
			TaskErr1,
			TaskErr2,
			Task4,
			Task2,
		),
	)

	if err != nil {
		// Because we want to receive the modified context even in error cases, we
		// do not return the contexts zero value here.
		return ctx, err
	}

	return ctx, nil
}

func TestErrorQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "error-queue")
}

var _ = Describe("error-queue", func() {
	var (
		ctx *Ctx
		err error
	)

	BeforeEach(func() {
		ctx, err = ErrorQueue()
	})

	Context("executing error queue", func() {
		// TaskErr throws an error, thus there must be an error returned.
		It("should throw error", func() {
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("TaskErr1"))
		})

		// Task1 is called before TaskErr, thus it must NOT be empty.
		It("should create correct context value for task1", func() {
			Expect(ctx.Task1).To(Equal("task1"))
		})

		// Task2 is called after TaskErr, thus it must be empty.
		It("should create no context value for task2", func() {
			Expect(ctx.Task2).To(Equal(0))
		})

		// Task3 is not called in this queue, thus it must be empty.
		It("should create no context value for task3", func() {
			Expect(ctx.Task3).To(HaveLen(0))
		})

		// Task4 is called after TaskErr, thus it must be empty.
		It("should create no context value for task4", func() {
			Expect(ctx.Task4).To(Equal(float64(0)))
		})
	})
})
