package entity_test

import (
	"go-rest-echo/entity"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestIncrement(t *testing.T) {
	t.Run("1 Struct Increment Field", func(t *testing.T) {
		got := entity.Increment{ID: 1}
		want := entity.Increment{ID: 1}

		if got.ID != want.ID {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})

	t.Run("2 Struct Increment Pointer", func(t *testing.T) {
		got := &entity.Increment{ID: 1}
		want := &entity.Increment{ID: 1}

		if got == want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})
}

func TestUUID(t *testing.T) {
	t.Run("1 Struct UUID Field", func(t *testing.T) {
		got := entity.UUID{ID: "a"}
		want := entity.UUID{ID: "a"}

		if got.ID != want.ID {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})

	t.Run("2 Struct UUID Pointer", func(t *testing.T) {
		got := &entity.UUID{ID: "a"}
		want := &entity.UUID{ID: "a"}

		if got == want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})
}

func TestTimestamps(t *testing.T) {
	t.Run("1 Struct Timestamps Field", func(t *testing.T) {
		t.Run("1.1 Field CreatedAt", func(t *testing.T) {
			got := entity.Timestamps{CreatedAt: time.Time{}}
			want := entity.Timestamps{CreatedAt: time.Time{}}

			if got != want {
				t.Errorf("Expected `%v`, but got %v", want, got)
			}
		})

		t.Run("1.2 Field UpdatedAt", func(t *testing.T) {
			got := entity.Timestamps{UpdatedAt: time.Time{}}
			want := entity.Timestamps{UpdatedAt: time.Time{}}

			if got != want {
				t.Errorf("Expected `%v`, but got %v", want, got)
			}
		})

		t.Run("1.3 Field DeletedAt", func(t *testing.T) {
			got := entity.Timestamps{DeletedAt: gorm.DeletedAt{}}
			want := entity.Timestamps{DeletedAt: gorm.DeletedAt{}}

			if got != want {
				t.Errorf("Expected `%v`, but got %v", want, got)
			}
		})
	})

	t.Run("2 Struct Timestamps Pointer", func(t *testing.T) {
		got := &entity.Timestamps{}
		want := &entity.Timestamps{}

		if got == want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})
}
