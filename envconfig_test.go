package envconfig

import (
	"testing"
)

func TestConfig(t *testing.T) {

	t.Run("ENV undefined", func(t *testing.T) {
		const wanted = ModeDefault
		err := Load()

		if err != nil {
			t.Errorf("%s|%v\n", t.Name(), err)
		}

		if Mode != wanted {
			t.Errorf("%s|%s != %s \n", t.Name(), Mode, wanted)
		}
	})

	t.Run("ENV invalid", func(t *testing.T) {
		t.Setenv("ENV", "invalid")
		const wanted = ModeDefault
		err := Load()

		if err == nil {
			t.Errorf("%s|%v\n", t.Name(), err)
		}

		if Mode != ModeDefault {
			t.Errorf("%s|%s != %s \n", t.Name(), Mode, wanted)
		}

	})

	t.Run("ENV staging", func(t *testing.T) {
		t.Setenv("ENV", "staging")
		const wanted = Staging
		err := Load()

		if err != nil {
			t.Errorf("%s|%v\n", t.Name(), err)
		}

		if Mode != wanted {
			t.Errorf("%s|%s != %s \n", t.Name(), Mode, wanted)
		}

	})

}
