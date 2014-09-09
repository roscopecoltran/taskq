/*
Usage

  // A custom context tasks can operate on.
  type Ctx struct {
    Task1 string
    Task2 int
  }

  // Simple tasks fullfilling the Task interface.
  func Task1(ctx interface{}) error {
    ctx.(*Ctx).Task1 = "task1"
    return nil
  }

  func Task2(ctx interface{}) error {
    ctx.(*Ctx).Task2 = 2
    return nil
  }

  func main() {
    // Creating a pointer to your custom context.
    ctx := &Ctx{}

    // Create a new queue configured with you context.
    q := taskqPkg.NewQueue(ctx)

    // Run your tasks in a specified order.
    err := q.RunTasks(
      taskqPkg.InParallel(
        Task1,
        Task2,
      ),
    )

    // As soon as an error occurs, no more tasks will be executed. That is, the
    // first occuring error will be returned.
    if err != nil {
      panic(err)
    }

    // Output:
    // &taskq_test.Ctx{Task1:"task1", Task2:2}
    fmt.Printf("%#v\n", ctx)
  }
*/

package taskq
