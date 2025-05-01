package servercontext

import (
	"errors"
	"fmt"
	"testing"
)

func TestContextStop(t *testing.T) {
	ctx := NewServerContext()

	ctx.StopTask()
	select {
	case <-ctx.Listen():
		if !ctx.IsStop() {
			t.Errorf("stop failed: IsStop is not true")
		}

		if ctx.Reason() != StopReasonStop {
			t.Errorf("stop failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStop)
		}

		if ctx.Error() != nil {
			t.Errorf("stop failed: error is not nil: %s", ctx.Error().Error())
		}
	default:
		t.Fatalf("stop failed")
	}
}

func TestContextStopError(t *testing.T) {
	t.Run("test with error", func(t *testing.T) {
		ctx := NewServerContext()

		testErr := fmt.Errorf("test error")
		ctx.StopTaskError(testErr)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("stop failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonStop {
				t.Errorf("stop failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStop)
			}

			if !errors.Is(ctx.Error(), testErr) {
				t.Errorf("stop failed: error is not testErr: %v", ctx.Error())
			}
		default:
			t.Fatalf("stop failed")
		}
	})

	t.Run("test without error", func(t *testing.T) {
		ctx := NewServerContext()
		ctx.StopTaskError(nil)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("stop failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonStop {
				t.Errorf("stop failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStop)
			}

			if ctx.Error() != nil {
				t.Errorf("stop failed: error is not nil: %s", ctx.Error().Error())
			}
		default:
			t.Fatalf("stop failed")
		}
	})
}

func TestContextStopAllTask(t *testing.T) {
	ctx := NewServerContext()

	ctx.StopAllTask()
	select {
	case <-ctx.Listen():
		if !ctx.IsStop() {
			t.Errorf("stop failed: IsStop is not true")
		}

		if ctx.Reason() != StopReasonStopAllTask {
			t.Errorf("stop failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStopAllTask)
		}

		if ctx.Error() != nil {
			t.Errorf("stop failed: error is not nil: %s", ctx.Error().Error())
		}
	default:
		t.Fatalf("stop failed")
	}
}

func TestContextStopAllTaskError(t *testing.T) {
	t.Run("test with error", func(t *testing.T) {
		ctx := NewServerContext()

		testErr := fmt.Errorf("test error")
		ctx.StopAllTaskError(testErr)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("stop all task failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonStopAllTask {
				t.Errorf("stop all task failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStopAllTask)
			}

			if !errors.Is(ctx.Error(), testErr) {
				t.Errorf("stop all task failed: error is not testErr: %v", ctx.Error())
			}
		default:
			t.Fatalf("stop all task failed")
		}
	})

	t.Run("test without error", func(t *testing.T) {
		ctx := NewServerContext()
		ctx.StopAllTaskError(nil)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("stop all task failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonStopAllTask {
				t.Errorf("stop all task failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonStopAllTask)
			}

			if ctx.Error() != nil {
				t.Errorf("stop all task failed: error is not nil: %s", ctx.Error().Error())
			}
		default:
			t.Fatalf("stop all task failed")
		}
	})
}

func TestContextFinish(t *testing.T) {
	ctx := NewServerContext()
	ctx.Finish()

	select {
	case <-ctx.Listen():
		if !ctx.IsStop() {
			t.Errorf("finish failed: IsStop is not true")
		}

		if ctx.Reason() != StopReasonFinish {
			t.Errorf("finish failed: reason is %d (not finish %d)", ctx.Reason(), StopReasonFinish)
		}

		if ctx.Error() != nil {
			t.Errorf("finish failed: error is not nil: %s", ctx.Error().Error())
		}
	default:
		t.Fatalf("finish failed")
	}
}

func TestContextFinishError(t *testing.T) {
	t.Run("test with error", func(t *testing.T) {
		ctx := NewServerContext()

		testErr := fmt.Errorf("test error")
		ctx.FinishError(testErr)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("finish failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonFinish {
				t.Errorf("finish failed: reason is %d (not finish %d)", ctx.Reason(), StopReasonFinish)
			}

			if !errors.Is(ctx.Error(), testErr) {
				t.Errorf("finish failed: error is not testErr: %v", ctx.Error())
			}
		default:
			t.Fatalf("finish failed")
		}
	})

	t.Run("test without error", func(t *testing.T) {
		ctx := NewServerContext()
		ctx.FinishError(nil)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("finish failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonFinish {
				t.Errorf("finish failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonFinish)
			}

			if ctx.Error() != nil {
				t.Errorf("finish failed: error is not nil: %s", ctx.Error().Error())
			}
		default:
			t.Fatalf("finish failed")
		}
	})
}

func TestContextFinishAndStopAllTask(t *testing.T) {
	ctx := NewServerContext()
	ctx.FinishAndStopAllTask()

	select {
	case <-ctx.Listen():
		if !ctx.IsStop() {
			t.Errorf("stop failed: IsStop is not true")
		}

		if ctx.Reason() != StopReasonFinishAndStopAllTask {
			t.Errorf("stop failed: reason is %d (not finish %d)", ctx.Reason(), StopReasonFinishAndStopAllTask)
		}

		if ctx.Error() != nil {
			t.Errorf("stop failed: error is not nil: %s", ctx.Error().Error())
		}
	default:
		t.Fatalf("stop failed")
	}
}

func TestContextFinishErrorAndStopAllTask(t *testing.T) {
	t.Run("test with error", func(t *testing.T) {
		ctx := NewServerContext()

		testErr := fmt.Errorf("test error")
		ctx.FinishErrorAndStopAllTask(testErr)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("finish and stop all task failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonFinishAndStopAllTask {
				t.Errorf("finish and stop all task failed: reason is %d (not finish %d)", ctx.Reason(), StopReasonFinishAndStopAllTask)
			}

			if !errors.Is(ctx.Error(), testErr) {
				t.Errorf("finish and stop all task failed: error is not testErr: %v", ctx.Error())
			}
		default:
			t.Fatalf("finish and stop all task failed")
		}
	})

	t.Run("test without error", func(t *testing.T) {
		ctx := NewServerContext()
		ctx.FinishErrorAndStopAllTask(nil)
		select {
		case <-ctx.Listen():
			if !ctx.IsStop() {
				t.Errorf("finish and stop all task failed: IsStop is not true")
			}

			if ctx.Reason() != StopReasonFinishAndStopAllTask {
				t.Errorf("finish and stop all task failed: reason is %d (not stop %d)", ctx.Reason(), StopReasonFinishAndStopAllTask)
			}

			if ctx.Error() != nil {
				t.Errorf("finish and stop all task failed: error is not nil: %s", ctx.Error().Error())
			}
		default:
			t.Fatalf("finish and stop all task failed")
		}
	})
}
