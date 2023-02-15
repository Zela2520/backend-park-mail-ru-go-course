package uniq

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
