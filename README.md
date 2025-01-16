# CSS Generator Package

Package for generating CSS classes programmatically in Go.

## API Reference

### Methods

#### AddClass(name string) *class
Creates a new CSS class with the specified name. If a class with the same name already exists, returns the existing class.

#### AddProperty(key string, values ...string) *class
Adds a CSS property to the class. Supports multiple values for properties like margin, padding, etc.

#### GenerateCSS() string
Generates the CSS code for the class respecting the insertion order of properties.

## Usage Example

```go
package main

import (
    "fmt"
    "github.com/cdvelop/css"
)

func main() {
    // Create stylesheet
    stylesheet := &css.Stylesheet{}

    // Create button class
    buttonClass := stylesheet.AddClass("btn-primary").
        AddProperty("font-family", "Arial", "Helvetica", "sans-serif").
        AddProperty("background", "linear-gradient(to right)", "blue", "purple").
        AddProperty("padding", "10px", "15px").
        AddProperty("box-shadow", "0 2px 4px rgba(0,0,0,0.1)")

    // Create card class  
    cardClass := stylesheet.AddClass("card").
        AddProperty("border", "1px", "solid", "#ccc").
        AddProperty("transition", "all", "0.3s", "ease-in-out")

    // Generate CSS
    css := stylesheet.GenerateStylesheet()
    fmt.Println(css)
}
```

## Generated CSS Example

```css
.btn-primary {
    font-family: Arial Helvetica sans-serif;
    background: linear-gradient(to right) blue purple;
    padding: 10px 15px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.card {
    border: 1px solid #ccc;
    transition: all 0.3s ease-in-out;
}
```

## Notes

- Property values are space-separated in the generated CSS
- Class names are prefixed with '.' in the generated CSS
- Property order is preserved as added