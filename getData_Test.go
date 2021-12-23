package getdata

import "testing"

func TestMdbConnect(t *testing.T) {
	text := "Hey Dude"
	if out := MdbConnect(); out != text {
		t.Errorf("MbConnect, = %q, want %q", out, text)
	}
}
