package tvm

const (
	TOK_INCLUDE string = "%include"
	TOK_DEFINE  string = "%define"
)

// tvm opcode map
type opcode int

const (
	NOP opcode = iota
	INT
	MOV

	PUSH
	POP
	PUSHF
	POPF

	INC
	DEC
	ADD
	SUB
	MUL
	DIV
	MOD
	REM

	NOT
	XOR
	OR
	AND
	SHL
	SHR

	CMP
	JMP
	CALL
	RET

	JE
	JNE
	JG
	JGE
	JL
	JLE

	PRN
)

var opcodeMap map[string]opcode = map[string]opcode{
	"nop": NOP,
	"int": INT,
	"mov": MOV,

	"push":  PUSH,
	"pop":   POP,
	"pushf": PUSHF,
	"popf":  POPF,

	"inc": INC,
	"dec": DEC,
	"add": ADD,
	"sub": SUB,
	"mul": MUL,
	"div": DIV,
	"mod": MOD,
	"rem": REM,

	"not": NOT,
	"xor": XOR,
	"or":  OR,
	"and": AND,
	"shl": SHL,
	"shr": SHR,

	"cmp":  CMP,
	"jmp":  JMP,
	"call": CALL,
	"ret":  RET,

	"je":  JE,
	"jne": JNE,
	"jg":  JG,
	"jge": JGE,
	"jl":  JL,
	"jle": JLE,

	"prn": PRN,
}

/*
static const char *tvm_opcode_map[] = {
	"nop", "int", "mov",
	"push", "pop", "pushf", "popf",
	"inc", "dec", "add", "sub", "mul", "div", "mod", "rem",
	"not", "xor", "or", "and", "shl", "shr",
	"cmp", "jmp", "call", "ret",
	"je", "jne", "jg", "jge", "jl", "jle",
	"prn", 0
};
*/

type register int

const (
	EAX register = iota
	EBX
	ECX
	EDX

	ESI
	EDI
	ESP
	EBP

	EIP

	R08
	R09
	R10
	R11

	R12
	R13
	R14
	R15
)

var registerMap map[string]register = map[string]register{
	"eax": EAX,
	"ebx": EBX,
	"ecx": ECX,
	"edx": EDX,

	"esi": ESI,
	"edi": EDI,
	"esp": ESP,
	"ebp": EBP,

	"eip": EIP,

	"r08": R08,
	"r09": R09,
	"r10": R10,
	"r11": R11,

	"r12": R12,
	"r13": R13,
	"r14": R14,
	"r15": R15,
}

/*
static const char *tvm_register_map[] = {
	"eax", "ebx", "ecx", "edx",
	"esi", "edi", "esp", "ebp",
	"eip",
	"r08", "r09", "r10", "r11",
	"r12", "r13", "r14", "r15", 0};
*/
