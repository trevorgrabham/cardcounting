package html

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
)

type IndexData struct {
	Settings ElementData
}

type ButtonData struct {
	Button ElementData
	Text string
}

type HandData struct {
	Cards []cards.Card
	Value uint8
}

type TrainingData struct {
	Dealer HandData 
	Player HandData
	Split ButtonData 
	Stand ButtonData
	Hit ButtonData
	Double ButtonData
}

type TrainingMessageData struct {
	Message string 
	ButtonData
}
// type CountValidationData struct {
// 	Correct bool 
// 	Count int16
// }

// type TrainingData struct {
// 	Dealer []cards.Card
// 	Player []cards.Card
// }

// func NewCountValidationData(correct bool, count int16) CountValidationData {
// 	return CountValidationData{
// 		Correct: correct,
// 		Count: count}
// }

type Attribute template.HTMLAttr
type ElementDataFunc func(*ElementData)
type ElementData struct {
	Attributes []Attribute
	Errors []string
}

func NewElementData(funcs ...ElementDataFunc) (e ElementData) {
	for _, f := range funcs {
		f(&e)
	}
	return 
}

func WithID(id string) ElementDataFunc {
	return func(e *ElementData) {
		e.Attributes = append(e.Attributes, Attribute(fmt.Sprintf(`id="%s"`, id)))
	}
}

func WithClasses(classes ...string) ElementDataFunc {
	return func(e *ElementData) {
		e.Attributes = append(e.Attributes, Attribute(fmt.Sprintf(`class="%s"`, strings.Join(classes, " "))))
	}
}

// Pairings of the form hx-'attr', 'value', where 'attr' has the 'hx-' prefix removed
// i.e. "get", "/index", "target", "body", "swap", "beforeend"
func WithHTMX(pairings ...string) ElementDataFunc {
	return func(e *ElementData) {
		for i := 0; i < len(pairings)-1; i+=2 {
			e.Attributes = append(e.Attributes, Attribute(fmt.Sprintf(`hx-%s=%s`, pairings[i], pairings[i+1])))
		}
	}
}

func WithHyperscript(events ...string) ElementDataFunc {
	return func(e *ElementData) {
		var s strings.Builder 
		s.WriteString(`_="`)
		for _, ev := range events {
			s.WriteString(ev)
			s.WriteByte('\n')
		}
		s.WriteByte('"')
	}
}

func WithAttrsNoValue(attrs ...string) ElementDataFunc {
	return func(e *ElementData) {
		for _, attr := range attrs {
			e.Attributes = append(e.Attributes, Attribute(attr))
		}
	}
}

// Pairings of the form 'attr', 'value'
// i.e. "type", "text", "placeholder", "name", "name", "name"
func WithAttrsWithValues(pairings ...string) ElementDataFunc {
	return func(e *ElementData) {
		for i := 0; i < len(pairings)-1; i += 2 {
			e.Attributes = append(e.Attributes, Attribute(fmt.Sprintf(`%s="%s"`, pairings[i], pairings[i+1])))
		}
	}
}

func WithErrors(errors ...string) ElementDataFunc {
	return func(e *ElementData) {
		e.Errors = append(e.Errors, errors...)
	}
}
