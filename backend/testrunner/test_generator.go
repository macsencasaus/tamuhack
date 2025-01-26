package testrunner

// import "fmt"
import "strconv"
//TODO -> double check booleans and strings
//TODO -> change cases to a struct?
type reflectionData struct {
	numParams       int//follow python convension for everything in this struct
	numCases        int
	paramTypes      []string //int, bool(follow python i.e. True False), string, float, list type(has a space)
	cases           [][][]string //[case][parameter][item] item is just 0 for non lists
	expectedResults []string//if is list each test case starts with a number saying number of values then is sequence of values
	methodName      string
	returnType      string
}

var questionMap = map[int]reflectionData{
	1: Q1, 2:Q2, 3:Q3,
}

var headerMap = map[int]map[string]string{
	1: Q1Template, 2:Q2Template, 3:Q3Template,
}

var pythonToC = map[string]string{
	"bool":"bool", "int":"int", "string":"string","float":"double", "list int":"vector<int>", "list float":"vector<float>",
	"list bool":"vector<bool>", "list string":"vector<string>",
}

	//following is test cases
	//python
	// fmt.Println(generate(`def add(a:int,b:int):
	// return a+b`, "python", "747474747", 1))

	// fmt.Println(generate(`def addLots(ls:list):
	// sum=0
	// for l in ls:
	// 	sum+=l
	// return [sum]`, "python", "747474747", 2))

	// fmt.Println(generate(`def returnList(ls:list):
	// return ls`, "python", "747474747", 3))

	//c++
	// fmt.Println(generate(`int add(int a, int b){
	// return a+b;}`, "c++", "747474747", 1))

	// fmt.Println(generate(`vector<int> addLots(vector<int> list){
	// int sum=0;
	// for(int i=0; i<list.size();i++)	
	// sum+=list.at(i);
	// vector<int> answer = {sum};
	// return answer;}`, "c++", "747474747", 2))

	// fmt.Println(generate(`vector<int> returnList(vector<int> list){return list;}`, "c++", "747474747",3))

	//javascript
	// fmt.Println(generate(`function add(a,b){
    // return a+b;}`, "javascript","747474747",1))

	// fmt.Println(generate(`function addLots(ls){
	// 	let total=0;
	// 	for(let i=0; i<ls.length; i++){
	// 		total+=ls[i];
	// 	}
	// 	return [total];}`, "javascript","747474747",2))

	// fmt.Println(generate(`function returnList(ls){
	// 	return ls;}`, "javascript","747474747",3))

func generate(userInput string, language string, randomNumber string, questionNumber int) string {

	r := questionMap[questionNumber]
	if language == "c++" {
		return generateC(userInput, randomNumber, r)
	} else if language == "python" {
		return generatePython(userInput, randomNumber, r)
	} else if language == "javascript" {
		return generateJavacript(userInput, randomNumber, r)
	}
	return "bruh"
}

func isNotAList(typee string) bool {
	for i:=0; i<len(typee);i++{
		if typee[i]==' '{
			return false
		}
	}
	return true
	// return len(typee) > 4 && typee[0:4]=="List"
}

