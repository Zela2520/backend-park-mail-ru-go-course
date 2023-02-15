package uniq

import (
	"flag"

	"github.com/pkg/errors"
)

const (
	numberOfBoolFlags         = 4
	numberOfIntegerFlags      = 2
	countFlag                 = "c" // bool
	countFlagMessage          = "count the number of occurrences of a string in the input. Output this number before the string separated by a space."
	repeatFlag                = "d" // bool
	repeatFlagMessage         = "output only those lines that are repeated in the input"
	uniqFlag                  = "u" // bool
	uniqFlagMessage           = "output only those lines that are not repeated in the input"
	compareNWord              = "f" // integer
	compareNWordMessage       = "compare starting from the nth word"
	compareNCharacters        = "s" // integer
	compareNCharactersMessage = "compare starting from the nth character"
	arbitraryСaseFlag         = "i" // bool
	arbitraryСaseFlagMessage  = "ignore letter case"
)

type Param struct {
	Option        string
	OptionMessage string
	OptionValue   interface{}
}

// создание bool параметра (3-ий аргумент по указателю, потому что мы должны хранить именно адрес на флаг)
func NewBoolParam(initParam string, objMessage string, objValue *bool) *Param {
	return &Param{
		Option:        initParam,
		OptionMessage: objMessage,
		OptionValue:   objValue,
	}
}

// создание int параметра (3-ий аргумент по указателю, потому что мы должны хранить именно адрес на флаг)
func NewIntParam(initParam string, objMessage string, objValue *int) *Param {
	return &Param{
		Option:        initParam,
		OptionMessage: objMessage,
		OptionValue:   objValue,
	}
}

// создание string параметра (3-ий аргумент по указателю, потому что мы должны хранить именно адрес на флаг)
func NewStringParam(initParam string, objMessage string, objValue string) *Param {
	return &Param{
		Option:        initParam,
		OptionMessage: objMessage,
		OptionValue:   objValue,
	}
}

func handleBoolParam(params *[]Param, srcParams []Param, requiredParamNumber int) (*[]bool, error) {
	var (
		boolFlags []bool
	)

	for i := 0; i < requiredParamNumber; i++ {
		boolFlags = append(boolFlags, false)
	}

	if requiredParamNumber != len(srcParams) {
		return &boolFlags, errors.Wrap(errors.New("The specified number of options does not match their actual number"), "handleBoolParam function")
	}

	for i := 0; i < requiredParamNumber; i++ {
		flag.BoolVar(&boolFlags[i], srcParams[i].Option, false, srcParams[i].OptionMessage)
		*params = append(*params, srcParams[i])
	}

	return &boolFlags, nil
}

func handleIntParam(destParams *[]Param, srcParams []Param, requiredParamNumber int) (*[]int, error) {
	var (
		intFlags []int
	)

	for i := 0; i < requiredParamNumber; i++ {
		intFlags = append(intFlags, 0)
	}

	if requiredParamNumber != len(srcParams) {
		return &intFlags, errors.Wrap(errors.New("The specified number of options does not match their actual number"), "handleBoolParam function")
	}

	for i := 0; i < requiredParamNumber; i++ {
		flag.IntVar(&intFlags[i], srcParams[i].Option, 0, srcParams[i].OptionMessage)
		*destParams = append(*destParams, srcParams[i])
	}

	return &intFlags, nil
}

func flagsInit(boolFlags *[]bool, intFlags *[]int, paramsList []Param) []Param {
	i := 0
	for _, val := range *boolFlags {
		paramsList[i].OptionValue = val
		i++
	}
	for _, val := range *intFlags {
		paramsList[i].OptionValue = val
		i++
	}
	for _, val := range flag.Args() {
		if val != "" {
			curParam := NewStringParam("filepath", "Path to the file", val)
			paramsList = append(paramsList, *curParam)
		}
	}

	return paramsList
}

func GetParams() ([]Param, error) {
	var (
		boolFlags *[]bool
		intFlags  *[]int
		err       error
	)

	boolParams := []Param{
		{Option: countFlag, OptionMessage: countFlagMessage},
		{Option: repeatFlag, OptionMessage: repeatFlagMessage},
		{Option: uniqFlag, OptionMessage: uniqFlagMessage},
		{Option: arbitraryСaseFlag, OptionMessage: arbitraryСaseFlagMessage},
	}

	intParams := []Param{
		{Option: compareNWord, OptionMessage: compareNWordMessage},
		{Option: compareNCharacters, OptionMessage: compareNCharactersMessage},
	}

	paramsList := make([]Param, 0)
	// TODO: paramsList := make(map[string]interface{})
	// Итоговые данные лучше хранить в словаре (key = Option, value = OptionValue),
	// чтобы можно было по ключу понять значение опции // m["d"] = "string"

	boolFlags, err = handleBoolParam(&paramsList, boolParams, numberOfBoolFlags)
	if err != nil {
		return nil, errors.Wrap(err, "Getparam function")
	}

	intFlags, err = handleIntParam(&paramsList, intParams, numberOfIntegerFlags)
	if err != nil {
		return nil, errors.Wrap(err, "Getparam function")
	}

	flag.Parse()

	paramsList = flagsInit(boolFlags, intFlags, paramsList)

	return paramsList, nil
}
