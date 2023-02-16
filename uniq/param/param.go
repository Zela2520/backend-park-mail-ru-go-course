package param

import (
	"flag"

	"github.com/pkg/errors"
)

type Param struct {
	Option        string
	OptionMessage string
	OptionValue   interface{}
}

func NewBoolParam(initParam string, objMessage string, objValue *bool) *Param {
	return &Param{
		Option:        initParam,
		OptionMessage: objMessage,
		OptionValue:   objValue,
	}
}

func NewIntParam(initParam string, objMessage string, objValue *int) *Param {
	return &Param{
		Option:        initParam,
		OptionMessage: objMessage,
		OptionValue:   objValue,
	}
}

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

func flagsInit(boolFlags *[]bool, intFlags *[]int, paramsList []Param) ([]Param, error) {
	i := 0
	for _, val := range *boolFlags {
		paramsList[i].OptionValue = val
		i++
	}
	for _, val := range *intFlags {
		paramsList[i].OptionValue = val
		i++
	}

	if len(flag.Args()) > 2 {
		return paramsList, errors.Wrap(errors.New("More arguments passed than required"), "flagInit function")
	}

	for _, val := range flag.Args() {
		if val != "" {
			curParam := NewStringParam("filepath", "Path to the file", val)
			paramsList = append(paramsList, *curParam)
		}
	}

	return paramsList, nil
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

	boolFlags, err = handleBoolParam(&paramsList, boolParams, numberOfBoolFlags)
	if err != nil {
		return nil, errors.Wrap(err, "Getparam function")
	}

	intFlags, err = handleIntParam(&paramsList, intParams, numberOfIntegerFlags)
	if err != nil {
		return nil, errors.Wrap(err, "Getparam function")
	}

	flag.Parse()

	paramsList, err = flagsInit(boolFlags, intFlags, paramsList)
	if err != nil {
		return nil, errors.Wrap(err, "Getparam function")
	}

	return paramsList, nil
}