func generatePython(userInput, randomNumber string, r reflectionData) string {
	answer := userInput
	answer += "\n\n"
	//random number print method Follows:user_output \nmagic_number\n result \nmagic_number\n user_output...
	answer+="def magic(thingToPrint):\n"
	answer+="\tprint('\\n',str('"+string(randomNumber)+"'),'\\n',thingToPrint,'\\n',str('"+string(randomNumber)+"'),'\\n',sep='')\n\n"
	answer += "def main():\n"
	//results array
	if isNotAList(r.returnType) {
		answer += "\texpected_results = ["
		for i := 0; i < len(r.expectedResults); i++ {
			answer += r.expectedResults[i]
			if i+1 != len(r.expectedResults) {
				answer += ","
			}
		}
	} else {//forming expected results as list of lists
		answer += "\texpected_results = ["
		currentIndex :=0
		for j:=0; j<r.numCases;j++{
			answer+="["
			lengthOfCurrentTestCase,_:=strconv.Atoi(r.expectedResults[currentIndex])//get length of this test cases's array
			currentIndex++
			for i := 0; i < lengthOfCurrentTestCase; i++ {
				answer += r.expectedResults[currentIndex]
				currentIndex++
				if i+1 != lengthOfCurrentTestCase {
					answer += ","
				}
			}
			answer+="]"
			if j+1 != r.numCases{
				answer+=","
			}
		}
	}
	answer += "]\n"

	//parameter array
	for i := 0; i < len(r.paramTypes); i++ {//i=current_paramter
		if !(isNotAList(r.paramTypes[i])) {
			//initializes lists and such with name: a<test_case_number>_<parameter_number>
			for j:=0; j<r.numCases; j++{ //j==current_Case
				answer+="\ta"+strconv.Itoa(j)+"_"+strconv.Itoa(i)+" = ["
				for k:=0; k<len(r.cases[j][i]); k++{
					answer+=r.cases[j][i][k]
					if k+1 != len(r.cases[j][i]){
						answer+=","
					}
				}
				answer+="]\n"
			}
			
		}
	}

	answer += "\tcases = [["
	for i := 0; i < len(r.cases); i++ {
		for j := 0; j < len(r.cases[i]); j++ {
			if isNotAList(r.paramTypes[j]) { //if regular
				answer += r.cases[i][j][0]
			} else {
				answer += "a" + strconv.Itoa(i) + "_" + strconv.Itoa(j)
			}
			if j+1 != len(r.cases[i]) {
				answer += ","
			}
		}
		answer += "]"
		if i+1 != len(r.cases) {
			answer += ",["
		}
	}
	answer += "]\n"
	if isNotAList(r.returnType){
		answer+="\tsimple_return = True\n"
	}else {
		answer+="\tsimple_return = False\n"
	}
	answer+="\tfor index, case in enumerate(cases):\n"
	answer+="\t\ttry:\n"
	answer+="\t\t\tresult = "
	answer+=r.methodName+"("
	for i:=0; i<r.numParams; i++{
		answer+="case["+strconv.Itoa(i)+"]"
		if(i+1!=r.numParams){
			answer+=","
		}
	}
	answer+=")\n"
	answer+="\t\t\tif(simple_return):\n"
	answer+="\t\t\t\tif result == expected_results[index]:\n"
	answer+="\t\t\t\t\tmagic('AC')\n"
	answer+="\t\t\t\telse:\n"
	answer+="\t\t\t\t\tmagic('WA')\n"
	answer+="\t\t\telse:\n"
	answer+="\t\t\t\tfailed=len(expected_results[index])!=len(result)\n"
	answer+="\t\t\t\tfor i in range(len(expected_results[index])):\n"
	answer+="\t\t\t\t\tif failed:\n"
	answer+="\t\t\t\t\t\tbreak\n"
	answer+="\t\t\t\t\tfailed = expected_results[index][i]!=result[i]\n"
	answer+="\t\t\t\tif failed:\n"
	answer+="\t\t\t\t\tmagic('WA')\n"
	answer+="\t\t\t\telse:\n"
	answer+="\t\t\t\t\tmagic('AC')\n"
	answer+="\t\texcept:\n"
	answer+="\t\t\tmagic('RE')\n"
	answer+="\nmain()"
	return answer
}

