package testrunner

var Q2 = QuestionData{
	Title:      "Sum list",
	Prompt:     "Write a function thats sums a list.",
	numParams:  1,
	paramTypes: []string{"list int"},
	methodName: "addLots",
	NumCases:   5,
	VisibleCases: 3,
	Cases: [][][]string{{{"1", "2"}},
		{{"-10", "7"}},
		{{"2147483647", "-2147483648"}},
		{{"-543", "543"}},
		{{"-500", "4500"}}},
	ExpectedResults: []string{"1", "3", "1", "-3", "1", "-1", "1", "0", "1", "4000"},
	returnType:      "list int",
	Templates: LanguageFunctionTemplates{
		Python:     "def addLots(ls:list):\n    \n",
		Javascript: "function addLots(ls) {\n    \n}",
		Cpp:        "vector<int> addLots(vector<int> list) {\n    \n}",
	},
}
