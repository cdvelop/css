package css

import (
	"testing"
)

func TestVars(t *testing.T) {

	expected := `:root {
    --FontSizeNormal: 1.1rem;
    --FontSizeSmall: .6rem;
    --ColorPrimary: #ffffff;
    --ColorSecondary: #3f88bf;
    --ColorTertiary: #c2c1c1;
    --ColorQuaternary: #000000;
    --ColorGray: #e9e9e9;
    --ColorSelection: #ff9300;
    --ColorHover: #ff95008e;
    --ColorSuccess: #aadaff7c;
    --ColorError: #f20707;
    --MenuSize: 6vh;
    --ContentHeight: 94vh;
    --ContentWidth: 100vw;
    --TransitionWait: 0s;
}
`

	s := NewStyleSheet()
	result := s.generateRoot()

	// Verificaciones de contenido
	if result != expected {
		t.Errorf("error:\nresult:\n[%v]\nexpected:\n[%v]", result, expected)
	}

	s.Vars.ColorPrimary = "#000000"
	// result = css.GenerateRoot()
	if s.Vars.ColorPrimary != "#000000" {
		t.Errorf("error:\nresult:\n[%v]\nexpected:\n[%v]", result, expected)
	}

	// fmt.Println(result)

}
