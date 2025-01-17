package css_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cdvelop/css"
)

func TestCSSGeneration(t *testing.T) {

	// Test case 1: Basic CSS Generation
	t.Run("Basic CSS Generation", func(t *testing.T) {

		sheet := css.NewStyleSheet()
		sheet.AddSelector(".btn-primary").
			AddProperty("font-family", "Arial", "Helvetica", "sans-serif").
			AddProperty("background", "linear-gradient(to right)", "blue", "purple").
			AddProperty("padding", "10px", "15px").
			AddProperty("box-shadow", "0 2px 4px rgba(0,0,0,0.1)")

		sheet.AddSelector("#card").
			AddProperty("border", "1px", "solid", "#ccc").
			AddProperty("transition", "all", "0.3s", "ease-in-out")

		cssResult, err := sheet.Generate()
		if err != nil {
			t.Errorf("Error generating StyleSheet: %v", err)
		}

		// Content verifications
		expectedParts := []string{
			".btn-primary {",
			"font-family: Arial Helvetica sans-serif;",
			"background: linear-gradient(to right) blue purple;",
			"padding: 10px 15px;",
			"box-shadow: 0 2px 4px rgba(0,0,0,0.1);",
			"#card {",
			"border: 1px solid #ccc;",
			"transition: all 0.3s ease-in-out;",
		}

		for _, part := range expectedParts {
			if !strings.Contains(cssResult, part) {
				t.Errorf("Generated CSS missing expected part: %s", part)
			}
		}
	})

	// Test case 2: Property Order Verification
	t.Run("Property Order Preservation", func(t *testing.T) {

		sheet := css.NewStyleSheet()
		result := sheet.AddSelector(".test-order").
			// Properties in a specific order
			AddProperty("display", "flex").
			AddProperty("justify-content", "center").
			AddProperty("align-items", "center").
			GenerateCSS()

		// Verify property order
		cssLines := strings.Split(result, "\n")

		expectedOrder := []string{
			".test-order {",
			"    display: flex;",
			"    justify-content: center;",
			"    align-items: center;",
			"}",
		}

		for i, expectedLine := range expectedOrder {
			if i < len(cssLines) && strings.TrimSpace(cssLines[i]) != strings.TrimSpace(expectedLine) {
				t.Errorf("Unexpected line order. Expected %s, got %s",
					expectedLine, cssLines[i])
			}
		}
	})

	// Test case 3: Multiple Values
	t.Run("Multiple Values", func(t *testing.T) {
		sheet := css.NewStyleSheet()
		result := sheet.AddSelector("multi-value").
			AddProperty("background", "linear-gradient(45deg)", "red", "blue").
			GenerateCSS()

		expectedValue := "background: linear-gradient(45deg) red blue;"
		if !strings.Contains(result, expectedValue) {
			t.Errorf("Failed to generate correct multiple value property. Got: %s", result)
		}
	})

	// Test case 4: Complete StyleSheet
	t.Run("Full StyleSheet Generation", func(t *testing.T) {
		sheet := css.NewStyleSheet()
		sheet.AddSelector("button").
			AddProperty("color", "white")

		sheet.AddSelector(".card").
			AddProperty("border", "1px", "solid", "black")

		ssResult, err := sheet.Generate()
		if err != nil {
			t.Errorf("Error generating StyleSheet: %v", err)
		}

		expectedClasses := []string{
			"button {",
			"color: white;",
			".card {",
			"border: 1px solid black;",
		}

		for _, expectedClass := range expectedClasses {
			if !strings.Contains(ssResult, expectedClass) {
				t.Errorf("StyleSheet missing expected class: %s", expectedClass)
			}
		}
	})

	// Test case 5: Duplicate Property Prevention
	t.Run("Prevent Duplicate Properties", func(t *testing.T) {
		sheet := css.NewStyleSheet()
		class := sheet.AddSelector(".test-duplicates")

		// Add initial property
		class.AddProperty("color", "red")
		// Try to add same property again
		class.AddProperty("color", "red")

		// Verify only one property exists
		cssResult := class.GenerateCSS()

		expected := `.test-duplicates{color:red;}`
		normalizedResult := strings.ReplaceAll(strings.ReplaceAll(cssResult, " ", ""), "\n", "")
		if normalizedResult != expected {
			t.Errorf("Duplicate property detected. Expected:\n%s\nGot:\n%s", expected, normalizedResult)
		}
	})
}

// Example of benchmark execution for performance
func BenchmarkCSSGeneration(b *testing.B) {
	sheet := css.NewStyleSheet()
	// Prepare a set of classes
	for i := 0; i < 100; i++ {
		sheet.AddSelector(fmt.Sprintf("class-%d", i)).
			AddProperty("color", "black").
			AddProperty("margin", "10px")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sheet.Generate()
	}
}
