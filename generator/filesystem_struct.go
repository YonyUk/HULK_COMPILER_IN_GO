package generator

import (
	"os"
	"path"
	"strings"

	. "hulk.com/app/generator/ast_package"
	. "hulk.com/app/generator/automaton_package"
	. "hulk.com/app/generator/context_package"
	. "hulk.com/app/generator/error_collector_package"
	. "hulk.com/app/generator/filesystem_package"
	. "hulk.com/app/generator/grammar_package"
	. "hulk.com/app/generator/interpreter_package"
	. "hulk.com/app/generator/lexer_package"
	. "hulk.com/app/generator/lexical_analisys_package"
	. "hulk.com/app/generator/parser_package"
	. "hulk.com/app/generator/regex_package"
	. "hulk.com/app/generator/tokens_package"
	. "hulk.com/app/generator/tools_package"
)

type NewProjectInput struct {
	TokenTypes []string
	Name       string
}

func BuildProject(folder_path string, input NewProjectInput) {
	os.MkdirAll(folder_path, os.ModePerm)
	makeTokensFolder(path.Join(folder_path, "tokens"), input)
	makeToolsFolder(path.Join(folder_path, "tools"), input)
	makeRegexFolder(path.Join(folder_path, "regex"), input)
	makeParserFolder(path.Join(folder_path, "parser"), input)
	makeLexicalAnalisysFolder(path.Join(folder_path, "lexical_analisys"), input)
	makeLexerFolder(path.Join(folder_path, "lexer"), input)
	makeInterpreterFolder(path.Join(folder_path, "interpreter"), input)
	makeGrammarFolder(path.Join(folder_path, "grammar"), input)
	makeFileSystemFolder(path.Join(folder_path, "filesystem"), input)
	makeContextFolder(path.Join(folder_path, "context"), input)
	makeErrorCollectorFolder(path.Join(folder_path, "compiler"), input) // Hay que cambiar nombres en demasiadas carpetas
	// dicha tarea se hara mas tarde
	makeAutomatonFolder(path.Join(folder_path, "automaton"), input)
	makeASTFolder(path.Join(folder_path, "ast"), input)
}

func makeASTFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"IAST", "AtomicAST", "BaseAST"}
	content_by_file := []string{IASTCode(), AtomicASTCode(), BaseASTCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeAutomatonFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{
		"Automaton",
		"AutomatonsBuilders",
		"AutomatonsOperations",
		"AutomatonTools",
		"IAutomaton",
		"IState",
		"State",
	}
	content_by_file := []string{
		AutomatonCode(),
		AutomatonsBuildersCode(),
		AutomatonsOperationsCode(),
		AutomatonToolsCode(),
		IAutomatonCode(),
		IStateCode(),
		StateCode(),
	}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeErrorCollectorFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"Error", "ErrorCollector", "ErrorType", "IError", "IErrorCollector"}
	content_by_file := []string{ErrorCode(), ErrorCollectorCode(), ErrorTypeCode(), IErrorCode(), IErrorCollectorCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeContextFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"IContext"}
	content_by_file := []string{IContextCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeFileSystemFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"FileReader", "IFileReader"}
	content_by_file := []string{FileReaderCode(), IFileReaderCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeGrammarFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{
		"AttributedGrammar",
		"Grammar",
		"GrammarBuilders",
		"GrammarOperations",
		"GrammarSymbol",
		"GrammarSymbolType",
		"IAttributedGrammar",
		"IGrammar",
		"IGrammarSymbol",
	}
	content_by_file := []string{
		AttributedGrammarCode(),
		GrammarCode(),
		GrammarBuidersCode(),
		GrammarOperationsCode(),
		GrammarSymbolCode(),
		GrammarSymbolTypeCode(),
		IAttributedGrammarCode(),
		IGrammarCode(),
		IGrammarSymbolCode(),
	}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeInterpreterFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"IInterpreter", "Interpreter"}
	content_by_file := []string{IInterpreterCode(), InterpreterCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeLexerFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"ILexer", "Lexer"}
	content_by_file := []string{ILexerCode(), LexerCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeTokensFolder(folder_path string, input NewProjectInput) {
	token_types := input.TokenTypes
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"IToken", "ITokenExtractor", "Token", "TokenExtractor"}
	content_by_file := []string{ITokenCode(), ITokenExtractorCode(), TokensCode(), TokenExtractorCode()}
	fileinfo, err := os.Stat(folder_path)
	if err != nil {
		if os.IsNotExist(err) {
			panic("The specified path doesn't exists")
		}
		panic(err)
	}

	token_type_file_content_template := `package tokens

type TokenType int

const (
	GarbageToken TokenType = iota
	EndToken
<Types>
)
`
	types := ""
	for _, t := range token_types {
		types += "\t" + t + "Token\n"
	}
	token_type_file_content_template = strings.Replace(token_type_file_content_template, "<Types>", types, -1)
	if fileinfo.IsDir() {
		for i, file := range files {
			file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
			os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
		}
	}
	os.WriteFile(path.Join(folder_path, "TokenType.go"), []byte(token_type_file_content_template), os.ModePerm)
}

func makeToolsFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"arrays", "IGenerator", "StringTools"}
	content_by_file := []string{ArraysCode(), IGeneratorCode(), StringCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeRegexFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"IRegexEngine", "RegexEngine"}
	content_by_file := []string{IRegexEngineCode(), RegexEngineCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeParserFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{
		"ActionStruct",
		"IItemLR0",
		"IItemLR0Collection",
		"IParser",
		"ItemLR0",
		"ItemLR0Collection",
		"parser_tools",
		"ParserAction",
		"ParserSLR",
		"ParserSLRFromAttributed",
		"ReduceStruct",
	}
	content_by_file := []string{
		ActionStructCode(),
		IItemLR0Code(),
		IItemLR0CollectionCode(),
		IParserCode(),
		ItemLR0Code(),
		ItemLR0CollectionCode(),
		ParserToolsCode(),
		ParserActionCode(),
		ParserSLRCode(),
		ParserSLRFromAttributedGrammarCode(),
		ReduceStructCode(),
	}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}

func makeLexicalAnalisysFolder(folder_path string, input NewProjectInput) {
	os.Mkdir(folder_path, os.ModePerm)
	files := []string{"ILexicalAnalizer", "ILexicalRule", "LexicalAnalizer", "LexicalRule"}
	content_by_file := []string{ILexicalAnalizerCode(), ILexicalRuleCode(), LexicalAnalizerCode(), LexicalRuleCode()}
	for i, file := range files {
		file_content := strings.Replace(content_by_file[i], "hulk.com/app", input.Name+".com/app", -1)
		os.WriteFile(path.Join(folder_path, file+".go"), []byte(file_content), os.ModePerm)
	}
}
