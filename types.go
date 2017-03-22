package argparser

import (
	"strconv"
)

type Verb struct {
	Name string
	Function func(*Verb)
	Options []Option
	ChildVerbs []*Verb
}
func (v *Verb) CallFunction() { v.Function(v) }

type Option interface {
	IsRequired() bool
	IsUsed() bool
	IsHaveFunction() bool
	CallFunction()
	GetNumberArgs() int
	GetNumberArgsMax() int
	SetValue(string)
	GetLongName() string
	GetShortName() string
}

type String struct {
	Required bool
	ArgsMax int
	value []string
	LongName string
	ShortName string
	Function func(*String)
}
func (s String) IsRequired() bool { return s.Required }
func (s String) IsUsed() bool { if len(s.value) != 0 {return true} else {return false}}
func (s String) IsHaveFunction() bool { if s.Function != nil { return true} else {return false}}
func (s *String) CallFunction() { s.Function(s) }
func (s String) GetNumberArgs() int { return s.ArgsMax }
func (s String) GetNumberArgsMax() int { return len(s.value) }
func (s String) GetLongName() string { return s.LongName }
func (s String) GetShortName() string { return s.ShortName }
func (s *String) SetValue( val string) { s.value = append(s.value, val) }
func (s String) Value() []string { return s.value }


type Int struct {
	Required bool
	ArgsMax int
	value []int
	LongName string
	ShortName string
	Function func(*Int)
}
func (i Int) IsRequired() bool { return i.Required }
func (i Int) IsUsed() bool { if len(i.value) != 0 {return true} else {return false}}
func (i Int) IsHaveFunction() bool { if i.Function != nil { return true} else {return false}}
func (i *Int) CallFunction() { i.Function(i) }
func (i Int) GetNumberArgs() int { return i.ArgsMax }
func (i Int) GetNumberArgsMax() int { return len(i.value) }
func (i Int) GetLongName() string { return i.LongName }
func (i Int) GetShortName() string { return i.ShortName }
func (i *Int) SetValue( valStr string) {
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return
	}
	i.value = append(i.value, val)
}
func (i Int) Value() []int { return i.value }



type Bool struct {
	value bool
	LongName string
	ShortName string
	Function func(*Bool)
}
func (Bool) IsRequired() bool { return false }
func (b Bool) IsUsed() bool { if b.value {return true} else {return false}}
func (b Bool) IsHaveFunction() bool { if b.Function != nil { return true} else {return false}}
func (b *Bool) CallFunction() { b.Function(b) }
func (Bool) GetNumberArgs() int { return 0 }
func (Bool) GetNumberArgsMax() int { return 0 }
func (b Bool) GetLongName() string { return b.LongName }
func (b Bool) GetShortName() string { return b.ShortName }
func (b *Bool) SetValue(string) { b.value = true }
func (b Bool) Value() bool { return b.value}