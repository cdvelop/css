package css

import (
	"strings"
)

// CSS class structure that maintains order
type class struct {
	Name       string   //eg: "normal", "border", "width-auto"
	Properties []string //eg: "width: 100%", "min-width: 100%"
}

// Method to add a class to the stylesheet
func AddClass(name string) *class {
	// Check if the class already exists
	for _, existingClass := range ss.classes {
		if existingClass.Name == name {
			return existingClass
		}
	}

	new := &class{
		Name:       name,
		Properties: []string{},
	}

	ss.classes = append(ss.classes, new)
	return new
}

// Method to add a CSS property with multiple values
// Example: AddProperty("margin", "10px", "20px") -> margin: 10px 20px
func (c *class) AddProperty(key string, values ...string) *class {
	property := key + ": " + strings.Join(values, " ")
	c.Properties = append(c.Properties, property)
	return c
}

// getClassName returns the class name to use in HTML
// func (c *class) getClassName() string {
// 	return "." + c.Name
// }

// Method to generate CSS respecting insertion order
func (c *class) GenerateCSS() string {
	var cssBuilder strings.Builder

	// Pre-estimate size
	estimatedSize := len(c.Name) + 50
	for _, prop := range c.Properties {
		estimatedSize += len(prop) + 20
	}
	cssBuilder.Grow(estimatedSize)

	cssBuilder.WriteString(".")
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
