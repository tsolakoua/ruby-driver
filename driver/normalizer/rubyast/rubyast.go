// Package rubyast defines constants for Ruby AST.
package rubyast

import "gopkg.in/bblfsh/sdk.v1/uast/ann"

var (
	Alias            = ann.HasInternalType("alias")
	And              = ann.HasInternalType("and")
	AndAsgn          = ann.HasInternalType("and_asgn")
	Arg              = ann.HasInternalType("arg")
	ArgExpr          = ann.HasInternalType("arg_expr")
	Args             = ann.HasInternalRole("args")
	Array            = ann.HasInternalType("array")
	Base             = ann.HasInternalType("base")
	Begin            = ann.HasInternalType("begin")
	Block            = ann.HasInternalType("block")
	BlockArg         = ann.HasInternalType("blockarg")
	BlockData        = ann.HasInternalType("blockdata")
	BlockPass        = ann.HasInternalType("block_pass")
	Body             = ann.HasInternalType("body")
	BodyRole         = ann.HasInternalRole("body")
	Break            = ann.HasInternalType("break")
	CAsgn            = ann.HasInternalType("casgn")
	CBase            = ann.HasInternalType("cbase")
	CSend            = ann.HasInternalType("csend")
	CVAsgn           = ann.HasInternalType("cvasgn")
	CVar             = ann.HasInternalType("cvar")
	Case             = ann.HasInternalType("case")
	CaseVar          = ann.HasInternalType("casevar")
	Children         = ann.HasInternalType("children")
	Class            = ann.HasInternalType("class")
	Complex          = ann.HasInternalType("complex")
	Condition        = ann.HasInternalType("condition")
	Conditions       = ann.HasInternalType("conditions")
	Const            = ann.HasInternalType("const")
	Contents         = ann.HasInternalType("contents")
	DStr             = ann.HasInternalType("dstr")
	DSym             = ann.HasInternalType("dsym")
	Def              = ann.HasInternalType("def")
	Default          = ann.HasInternalType("default")
	Defined          = ann.HasInternalType("defined?")
	Defs             = ann.HasInternalType("defs")
	Documentation    = ann.HasInternalType("documentation")
	EFlipFlop        = ann.HasInternalType("eflipflop")
	ERange           = ann.HasInternalType("erange")
	Else             = ann.HasInternalType("else")
	Ensure           = ann.HasInternalType("ensure")
	Ensurebody       = ann.HasInternalType("ensure_body")
	Exceptions       = ann.HasInternalType("exceptions")
	Expression       = ann.HasInternalType("expression")
	False            = ann.HasInternalType("false")
	Float            = ann.HasInternalType("float")
	For              = ann.HasInternalType("for")
	GVAsgn           = ann.HasInternalType("gvasgn")
	GVar             = ann.HasInternalType("gvar")
	Handlers         = ann.HasInternalType("handlers")
	Hash             = ann.HasInternalType("hash")
	IFlipFlop        = ann.HasInternalType("iflipflop")
	IRange           = ann.HasInternalType("irange")
	IVAsgn           = ann.HasInternalType("ivasgn")
	IVar             = ann.HasInternalType("ivar")
	If               = ann.HasInternalType("if")
	Inline           = ann.HasInternalType("inline")
	InnerComplex     = ann.HasInternalType("Complex")
	InnerRational    = ann.HasInternalType("Rational")
	InnerSymbol      = ann.HasInternalType("Symbol")
	Int              = ann.HasInternalType("int")
	Iterated         = ann.HasInternalType("iterated")
	Iterators        = ann.HasInternalType("iterators")
	KwArg            = ann.HasInternalType("kwarg")
	KwBegin          = ann.HasInternalType("kwbegin")
	KwOptArg         = ann.HasInternalType("kwoptarg")
	KwRestArg        = ann.HasInternalType("kwrestarg")
	KwSplat          = ann.HasInternalType("kwsplat")
	LVAsgn           = ann.HasInternalType("lvasgn")
	LVar             = ann.HasInternalType("lvar")
	Lambda           = ann.HasInternalType("lambda")
	MAsgn            = ann.HasInternalType("masgn")
	MatchCurrentLine = ann.HasInternalType("match_current_line")
	MatchWithLVAsgn  = ann.HasInternalType("match_with_lvasgn")
	MultipleLeftSide = ann.HasInternalType("mlhs")
	Module           = ann.HasInternalType("module")
	Name             = ann.HasInternalType("name")
	Next             = ann.HasInternalType("next")
	Nil              = ann.HasInternalType("nil")
	NilClass         = ann.HasInternalType("NilClass")
	Not              = ann.HasInternalType("not")
	Object           = ann.HasInternalType("object")
	OpAsgn           = ann.HasInternalType("op_asgn")
	Operator         = ann.HasInternalType("operator")
	OptArg           = ann.HasInternalType("optarg")
	Options          = ann.HasInternalRole("options")
	Or               = ann.HasInternalType("or")
	OrAsgn           = ann.HasInternalType("or_asgn")
	Pair             = ann.HasInternalType("pair")
	PairFirst        = ann.HasInternalRole("_1")
	PairSecond       = ann.HasInternalRole("_2")
	Parent           = ann.HasInternalType("parent")
	PostExe          = ann.HasInternalType("postexe")
	PreExe           = ann.HasInternalType("preexe")
	ProcArg0         = ann.HasInternalType("procarg0")
	Question         = ann.HasInternalType("question")
	Rational         = ann.HasInternalType("rational")
	Redo             = ann.HasInternalType("redo")
	RegExp           = ann.HasInternalType("regexp")
	RegExpBackRef    = ann.HasInternalType("back_ref")
	RegExpRef        = ann.HasInternalType("nth_ref")
	RegOpt           = ann.HasInternalType("regopt")
	ResBody          = ann.HasInternalType("resbody")
	Rescue           = ann.HasInternalType("rescue")
	RestArg          = ann.HasInternalType("restarg")
	Retry            = ann.HasInternalType("retry")
	Return           = ann.HasInternalType("return")
	SClass           = ann.HasInternalType("sclass")
	SName            = ann.HasInternalType("s_name")
	Selector         = ann.HasInternalType("selector")
	Self             = ann.HasInternalType("self")
	Send             = ann.HasInternalType("send")
	ShadowArg        = ann.HasInternalType("shadow_arg")
	Splat            = ann.HasInternalType("splat")
	Str              = ann.HasInternalType("str")
	Symbol           = ann.HasInternalType("Symbol")
	Super            = ann.HasInternalType("super")
	Sym              = ann.HasInternalType("sym")
	Target           = ann.HasInternalType("target")
	Text             = ann.HasInternalType("text")
	True             = ann.HasInternalType("true")
	Undef            = ann.HasInternalType("undef")
	Until            = ann.HasInternalType("until")
	UntilPost        = ann.HasInternalType("until_post")
	Value            = ann.HasInternalType("value")
	Values           = ann.HasInternalRole("values")
	When             = ann.HasInternalType("when")
	WhenClauses      = ann.HasInternalType("when_clauses")
	While            = ann.HasInternalType("while")
	WhilePost        = ann.HasInternalType("while_post")
	XStr             = ann.HasInternalType("xstr")
	Yield            = ann.HasInternalType("yield")
	ZSuper           = ann.HasInternalType("zsuper")
)