func generateC(userInput, randomNumber string, r reflectionData) string {
	answer := "#include<vector>\n#include<string>\n#include<iostream>\nusing namespace std;\n#include <tuple>\n\n"
	answer+=userInput
	answer += "\n\n"
	//random number print method Follows:user_output \nmagic_number\n result \nmagic_number\n user_output...
	answer+="void magic(string thingToPrint){\n"
	answer+="\tcout<<\"\\n\"<<"+randomNumber+"<<\"\\n\"<<thingToPrint<<\"\\n\"<<"+randomNumber+"<<\"\\n\";}\n\n"
	answer += "int main(){\n"
	//results array
	if isNotAList(r.returnType) {
		answer += "\tvector<"+pythonToC[r.returnType]+"> expected_results = {{"
		for i := 0; i < len(r.expectedResults); i++ {
			answer += r.expectedResults[i]
			if i+1 != len(r.expectedResults) {
				answer += "},{"
			}else{
				answer+="}"
			}
		}
	} else {//forming expected results as list of lists
		answer += "\tvector<"+pythonToC[r.returnType]+"> expected_results = {"
		currentIndex :=0
		for j:=0; j<r.numCases;j++{
			answer+="{"
			lengthOfCurrentTestCase,_:=strconv.Atoi(r.expectedResults[currentIndex])//get length of this test cases's array
			currentIndex++
			for i := 0; i < lengthOfCurrentTestCase; i++ {
				answer += r.expectedResults[currentIndex]
				currentIndex++
				if i+1 != lengthOfCurrentTestCase {
					answer += ","
				}
			}
			answer+="}"
			if j+1 != r.numCases{
				answer+=","
			}
		}
	}
	answer += "};\n"

	//parameter array 
	for i := 0; i < len(r.paramTypes); i++ {//i=current_paramter
		if !(isNotAList(r.paramTypes[i])) {
			//initializes lists and such with name: a<test_case_number>_<parameter_number>
			for j:=0; j<r.numCases; j++{ //j==current_case
				answer+="\t"+pythonToC[r.paramTypes[i]]+" a"+strconv.Itoa(j)+"_"+strconv.Itoa(i)+" = {"
				for k:=0; k<len(r.cases[j][i]); k++{
					answer+=r.cases[j][i][k]
					if k+1 != len(r.cases[j][i]){
						answer+=","
					}
				}
				answer+="};\n"
			}
			
		}
	}

	answer += "\tvector<tuple<"
	for i:=0; i<len(r.paramTypes);i++ {
		answer+=pythonToC[r.paramTypes[i]]
		if(i+1!=len(r.paramTypes)){
			answer+=", "
		}
	}
	answer+= ">> cases = {make_tuple("
	for i := 0; i < len(r.cases); i++ {
		for j := 0; j < len(r.cases[i]); j++ {
			if isNotAList(r.paramTypes[j]) { //if regular
				answer += r.cases[i][j][0]
			} else {
				answer += "a" + strconv.Itoa(i) + "_" + strconv.Itoa(j)
			}
			if j+1 != len(r.cases[i]) {
				answer += ","
			}
		}
		answer += ")"
		if i+1 != len(r.cases) {
			answer += ", make_tuple("
		}
	}
	answer += "};\n"
	answer+="\tfor (int index=0; index<cases.size();index++){\n"
	answer+="\t\ttry{\n"
	answer+="\t\t\t"+pythonToC[r.returnType]+" result = "
	answer+=r.methodName+"("
	for i:=0; i<r.numParams; i++{
		answer+="get<"+strconv.Itoa(i)+">(cases[index])"
		if(i+1!=r.numParams){
			answer+=","
		}
	}
	answer+=");\n"
	if isNotAList(r.returnType){
		answer+="\t\t\t\tif (result == expected_results[index])\n"
		answer+="\t\t\t\t\tmagic(\"AC\");\n"
		answer+="\t\t\t\telse\n"
		answer+="\t\t\t\t\tmagic(\"WA\");}\n"
	}else {
		answer+="\t\t\t\tbool failed=expected_results[index].size()!=result.size();\n"
		answer+="\t\t\t\tfor (int i=0; i<expected_results[index].size();i++){\n"
		answer+="\t\t\t\t\tif (failed)\n"
		answer+="\t\t\t\t\t\tbreak;\n"
		answer+="\t\t\t\t\tfailed = expected_results[index][i]!=result[i];}\n"
		answer+="\t\t\t\tif (failed)\n"
		answer+="\t\t\t\t\tmagic(\"WA\");\n"
		answer+="\t\t\t\telse\n"
		answer+="\t\t\t\t\tmagic(\"AC\");}\n"
	}
	answer+="\t\tcatch(...){\n"
	answer+="\t\t\tmagic(\"RE\");}}}"
	return answer
}

