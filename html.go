package gofakeit

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6/data"
)

// InputName will return a random input field name
func InputName() string {
	return inputName(globalFaker.Rand)
}

// InputName will return a random input field name
func (f *Faker) InputName() string {
	return inputName(f.Rand)
}

func inputName(r *rand.Rand) string {
	return getRandValue(r, []string{"html", "input_name"})
}

type SVGOptions struct {
	Height int
	Width  int
	Type   string
	Colors []string
}

// Generate a random svg generator
func Svg(options *SVGOptions) string { return svg(globalFaker.Rand, options) }

// Generate a random svg generator
func (f *Faker) Svg(options *SVGOptions) string { return svg(f.Rand, options) }

func svg(r *rand.Rand, options *SVGOptions) string {
	// If options is nil, set it to empty struct
	if options == nil {
		options = &SVGOptions{}
	}

	// If options height and weight is not set, set it to random number between 100 and 500
	if options.Width == 0 {
		options.Width = number(r, 100, 500)
	}
	widthStr := strconv.Itoa(options.Width)
	if options.Height == 0 {
		options.Height = number(r, 100, 500)
	}
	heightStr := strconv.Itoa(options.Height)

	// Check if type is set, if not set to random type
	if options.Type == "" {
		options.Type = randomString(r, data.GetSubData("html", "svg"))
	}

	// If the colors are not set, set it to a set of nice colors
	if len(options.Colors) == 0 {
		options.Colors = niceColors(r)
	}

	// Start svg string
	svgStr := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ` + widthStr + ` ` + heightStr + `" width="` + widthStr + `" height="` + heightStr + `">`

	// Add a rect for the background
	svgStr += `<rect x="0" y="0" width="100%" height="100%" fill="` + randomString(r, options.Colors) + `" />`

	// Add a random number of shapes
	for i := 0; i < number(r, 10, 20); i++ {
		// Add a random shape
		switch options.Type {
		case "rect":
			svgStr += `<rect x="` + strconv.Itoa(number(r, 0, options.Width)) + `" y="` + strconv.Itoa(number(r, 0, options.Height)) + `" width="` + strconv.Itoa(number(r, 0, options.Width)) + `" height="` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "circle":
			svgStr += `<circle cx="` + strconv.Itoa(number(r, 0, options.Width)) + `" cy="` + strconv.Itoa(number(r, 0, options.Height)) + `" r="` + strconv.Itoa(number(r, 0, options.Width)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "ellipse":
			svgStr += `<ellipse cx="` + strconv.Itoa(number(r, 0, options.Width)) + `" cy="` + strconv.Itoa(number(r, 0, options.Height)) + `" rx="` + strconv.Itoa(number(r, 0, options.Width)) + `" ry="` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "line":
			svgStr += `<line x1="` + strconv.Itoa(number(r, 0, options.Width)) + `" y1="` + strconv.Itoa(number(r, 0, options.Height)) + `" x2="` + strconv.Itoa(number(r, 0, options.Width)) + `" y2="` + strconv.Itoa(number(r, 0, options.Height)) + `" stroke="` + randomString(r, options.Colors) + `" />`
		case "polyline":
			svgStr += `<polyline points="` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "polygon":
			svgStr += `<polygon points="` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		}
	}

	// End svg string
	svgStr += `</svg>`

	return svgStr
}

func addHtmlLookup() {
	AddFuncLookup("inputname", Info{
		Display:     "Input Name",
		Category:    "html",
		Description: "Random HTML input field name",
		Example:     "first_name",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return inputName(r), nil
		},
	})

	AddFuncLookup("svg", Info{
		Display:     "Image SVG",
		Category:    "html",
		Description: "Random svg and ",
		Example: `<svg width="369" height="289">
	<rect fill="#4f2958" />
	<polygon points="382,87 418,212 415,110" fill="#fffbb7" />
</svg>`,
		Output:      "string",
		ContentType: "image/svg+xml",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Height in px"},
			{Field: "type", Display: "Type", Type: "string", Optional: true, Options: data.GetSubData("html", "svg"), Description: "Sub child element type"},
			{Field: "colors", Display: "Colors", Type: "[]string", Optional: true, Description: "Hex or RGB array of colors to use"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			// Setup new options
			options := SVGOptions{}
			var err error

			options.Width, err = info.GetInt(m, "width")
			if err != nil {
				return nil, err
			}
			if options.Width < 10 || options.Width >= 1000 {
				return nil, errors.New("invalid image width, must be greater than 10, less than 1000")
			}

			options.Height, err = info.GetInt(m, "height")
			if err != nil {
				return nil, err
			}
			if options.Height < 10 || options.Height >= 1000 {
				return nil, errors.New("invalid image height, must be greater than 10, less than 1000")
			}

			options.Type, err = info.GetString(m, "type")
			svgData := data.GetSubData("html", "svg")
			if err != nil {
				return nil, err
			}

			// If type is empty, set with random type
			if options.Type == "" {
				options.Type = randomString(r, svgData)
			}

			// If not in date html svg type array, return error
			if !stringInSlice(options.Type, svgData) {
				return nil, errors.New("invalid svg type, must be one of " + strings.Join(svgData, ","))
			}

			// Get colors
			options.Colors, err = info.GetStringArray(m, "colors")
			if err != nil {
				return nil, err
			}

			// If colors is empty, set with random colors
			if len(options.Colors) == 0 {
				options.Colors = niceColors(r)
			}

			return svg(r, &options), nil
		},
	})
}
