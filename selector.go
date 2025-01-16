package css

import (
	"strings"
)

// class structure that maintains order
type selector struct {
	Name       string   //eg: "normal", "border", "width-auto"
	Properties []string //eg: "width: 100%", "min-width: 100%"
}

// Method to add a class to the StyleSheet
// Examples:
// - AddSelector(".my-class") -> for class selectors
// - AddSelector("#my-id") -> for ID selectors
// - AddSelector("div") -> for element selectors
// - AddSelector("div > p") -> for child selectors
// - AddSelector("div.my-class") -> for element with class
func (s *StyleSheet) AddSelector(name string) *selector {
	// Check if the class already exists
	for _, existingClass := range s.selectors {
		if existingClass.Name == name {
			return existingClass
		}
	}

	new := &selector{
		Name:       name,
		Properties: []string{},
	}

	s.selectors = append(s.selectors, new)
	return new
}

// Method to add a CSS property with multiple values
// Example: AddProperty("margin", "10px", "20px") -> margin: 10px 20px
func (c *selector) AddProperty(key string, values ...string) *selector {
	property := key + ": " + strings.Join(values, " ")

	// Check if property already exists
	for _, prop := range c.Properties {
		if prop == property {
			return c
		}
	}

	c.Properties = append(c.Properties, property)
	return c
}

// getClassName returns the class name to use in HTML
// func (c *class) getClassName() string {
// 	return "." + c.Name
// }

// Method to generate CSS respecting insertion order
func (c *selector) GenerateCSS() string {
	var cssBuilder strings.Builder

	// Pre-estimate size
	estimatedSize := len(c.Name) + 50
	for _, prop := range c.Properties {
		estimatedSize += len(prop) + 20
	}
	cssBuilder.Grow(estimatedSize)

	cssBuilder.WriteString(c.Name)
	cssBuilder.WriteString(" {\n")

	for _, prop := range c.Properties {
		cssBuilder.WriteString("    ")
		cssBuilder.WriteString(prop)
		cssBuilder.WriteString(";\n")
	}

	cssBuilder.WriteString("}\n")

	return cssBuilder.String()
}