func generateJavacript(userInput, randomNumber string, r reflectionData) string {
	answer := userInput
	answer += "\n\n"
	//random number print method Follows:user_output \nmagic_number\n result \nmagic_number\n user_output...
	answer+="function magic(thingToPrint){\n"
	answer+="\tconsole.log(\"\\n\"+"+randomNumber+"+\"\\n\"+thingToPrint+\"\\n\"+"+randomNumber+"+\"\\n\");}\n\n"
	answer += "function main(){\n"
	//results array
	if isNotAList(r.returnType) {
		answer += "\tlet expected_results = ["
		for i := 0; i < len(r.expectedResults); i++ {
			answer += r.expectedResults[i]
			if i+1 != len(r.expectedResults) {
				answer += ","
			}
		}
	} else {//forming expected results as list of lists
		answer += "\tlet expected_results = ["
		currentIndex :=0
		for j:=0; j<r.numCases;j++{
			answer+="["
			lengthOfCurrentTestCase,_:=strconv.Atoi(r.expectedResults[currentIndex])//get length of this test cases's array
			currentIndex++
			for i := 0; i < lengthOfCurrentTestCase; i++ {
				answer += r.expectedResults[currentIndex]
				currentIndex++
				if i+1 != lengthOfCurrentTestCase {
					answer += ","
				}
			}
			answer+="]"
			if j+1 != r.numCases{
				answer+=","
			}
		}
	}
	answer += "];\n"

	//parameter array
	for i := 0; i < len(r.paramTypes); i++ {//i=current_paramter
		if !(isNotAList(r.paramTypes[i])) {
			//initializes lists and such with name: a<test_case_number>_<parameter_number>
			for j:=0; j<r.numCases; j++{ //j==current_Case
				answer+="\tlet a"+strconv.Itoa(j)+"_"+strconv.Itoa(i)+" = ["
				for k:=0; k<len(r.cases[j][i]); k++{
					answer+=r.cases[j][i][k]
					if k+1 != len(r.cases[j][i]){
						answer+=","
					}
				}
				answer+="];\n"
			}
			
		}
	}

	answer += "\tlet cases = [["
	for i := 0; i < len(r.cases); i++ {
		for j := 0; j < len(r.cases[i]); j++ {
			if isNotAList(r.paramTypes[j]) { //if regular
				answer += r.cases[i][j][0]
			} else {
				answer += "a" + strconv.Itoa(i) + "_" + strconv.Itoa(j)
			}
			if j+1 != len(r.cases[i]) {
				answer += ","
			}
		}
		answer += "]"
		if i+1 != len(r.cases) {
			answer += ",["
		}
	}
	answer += "];\n"
	if isNotAList(r.returnType){
		answer+="\tlet simple_return = true;\n"
	}else {
		answer+="\tlet simple_return = false;\n"
	}
	answer+="\tfor (let index=0; index<cases.length; index++){\n"
	answer+="\t\ttry{\n"
	answer+="\t\t\tlet result = "
	answer+=r.methodName+"("
	for i:=0; i<r.numParams; i++{
		answer+="cases[index]["+strconv.Itoa(i)+"]"
		if(i+1!=r.numParams){
			answer+=","
		}
	}
	answer+=")\n"
	answer+="\t\t\tif(simple_return)\n"
	answer+="\t\t\t\tif (result == expected_results[index])\n"
	answer+="\t\t\t\t\tmagic('AC');\n"
	answer+="\t\t\t\telse\n"
	answer+="\t\t\t\t\tmagic('WA');\n"
	answer+="\t\t\telse{\n"
	answer+="\t\t\t\tlet failed=expected_results[index].length!=result.length;\n"
	answer+="\t\t\t\tfor (let i=0; i<expected_results[index].length;i++){\n"
	answer+="\t\t\t\t\tif (failed)\n"
	answer+="\t\t\t\t\t\tbreak;\n"
	answer+="\t\t\t\t\tfailed = expected_results[index][i]!=result[i];}\n"
	answer+="\t\t\t\tif (failed)\n"
	answer+="\t\t\t\t\tmagic('WA');\n"
	answer+="\t\t\t\telse\n"
	answer+="\t\t\t\t\tmagic('AC');}}\n"
	answer+="\t\tcatch(error){\n"
	answer+="\t\t\tmagic('RE');}}}\n"
	answer+="\nmain();"
	return answer
}