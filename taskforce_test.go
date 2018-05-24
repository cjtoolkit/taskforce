// +build debug

package taskforce

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTaskForce(t *testing.T) {
	type Mocks struct {
		util *MockutilI
	}

	let := func(t *testing.T) (Mocks, *TaskForce) {
		ctrl := gomock.NewController(t)
		mocks := Mocks{
			util: NewMockutilI(ctrl),
		}

		subject := &TaskForce{
			util: mocks.util,
		}

		return mocks, subject
	}

	t.Run("CheckError", func(t *testing.T) {
		t.Run("No Error", func(t *testing.T) {
			_, subject := let(t)

			subject.CheckError(nil)
		})

		t.Run("Has Error", func(t *testing.T) {
			mocks, subject := let(t)
			err := errors.New("I am error")

			mocks.util.EXPECT().DoPanic(err)

			subject.CheckError(err)
		})
	})

	t.Run("recover", func(t *testing.T) {
		t.Run("Does not panic", func(t *testing.T) {
			mocks, subject := let(t)

			mocks.util.EXPECT().DisplaySuccess()

			defer subject.recover()
		})

		t.Run("Does panic", func(t *testing.T) {
			mocks, subject := let(t)
			err := errors.New("I am error")

			mocks.util.EXPECT().DoRecover(err)

			defer subject.recover()
			panic(err)
		})
	})
}
