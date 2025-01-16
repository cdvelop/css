package css

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"unicode"
)

// Structure to handle multiple classes
type stylesheet struct {
	classes []*class
}

// NewStyleSheet creates and returns a new stylesheet instance
func NewStyleSheet() *stylesheet {
	return &stylesheet{
		classes: []*class{},
	}
}

type cssVars struct {
	/* Font Sizes */
	FontSizeNormal string
	FontSizeSmall  string
	/* Colors */
	ColorPrimary    string
	ColorSecondary  string
	ColorTertiary   string
	ColorQuaternary string
	ColorGray       string
	ColorSelection  string
	ColorHover      string
	ColorSuccess    string
	ColorError      string
	/* Layout Sizes */
	MenuSize      string
	ContentHeight string
	ContentWidth  string
	/* Timing */
	TransitionWait string

	externalVars map[string]string
}

var Var = cssVars{
	FontSizeNormal:  "1.1rem",
	FontSizeSmall:   ".6rem",
	ColorPrimary:    "#ffffff",
	ColorSecondary:  "#3f88bf",
	ColorTertiary:   "#c2c1c1",
	ColorQuaternary: "#000000",
	ColorGray:       "#e9e9e9",
	ColorSelection:  "#ff9300",
	ColorHover:      "#ff95008e",
	ColorSuccess:    "#aadaff7c",
	ColorError:      "#f20707",
	MenuSize:        "6vh",
	ContentHeight:   "94vh",
	ContentWidth:    "100vw",
	TransitionWait:  "0s",

	externalVars: map[string]string{},
}

// agregar una variable CSS externa
func (c *cssVars) AddVariable(name, value string) {
	c.externalVars[name] = value
}

// obtener el valor de una variable (en formato CSS: var(--miVar))
func GetVariable(name string) string {
	return `var(` + name + `)`
}

// generateRoot genera la clase ":root" con todas las variables CSS
func generateRoot() string {
	var sb strings.Builder
	sb.WriteString(":root {\n")

	// Usa reflect para iterar sobre los campos de cssVars
	v := reflect.ValueOf(Var)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).String()

		// skip if it starts with lowercase
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		sb.WriteString("    --" + fieldName + ": " + fieldValue + ";\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

// Method to generate the entire stylesheet and optionally save it to a file
// Example usage:
//
//	sheet := NewStyleSheet()
//	css, err := sheet.Generate()              // Generate stylesheet string only
//	css, err := sheet.Generate("styles.css")  // Generate and save to file
func (s *stylesheet) Generate(paths ...string) (string, error) {
	// Use strings.Builder with initial capacity
	var stylesheetBuilder strings.Builder

	stylesheetBuilder.WriteString(generateRoot())

	for _, class := range s.classes {
		stylesheetBuilder.WriteString(class.GenerateCSS())
	}

	css := stylesheetBuilder.String()

	if len(paths) > 0 {
		path := paths[0]
		if !strings.HasSuffix(path, ".css") {
			return css, errors.New("file path must end with .css extension")
		}

		err := os.WriteFile(path, []byte(css), 0644)
		if err != nil {
			return css, errors.New("error writing css file: " + err.Error())
		}
	}

	return css, nil
}
