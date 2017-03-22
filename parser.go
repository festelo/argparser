package argparser

import (
	"regexp"
	"errors"
)

type parser struct {
	verb *Verb
}
func NewParser( mainVerb *Verb) parser{ return parser{ verb: mainVerb} }

func (p parser) Parse(args []string) error{
	opts, verbs, err := parseAndGetToCall(p.verb, args)


	for _, t := range verbs {
		t.CallFunction()
	}
	for _, t := range opts {
		t.CallFunction()
	}
	return err
}

func parseAndGetToCall(verb *Verb, args []string) ([]Option, []*Verb, error){
	shortMap, longMap, requiredArray, err := OptionsToArrays(verb.Options)
	childVerbsMap :=  VerbsToMap(verb.ChildVerbs)
	if err != nil {
		return nil, nil, err
	}
	shortRegex := regexp.MustCompile("^-([^-](?:.+)?)")
	longRegex := regexp.MustCompile("^--([^-](?:.+)?)")

	toCallOptionFuncArr := []Option{}
	toCallVerbFuncArr := []*Verb{}

	var parsKey Option
	for i, a := range args{
		if childVerb, ok := childVerbsMap[a]; ok {
			opts , verbs, err := parseAndGetToCall(childVerb, args[i+1:])
			if err != nil{
				return nil, nil, err
			} else {
				toCallOptionFuncArr = append(toCallOptionFuncArr, opts...)
				toCallVerbFuncArr = append(toCallVerbFuncArr, childVerb)
				toCallVerbFuncArr = append(toCallVerbFuncArr, verbs...)
				break
			}
		}
		find := shortRegex.FindStringSubmatch(a)
		if len(find) >= 2{
			if shortMap[find[1]] != nil{
				if shortMap[find[1]].IsHaveFunction(){
					toCallOptionFuncArr = appendIfMissing(toCallOptionFuncArr, shortMap[find[1]])
				}
				checkAndAdd(&parsKey, shortMap[find[1]])
				continue
			}
		}

		find = longRegex.FindStringSubmatch(a)
		if len(find) >= 2{
			if longMap[find[1]] != nil{
				if longMap[find[1]].IsHaveFunction(){
					toCallOptionFuncArr = appendIfMissing(toCallOptionFuncArr, longMap[find[1]])
				}
				checkAndAdd(&parsKey, longMap[find[1]])
				continue
			} else { return nil, nil, errors.New("Unknown argument " + a) }
		}

		if parsKey != nil{
			parsKey.SetValue(a)
			if parsKey.GetNumberArgs() == parsKey.GetNumberArgsMax(){
				parsKey = nil
			}
		} else {
			return nil, nil, errors.New("Unknown argument " + a)
		}
	}
	err = getRequiredError(requiredArray)
	if err != nil{
		return nil, nil, err
	}
	return toCallOptionFuncArr, toCallVerbFuncArr, nil
}

func appendIfMissing(slice []Option, i Option) []Option {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func getRequiredError(requiredArr []Option) error {
	isExpected := false
	errStr := "Required argument expected: "
	for _, t := range requiredArr{
		if !t.IsUsed() {
			errStr += t.GetLongName() + ", "
			isExpected = true
		}
	}
	if isExpected{
		return errors.New(errStr)
	} else {
		return nil
	}
}

func checkAndAdd(parent *Option, checkable Option){
	if checkable.GetNumberArgs() == 0{
		checkable.SetValue("")
		*parent = nil
	} else { *parent = checkable }
}

func VerbsToMap(verbArray []*Verb) map[string]*Verb{
	returnMap := make(map[string]*Verb)
	for _, t := range verbArray{
		returnMap[t.Name] = t
	}
	return returnMap
}

func OptionsToArrays(optionsArray []Option) (map[string]Option, map[string]Option, []Option, error){
	shortMap := make(map[string]Option)
	longMap := make(map[string]Option)
	requiredArray := []Option{}
	for _, t := range optionsArray {
		if t.GetShortName() != ""{
			shortMap[t.GetShortName()] = t
		}
		if t.GetLongName() != ""{
			longMap[t.GetLongName()] = t
		}
		if t.IsRequired(){
			requiredArray = append(requiredArray, t)
		}
	}
	if len(shortMap) == 0 && len(longMap) == 0{
		return nil, nil, nil, errors.New("No keys with short and long name to parse")
	}else{
		return shortMap, longMap, requiredArray, nil
	}
}