"".array_def STEXT size=1809 args=0x0 locals=0x260 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:15)	TEXT	"".array_def(SB), ABIInternal, $608-0
	0x0000 00000 (06-data_types/07-array.go:15)	LEAQ	-480(SP), R12
	0x0008 00008 (06-data_types/07-array.go:15)	CMPQ	R12, 16(R14)
	0x000c 00012 (06-data_types/07-array.go:15)	PCDATA	$0, $-2
	0x000c 00012 (06-data_types/07-array.go:15)	JLS	1799
	0x0012 00018 (06-data_types/07-array.go:15)	PCDATA	$0, $-1
	0x0012 00018 (06-data_types/07-array.go:15)	SUBQ	$608, SP
	0x0019 00025 (06-data_types/07-array.go:15)	MOVQ	BP, 600(SP)
	0x0021 00033 (06-data_types/07-array.go:15)	LEAQ	600(SP), BP
	0x0029 00041 (06-data_types/07-array.go:15)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0029 00041 (06-data_types/07-array.go:15)	FUNCDATA	$1, gclocals·9f1f24ff36424f3d0e5fd3a9c6a7b1cf(SB)
	0x0029 00041 (06-data_types/07-array.go:15)	FUNCDATA	$2, "".array_def.stkobj(SB)
	0x0029 00041 (06-data_types/07-array.go:17)	MOVQ	$0, "".a+144(SP)
	0x0035 00053 (06-data_types/07-array.go:17)	LEAQ	"".a+152(SP), CX
	0x003d 00061 (06-data_types/07-array.go:17)	MOVUPS	X15, (CX)
	0x0041 00065 (06-data_types/07-array.go:18)	MOVQ	$0, ""..autotmp_67+192(SP)
	0x004d 00077 (06-data_types/07-array.go:18)	LEAQ	""..autotmp_67+200(SP), CX
	0x0055 00085 (06-data_types/07-array.go:18)	MOVUPS	X15, (CX)
	0x0059 00089 (06-data_types/07-array.go:18)	LEAQ	type.[3]int(SB), AX
	0x0060 00096 (06-data_types/07-array.go:18)	LEAQ	""..autotmp_67+192(SP), BX
	0x0068 00104 (06-data_types/07-array.go:18)	PCDATA	$1, $0
	0x0068 00104 (06-data_types/07-array.go:18)	CALL	runtime.convT2Enoptr(SB)
	0x006d 00109 (06-data_types/07-array.go:18)	MOVUPS	X15, ""..autotmp_77+520(SP)
	0x0076 00118 (06-data_types/07-array.go:18)	MOVQ	AX, ""..autotmp_77+520(SP)
	0x007e 00126 (06-data_types/07-array.go:18)	MOVQ	BX, ""..autotmp_77+528(SP)
	0x0086 00134 (<unknown line number>)	NOP
	0x0086 00134 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x008d 00141 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0094 00148 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a: %v\n"(SB), CX
	0x009b 00155 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x00a0 00160 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_77+520(SP), SI
	0x00a8 00168 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x00ae 00174 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x00b1 00177 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x00b6 00182 (06-data_types/07-array.go:21)	MOVQ	$0, "".b+120(SP)
	0x00bf 00191 (06-data_types/07-array.go:21)	LEAQ	"".b+128(SP), CX
	0x00c7 00199 (06-data_types/07-array.go:21)	MOVUPS	X15, (CX)
	0x00cb 00203 (06-data_types/07-array.go:21)	MOVQ	$1, "".b+120(SP)
	0x00d4 00212 (06-data_types/07-array.go:21)	MOVQ	$2, "".b+128(SP)
	0x00e0 00224 (06-data_types/07-array.go:21)	MOVQ	$3, "".b+136(SP)
	0x00ec 00236 (06-data_types/07-array.go:22)	MOVQ	$1, ""..autotmp_67+192(SP)
	0x00f8 00248 (06-data_types/07-array.go:22)	MOVQ	$2, ""..autotmp_67+200(SP)
	0x0104 00260 (06-data_types/07-array.go:22)	MOVQ	$3, ""..autotmp_67+208(SP)
	0x0110 00272 (06-data_types/07-array.go:22)	LEAQ	type.[3]int(SB), AX
	0x0117 00279 (06-data_types/07-array.go:22)	LEAQ	""..autotmp_67+192(SP), BX
	0x011f 00287 (06-data_types/07-array.go:22)	NOP
	0x0120 00288 (06-data_types/07-array.go:22)	CALL	runtime.convT2Enoptr(SB)
	0x0125 00293 (06-data_types/07-array.go:22)	MOVUPS	X15, ""..autotmp_79+504(SP)
	0x012e 00302 (06-data_types/07-array.go:22)	MOVQ	AX, ""..autotmp_79+504(SP)
	0x0136 00310 (06-data_types/07-array.go:22)	MOVQ	BX, ""..autotmp_79+512(SP)
	0x013e 00318 (<unknown line number>)	NOP
	0x013e 00318 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0145 00325 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x014c 00332 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."b: %v\n"(SB), CX
	0x0153 00339 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0158 00344 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_79+504(SP), SI
	0x0160 00352 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0166 00358 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0169 00361 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x016e 00366 (06-data_types/07-array.go:25)	MOVQ	$0, "".c+96(SP)
	0x0177 00375 (06-data_types/07-array.go:25)	LEAQ	"".c+104(SP), CX
	0x017c 00380 (06-data_types/07-array.go:25)	MOVUPS	X15, (CX)
	0x0180 00384 (06-data_types/07-array.go:25)	MOVQ	$1, "".c+96(SP)
	0x0189 00393 (06-data_types/07-array.go:25)	MOVQ	$2, "".c+104(SP)
	0x0192 00402 (06-data_types/07-array.go:26)	MOVQ	$0, ""..autotmp_67+192(SP)
	0x019e 00414 (06-data_types/07-array.go:26)	LEAQ	""..autotmp_67+200(SP), CX
	0x01a6 00422 (06-data_types/07-array.go:26)	MOVUPS	X15, (CX)
	0x01aa 00426 (06-data_types/07-array.go:26)	MOVQ	$1, ""..autotmp_67+192(SP)
	0x01b6 00438 (06-data_types/07-array.go:26)	MOVQ	$2, ""..autotmp_67+200(SP)
	0x01c2 00450 (06-data_types/07-array.go:26)	LEAQ	type.[3]int(SB), AX
	0x01c9 00457 (06-data_types/07-array.go:26)	LEAQ	""..autotmp_67+192(SP), BX
	0x01d1 00465 (06-data_types/07-array.go:26)	CALL	runtime.convT2Enoptr(SB)
	0x01d6 00470 (06-data_types/07-array.go:26)	MOVUPS	X15, ""..autotmp_81+488(SP)
	0x01df 00479 (06-data_types/07-array.go:26)	MOVQ	AX, ""..autotmp_81+488(SP)
	0x01e7 00487 (06-data_types/07-array.go:26)	MOVQ	BX, ""..autotmp_81+496(SP)
	0x01ef 00495 (<unknown line number>)	NOP
	0x01ef 00495 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x01f6 00502 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x01fd 00509 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."c: %v\n"(SB), CX
	0x0204 00516 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0209 00521 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_81+488(SP), SI
	0x0211 00529 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0217 00535 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x021a 00538 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x021f 00543 (06-data_types/07-array.go:30)	MOVQ	$1, ""..autotmp_67+192(SP)
	0x022b 00555 (06-data_types/07-array.go:30)	MOVQ	$2, ""..autotmp_67+200(SP)
	0x0237 00567 (06-data_types/07-array.go:30)	MOVQ	$3, ""..autotmp_67+208(SP)
	0x0243 00579 (06-data_types/07-array.go:30)	LEAQ	type.[3]int(SB), AX
	0x024a 00586 (06-data_types/07-array.go:30)	LEAQ	""..autotmp_67+192(SP), BX
	0x0252 00594 (06-data_types/07-array.go:30)	CALL	runtime.convT2Enoptr(SB)
	0x0257 00599 (06-data_types/07-array.go:30)	MOVUPS	X15, ""..autotmp_83+472(SP)
	0x0260 00608 (06-data_types/07-array.go:30)	MOVQ	AX, ""..autotmp_83+472(SP)
	0x0268 00616 (06-data_types/07-array.go:30)	MOVQ	BX, ""..autotmp_83+480(SP)
	0x0270 00624 (<unknown line number>)	NOP
	0x0270 00624 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0277 00631 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x027e 00638 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."%T, d: %[1]v\n"(SB), CX
	0x0285 00645 ($GOROOT/src/fmt/print.go:213)	MOVL	$13, DI
	0x028a 00650 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_83+472(SP), SI
	0x0292 00658 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0298 00664 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x029b 00667 ($GOROOT/src/fmt/print.go:213)	NOP
	0x02a0 00672 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x02a5 00677 (06-data_types/07-array.go:33)	MOVQ	$0, "".e+72(SP)
	0x02ae 00686 (06-data_types/07-array.go:33)	LEAQ	"".e+80(SP), CX
	0x02b3 00691 (06-data_types/07-array.go:33)	MOVUPS	X15, (CX)
	0x02b7 00695 (06-data_types/07-array.go:33)	MOVQ	$1, "".e+72(SP)
	0x02c0 00704 (06-data_types/07-array.go:33)	MOVQ	$2, "".e+80(SP)
	0x02c9 00713 (06-data_types/07-array.go:33)	MOVQ	$3, "".e+88(SP)
	0x02d2 00722 (06-data_types/07-array.go:35)	MOVQ	"".a+144(SP), CX
	0x02da 00730 (06-data_types/07-array.go:35)	MOVQ	CX, "".e+72(SP)
	0x02df 00735 (06-data_types/07-array.go:35)	MOVUPS	"".a+152(SP), X0
	0x02e7 00743 (06-data_types/07-array.go:35)	MOVUPS	X0, "".e+80(SP)
	0x02ec 00748 (06-data_types/07-array.go:38)	MOVQ	$200, "".e+80(SP)
	0x02f5 00757 (06-data_types/07-array.go:39)	MOVQ	"".e+72(SP), CX
	0x02fa 00762 (06-data_types/07-array.go:39)	MOVQ	CX, ""..autotmp_67+192(SP)
	0x0302 00770 (06-data_types/07-array.go:39)	MOVUPS	"".e+80(SP), X0
	0x0307 00775 (06-data_types/07-array.go:39)	MOVUPS	X0, ""..autotmp_67+200(SP)
	0x030f 00783 (06-data_types/07-array.go:39)	LEAQ	type.[3]int(SB), AX
	0x0316 00790 (06-data_types/07-array.go:39)	LEAQ	""..autotmp_67+192(SP), BX
	0x031e 00798 (06-data_types/07-array.go:39)	NOP
	0x0320 00800 (06-data_types/07-array.go:39)	CALL	runtime.convT2Enoptr(SB)
	0x0325 00805 (06-data_types/07-array.go:39)	MOVUPS	X15, ""..autotmp_85+456(SP)
	0x032e 00814 (06-data_types/07-array.go:39)	MOVQ	AX, ""..autotmp_85+456(SP)
	0x0336 00822 (06-data_types/07-array.go:39)	MOVQ	BX, ""..autotmp_85+464(SP)
	0x033e 00830 (<unknown line number>)	NOP
	0x033e 00830 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0345 00837 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x034c 00844 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."e: %v\n"(SB), CX
	0x0353 00851 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0358 00856 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_85+456(SP), SI
	0x0360 00864 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0366 00870 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0369 00873 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x036e 00878 (06-data_types/07-array.go:42)	MOVQ	"".a+144(SP), CX
	0x0376 00886 (06-data_types/07-array.go:42)	MOVQ	CX, ""..autotmp_67+192(SP)
	0x037e 00894 (06-data_types/07-array.go:42)	MOVUPS	"".a+152(SP), X0
	0x0386 00902 (06-data_types/07-array.go:42)	MOVUPS	X0, ""..autotmp_67+200(SP)
	0x038e 00910 (06-data_types/07-array.go:42)	XORL	AX, AX
	0x0390 00912 (06-data_types/07-array.go:42)	JMP	1088
	0x0395 00917 (06-data_types/07-array.go:42)	MOVQ	AX, "".idx+64(SP)
	0x039a 00922 (06-data_types/07-array.go:42)	MOVQ	""..autotmp_67+192(SP)(AX*8), CX
	0x03a2 00930 (06-data_types/07-array.go:42)	MOVQ	CX, "".v+56(SP)
	0x03a7 00935 (06-data_types/07-array.go:43)	CALL	runtime.convT64(SB)
	0x03ac 00940 (06-data_types/07-array.go:43)	MOVQ	AX, ""..autotmp_154+432(SP)
	0x03b4 00948 (06-data_types/07-array.go:43)	MOVQ	"".v+56(SP), AX
	0x03b9 00953 (06-data_types/07-array.go:43)	PCDATA	$1, $1
	0x03b9 00953 (06-data_types/07-array.go:43)	CALL	runtime.convT64(SB)
	0x03be 00958 (06-data_types/07-array.go:43)	LEAQ	""..autotmp_91+568(SP), SI
	0x03c6 00966 (06-data_types/07-array.go:43)	MOVUPS	X15, (SI)
	0x03ca 00970 (06-data_types/07-array.go:43)	LEAQ	""..autotmp_91+584(SP), CX
	0x03d2 00978 (06-data_types/07-array.go:43)	MOVUPS	X15, (CX)
	0x03d6 00982 (06-data_types/07-array.go:43)	LEAQ	type.int(SB), CX
	0x03dd 00989 (06-data_types/07-array.go:43)	MOVQ	CX, ""..autotmp_91+568(SP)
	0x03e5 00997 (06-data_types/07-array.go:43)	MOVQ	""..autotmp_154+432(SP), DX
	0x03ed 01005 (06-data_types/07-array.go:43)	MOVQ	DX, ""..autotmp_91+576(SP)
	0x03f5 01013 (06-data_types/07-array.go:43)	MOVQ	CX, ""..autotmp_91+584(SP)
	0x03fd 01021 (06-data_types/07-array.go:43)	MOVQ	AX, ""..autotmp_91+592(SP)
	0x0405 01029 (<unknown line number>)	NOP
	0x0405 01029 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x040c 01036 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0413 01043 ($GOROOT/src/fmt/print.go:213)	MOVL	$20, DI
	0x0418 01048 ($GOROOT/src/fmt/print.go:213)	MOVL	$2, R8
	0x041e 01054 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0421 01057 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."index: %d\tvalue: %d\n"(SB), CX
	0x0428 01064 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $0
	0x0428 01064 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x042d 01069 (06-data_types/07-array.go:42)	MOVQ	"".idx+64(SP), CX
	0x0432 01074 (06-data_types/07-array.go:42)	LEAQ	1(CX), AX
	0x0436 01078 (06-data_types/07-array.go:42)	NOP
	0x0440 01088 (06-data_types/07-array.go:42)	CMPQ	AX, $3
	0x0444 01092 (06-data_types/07-array.go:42)	JLT	917
	0x044a 01098 (06-data_types/07-array.go:48)	MOVQ	$0, ""..autotmp_70+344(SP)
	0x0456 01110 (06-data_types/07-array.go:48)	LEAQ	""..autotmp_70+352(SP), DI
	0x045e 01118 (06-data_types/07-array.go:48)	PCDATA	$0, $-2
	0x045e 01118 (06-data_types/07-array.go:48)	LEAQ	-48(DI), DI
	0x0462 01122 (06-data_types/07-array.go:48)	DUFFZERO	$336
	0x0475 01141 (06-data_types/07-array.go:48)	PCDATA	$0, $-1
	0x0475 01141 (06-data_types/07-array.go:48)	MOVQ	$25, ""..autotmp_70+384(SP)
	0x0481 01153 (06-data_types/07-array.go:48)	MOVQ	$100, ""..autotmp_70+424(SP)
	0x048d 01165 (06-data_types/07-array.go:48)	LEAQ	type.[11]int(SB), AX
	0x0494 01172 (06-data_types/07-array.go:48)	LEAQ	""..autotmp_70+344(SP), BX
	0x049c 01180 (06-data_types/07-array.go:48)	NOP
	0x04a0 01184 (06-data_types/07-array.go:48)	CALL	runtime.convT2Enoptr(SB)
	0x04a5 01189 (06-data_types/07-array.go:48)	MOVUPS	X15, ""..autotmp_93+440(SP)
	0x04ae 01198 (06-data_types/07-array.go:48)	MOVQ	AX, ""..autotmp_93+440(SP)
	0x04b6 01206 (06-data_types/07-array.go:48)	MOVQ	BX, ""..autotmp_93+448(SP)
	0x04be 01214 (<unknown line number>)	NOP
	0x04be 01214 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x04c5 01221 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x04cc 01228 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."f type: %T\t, value: %[1]v\n"(SB), CX
	0x04d3 01235 ($GOROOT/src/fmt/print.go:213)	MOVL	$26, DI
	0x04d8 01240 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_93+440(SP), SI
	0x04e0 01248 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x04e6 01254 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x04e9 01257 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x04ee 01262 (06-data_types/07-array.go:52)	LEAQ	"".g+248(SP), CX
	0x04f6 01270 (06-data_types/07-array.go:52)	MOVUPS	X15, (CX)
	0x04fa 01274 (06-data_types/07-array.go:52)	LEAQ	"".g+264(SP), CX
	0x0502 01282 (06-data_types/07-array.go:52)	MOVUPS	X15, (CX)
	0x0506 01286 (06-data_types/07-array.go:52)	MOVQ	$1, "".g+248(SP)
	0x0512 01298 (06-data_types/07-array.go:52)	MOVQ	$2, "".g+256(SP)
	0x051e 01310 (06-data_types/07-array.go:52)	MOVQ	$3, "".g+264(SP)
	0x052a 01322 (06-data_types/07-array.go:52)	MOVQ	$4, "".g+272(SP)
	0x0536 01334 (06-data_types/07-array.go:53)	LEAQ	"".h+216(SP), CX
	0x053e 01342 (06-data_types/07-array.go:53)	MOVUPS	X15, (CX)
	0x0542 01346 (06-data_types/07-array.go:53)	LEAQ	"".h+232(SP), CX
	0x054a 01354 (06-data_types/07-array.go:53)	MOVUPS	X15, (CX)
	0x054e 01358 (06-data_types/07-array.go:53)	MOVQ	$1, "".h+216(SP)
	0x055a 01370 (06-data_types/07-array.go:53)	MOVQ	$2, "".h+224(SP)
	0x0566 01382 (06-data_types/07-array.go:53)	MOVQ	$3, "".h+232(SP)
	0x0572 01394 (06-data_types/07-array.go:53)	MOVQ	$4, "".h+240(SP)
	0x057e 01406 (06-data_types/07-array.go:55)	MOVQ	"".b+120(SP), CX
	0x0583 01411 (06-data_types/07-array.go:55)	MOVQ	CX, ""..autotmp_67+192(SP)
	0x058b 01419 (06-data_types/07-array.go:55)	MOVUPS	"".b+128(SP), X0
	0x0593 01427 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_67+200(SP)
	0x059b 01435 (06-data_types/07-array.go:55)	MOVQ	"".c+96(SP), CX
	0x05a0 01440 (06-data_types/07-array.go:55)	MOVQ	CX, ""..autotmp_71+168(SP)
	0x05a8 01448 (06-data_types/07-array.go:55)	MOVUPS	"".c+104(SP), X0
	0x05ad 01453 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_71+176(SP)
	0x05b5 01461 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_67+192(SP), CX
	0x05bd 01469 (06-data_types/07-array.go:55)	NOP
	0x05c0 01472 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_71+168(SP), CX
	0x05c8 01480 (06-data_types/07-array.go:55)	JNE	1521
	0x05ca 01482 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_67+200(SP), DX
	0x05d2 01490 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_71+176(SP), DX
	0x05da 01498 (06-data_types/07-array.go:55)	JNE	1521
	0x05dc 01500 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_67+208(SP), DX
	0x05e4 01508 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_71+184(SP), DX
	0x05ec 01516 (06-data_types/07-array.go:55)	SETEQ	DL
	0x05ef 01519 (06-data_types/07-array.go:55)	JMP	1523
	0x05f1 01521 (06-data_types/07-array.go:55)	XORL	DX, DX
	0x05f3 01523 (06-data_types/07-array.go:55)	MOVUPS	"".g+248(SP), X0
	0x05fb 01531 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_73+312(SP)
	0x0603 01539 (06-data_types/07-array.go:55)	MOVUPS	"".g+264(SP), X0
	0x060b 01547 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_73+328(SP)
	0x0613 01555 (06-data_types/07-array.go:55)	MOVUPS	"".h+216(SP), X0
	0x061b 01563 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_74+280(SP)
	0x0623 01571 (06-data_types/07-array.go:55)	MOVUPS	"".h+232(SP), X0
	0x062b 01579 (06-data_types/07-array.go:55)	MOVUPS	X0, ""..autotmp_74+296(SP)
	0x0633 01587 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_73+312(SP), R8
	0x063b 01595 (06-data_types/07-array.go:55)	NOP
	0x0640 01600 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_74+280(SP), R8
	0x0648 01608 (06-data_types/07-array.go:55)	JNE	1668
	0x064a 01610 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_73+320(SP), R8
	0x0652 01618 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_74+288(SP), R8
	0x065a 01626 (06-data_types/07-array.go:55)	JNE	1668
	0x065c 01628 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_73+328(SP), R8
	0x0664 01636 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_74+296(SP), R8
	0x066c 01644 (06-data_types/07-array.go:55)	JNE	1668
	0x066e 01646 (06-data_types/07-array.go:55)	MOVQ	""..autotmp_73+336(SP), R8
	0x0676 01654 (06-data_types/07-array.go:55)	CMPQ	""..autotmp_74+304(SP), R8
	0x067e 01662 (06-data_types/07-array.go:55)	SETEQ	R8B
	0x0682 01666 (06-data_types/07-array.go:55)	JMP	1671
	0x0684 01668 (06-data_types/07-array.go:55)	XORL	R8, R8
	0x0687 01671 (06-data_types/07-array.go:55)	LEAQ	""..autotmp_95+536(SP), CX
	0x068f 01679 (06-data_types/07-array.go:55)	MOVUPS	X15, (CX)
	0x0693 01683 (06-data_types/07-array.go:55)	LEAQ	""..autotmp_95+552(SP), R9
	0x069b 01691 (06-data_types/07-array.go:55)	MOVUPS	X15, (R9)
	0x069f 01695 (06-data_types/07-array.go:55)	LEAQ	type.bool(SB), R9
	0x06a6 01702 (06-data_types/07-array.go:55)	MOVQ	R9, ""..autotmp_95+536(SP)
	0x06ae 01710 (06-data_types/07-array.go:55)	MOVBLZX	DL, DX
	0x06b1 01713 (06-data_types/07-array.go:55)	LEAQ	runtime.staticuint64s(SB), R10
	0x06b8 01720 (06-data_types/07-array.go:55)	LEAQ	(R10)(DX*8), DX
	0x06bc 01724 (06-data_types/07-array.go:55)	MOVQ	DX, ""..autotmp_95+544(SP)
	0x06c4 01732 (06-data_types/07-array.go:55)	MOVQ	R9, ""..autotmp_95+552(SP)
	0x06cc 01740 (06-data_types/07-array.go:55)	MOVBLZX	R8B, DX
	0x06d0 01744 (06-data_types/07-array.go:55)	LEAQ	(R10)(DX*8), DX
	0x06d4 01748 (06-data_types/07-array.go:55)	MOVQ	DX, ""..autotmp_95+560(SP)
	0x06dc 01756 (<unknown line number>)	NOP
	0x06dc 01756 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x06e3 01763 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x06ea 01770 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x06ef 01775 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x06f2 01778 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x06f7 01783 (06-data_types/07-array.go:56)	MOVQ	600(SP), BP
	0x06ff 01791 (06-data_types/07-array.go:56)	ADDQ	$608, SP
	0x0706 01798 (06-data_types/07-array.go:56)	RET
	0x0707 01799 (06-data_types/07-array.go:56)	NOP
	0x0707 01799 (06-data_types/07-array.go:15)	PCDATA	$1, $-1
	0x0707 01799 (06-data_types/07-array.go:15)	PCDATA	$0, $-2
	0x0707 01799 (06-data_types/07-array.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x070c 01804 (06-data_types/07-array.go:15)	PCDATA	$0, $-1
	0x070c 01804 (06-data_types/07-array.go:15)	JMP	0
	0x0000 4c 8d a4 24 20 fe ff ff 4d 3b 66 10 0f 86 f5 06  L..$ ...M;f.....
	0x0010 00 00 48 81 ec 60 02 00 00 48 89 ac 24 58 02 00  ..H..`...H..$X..
	0x0020 00 48 8d ac 24 58 02 00 00 48 c7 84 24 90 00 00  .H..$X...H..$...
	0x0030 00 00 00 00 00 48 8d 8c 24 98 00 00 00 44 0f 11  .....H..$....D..
	0x0040 39 48 c7 84 24 c0 00 00 00 00 00 00 00 48 8d 8c  9H..$........H..
	0x0050 24 c8 00 00 00 44 0f 11 39 48 8d 05 00 00 00 00  $....D..9H......
	0x0060 48 8d 9c 24 c0 00 00 00 e8 00 00 00 00 44 0f 11  H..$.........D..
	0x0070 bc 24 08 02 00 00 48 89 84 24 08 02 00 00 48 89  .$....H..$....H.
	0x0080 9c 24 10 02 00 00 48 8b 1d 00 00 00 00 48 8d 05  .$....H......H..
	0x0090 00 00 00 00 48 8d 0d 00 00 00 00 bf 06 00 00 00  ....H...........
	0x00a0 48 8d b4 24 08 02 00 00 41 b8 01 00 00 00 4d 89  H..$....A.....M.
	0x00b0 c1 e8 00 00 00 00 48 c7 44 24 78 00 00 00 00 48  ......H.D$x....H
	0x00c0 8d 8c 24 80 00 00 00 44 0f 11 39 48 c7 44 24 78  ..$....D..9H.D$x
	0x00d0 01 00 00 00 48 c7 84 24 80 00 00 00 02 00 00 00  ....H..$........
	0x00e0 48 c7 84 24 88 00 00 00 03 00 00 00 48 c7 84 24  H..$........H..$
	0x00f0 c0 00 00 00 01 00 00 00 48 c7 84 24 c8 00 00 00  ........H..$....
	0x0100 02 00 00 00 48 c7 84 24 d0 00 00 00 03 00 00 00  ....H..$........
	0x0110 48 8d 05 00 00 00 00 48 8d 9c 24 c0 00 00 00 90  H......H..$.....
	0x0120 e8 00 00 00 00 44 0f 11 bc 24 f8 01 00 00 48 89  .....D...$....H.
	0x0130 84 24 f8 01 00 00 48 89 9c 24 00 02 00 00 48 8b  .$....H..$....H.
	0x0140 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x0150 00 00 00 bf 06 00 00 00 48 8d b4 24 f8 01 00 00  ........H..$....
	0x0160 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 c7  A.....M.......H.
	0x0170 44 24 60 00 00 00 00 48 8d 4c 24 68 44 0f 11 39  D$`....H.L$hD..9
	0x0180 48 c7 44 24 60 01 00 00 00 48 c7 44 24 68 02 00  H.D$`....H.D$h..
	0x0190 00 00 48 c7 84 24 c0 00 00 00 00 00 00 00 48 8d  ..H..$........H.
	0x01a0 8c 24 c8 00 00 00 44 0f 11 39 48 c7 84 24 c0 00  .$....D..9H..$..
	0x01b0 00 00 01 00 00 00 48 c7 84 24 c8 00 00 00 02 00  ......H..$......
	0x01c0 00 00 48 8d 05 00 00 00 00 48 8d 9c 24 c0 00 00  ..H......H..$...
	0x01d0 00 e8 00 00 00 00 44 0f 11 bc 24 e8 01 00 00 48  ......D...$....H
	0x01e0 89 84 24 e8 01 00 00 48 89 9c 24 f0 01 00 00 48  ..$....H..$....H
	0x01f0 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d  ......H......H..
	0x0200 00 00 00 00 bf 06 00 00 00 48 8d b4 24 e8 01 00  .........H..$...
	0x0210 00 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48  .A.....M.......H
	0x0220 c7 84 24 c0 00 00 00 01 00 00 00 48 c7 84 24 c8  ..$........H..$.
	0x0230 00 00 00 02 00 00 00 48 c7 84 24 d0 00 00 00 03  .......H..$.....
	0x0240 00 00 00 48 8d 05 00 00 00 00 48 8d 9c 24 c0 00  ...H......H..$..
	0x0250 00 00 e8 00 00 00 00 44 0f 11 bc 24 d8 01 00 00  .......D...$....
	0x0260 48 89 84 24 d8 01 00 00 48 89 9c 24 e0 01 00 00  H..$....H..$....
	0x0270 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x0280 0d 00 00 00 00 bf 0d 00 00 00 48 8d b4 24 d8 01  ..........H..$..
	0x0290 00 00 41 b8 01 00 00 00 4d 89 c1 0f 1f 44 00 00  ..A.....M....D..
	0x02a0 e8 00 00 00 00 48 c7 44 24 48 00 00 00 00 48 8d  .....H.D$H....H.
	0x02b0 4c 24 50 44 0f 11 39 48 c7 44 24 48 01 00 00 00  L$PD..9H.D$H....
	0x02c0 48 c7 44 24 50 02 00 00 00 48 c7 44 24 58 03 00  H.D$P....H.D$X..
	0x02d0 00 00 48 8b 8c 24 90 00 00 00 48 89 4c 24 48 0f  ..H..$....H.L$H.
	0x02e0 10 84 24 98 00 00 00 0f 11 44 24 50 48 c7 44 24  ..$......D$PH.D$
	0x02f0 50 c8 00 00 00 48 8b 4c 24 48 48 89 8c 24 c0 00  P....H.L$HH..$..
	0x0300 00 00 0f 10 44 24 50 0f 11 84 24 c8 00 00 00 48  ....D$P...$....H
	0x0310 8d 05 00 00 00 00 48 8d 9c 24 c0 00 00 00 66 90  ......H..$....f.
	0x0320 e8 00 00 00 00 44 0f 11 bc 24 c8 01 00 00 48 89  .....D...$....H.
	0x0330 84 24 c8 01 00 00 48 89 9c 24 d0 01 00 00 48 8b  .$....H..$....H.
	0x0340 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x0350 00 00 00 bf 06 00 00 00 48 8d b4 24 c8 01 00 00  ........H..$....
	0x0360 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b  A.....M.......H.
	0x0370 8c 24 90 00 00 00 48 89 8c 24 c0 00 00 00 0f 10  .$....H..$......
	0x0380 84 24 98 00 00 00 0f 11 84 24 c8 00 00 00 31 c0  .$.......$....1.
	0x0390 e9 ab 00 00 00 48 89 44 24 40 48 8b 8c c4 c0 00  .....H.D$@H.....
	0x03a0 00 00 48 89 4c 24 38 e8 00 00 00 00 48 89 84 24  ..H.L$8.....H..$
	0x03b0 b0 01 00 00 48 8b 44 24 38 e8 00 00 00 00 48 8d  ....H.D$8.....H.
	0x03c0 b4 24 38 02 00 00 44 0f 11 3e 48 8d 8c 24 48 02  .$8...D..>H..$H.
	0x03d0 00 00 44 0f 11 39 48 8d 0d 00 00 00 00 48 89 8c  ..D..9H......H..
	0x03e0 24 38 02 00 00 48 8b 94 24 b0 01 00 00 48 89 94  $8...H..$....H..
	0x03f0 24 40 02 00 00 48 89 8c 24 48 02 00 00 48 89 84  $@...H..$H...H..
	0x0400 24 50 02 00 00 48 8b 1d 00 00 00 00 48 8d 05 00  $P...H......H...
	0x0410 00 00 00 bf 14 00 00 00 41 b8 02 00 00 00 4d 89  ........A.....M.
	0x0420 c1 48 8d 0d 00 00 00 00 e8 00 00 00 00 48 8b 4c  .H...........H.L
	0x0430 24 40 48 8d 41 01 66 0f 1f 84 00 00 00 00 00 90  $@H.A.f.........
	0x0440 48 83 f8 03 0f 8c 4b ff ff ff 48 c7 84 24 58 01  H.....K...H..$X.
	0x0450 00 00 00 00 00 00 48 8d bc 24 60 01 00 00 48 8d  ......H..$`...H.
	0x0460 7f d0 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00  ..H.l$.H.l$.....
	0x0470 00 48 8b 6d 00 48 c7 84 24 80 01 00 00 19 00 00  .H.m.H..$.......
	0x0480 00 48 c7 84 24 a8 01 00 00 64 00 00 00 48 8d 05  .H..$....d...H..
	0x0490 00 00 00 00 48 8d 9c 24 58 01 00 00 0f 1f 40 00  ....H..$X.....@.
	0x04a0 e8 00 00 00 00 44 0f 11 bc 24 b8 01 00 00 48 89  .....D...$....H.
	0x04b0 84 24 b8 01 00 00 48 89 9c 24 c0 01 00 00 48 8b  .$....H..$....H.
	0x04c0 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x04d0 00 00 00 bf 1a 00 00 00 48 8d b4 24 b8 01 00 00  ........H..$....
	0x04e0 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8d  A.....M.......H.
	0x04f0 8c 24 f8 00 00 00 44 0f 11 39 48 8d 8c 24 08 01  .$....D..9H..$..
	0x0500 00 00 44 0f 11 39 48 c7 84 24 f8 00 00 00 01 00  ..D..9H..$......
	0x0510 00 00 48 c7 84 24 00 01 00 00 02 00 00 00 48 c7  ..H..$........H.
	0x0520 84 24 08 01 00 00 03 00 00 00 48 c7 84 24 10 01  .$........H..$..
	0x0530 00 00 04 00 00 00 48 8d 8c 24 d8 00 00 00 44 0f  ......H..$....D.
	0x0540 11 39 48 8d 8c 24 e8 00 00 00 44 0f 11 39 48 c7  .9H..$....D..9H.
	0x0550 84 24 d8 00 00 00 01 00 00 00 48 c7 84 24 e0 00  .$........H..$..
	0x0560 00 00 02 00 00 00 48 c7 84 24 e8 00 00 00 03 00  ......H..$......
	0x0570 00 00 48 c7 84 24 f0 00 00 00 04 00 00 00 48 8b  ..H..$........H.
	0x0580 4c 24 78 48 89 8c 24 c0 00 00 00 0f 10 84 24 80  L$xH..$.......$.
	0x0590 00 00 00 0f 11 84 24 c8 00 00 00 48 8b 4c 24 60  ......$....H.L$`
	0x05a0 48 89 8c 24 a8 00 00 00 0f 10 44 24 68 0f 11 84  H..$......D$h...
	0x05b0 24 b0 00 00 00 48 8b 8c 24 c0 00 00 00 0f 1f 00  $....H..$.......
	0x05c0 48 39 8c 24 a8 00 00 00 75 27 48 8b 94 24 c8 00  H9.$....u'H..$..
	0x05d0 00 00 48 39 94 24 b0 00 00 00 75 15 48 8b 94 24  ..H9.$....u.H..$
	0x05e0 d0 00 00 00 48 39 94 24 b8 00 00 00 0f 94 c2 eb  ....H9.$........
	0x05f0 02 31 d2 0f 10 84 24 f8 00 00 00 0f 11 84 24 38  .1....$.......$8
	0x0600 01 00 00 0f 10 84 24 08 01 00 00 0f 11 84 24 48  ......$.......$H
	0x0610 01 00 00 0f 10 84 24 d8 00 00 00 0f 11 84 24 18  ......$.......$.
	0x0620 01 00 00 0f 10 84 24 e8 00 00 00 0f 11 84 24 28  ......$.......$(
	0x0630 01 00 00 4c 8b 84 24 38 01 00 00 0f 1f 44 00 00  ...L..$8.....D..
	0x0640 4c 39 84 24 18 01 00 00 75 3a 4c 8b 84 24 40 01  L9.$....u:L..$@.
	0x0650 00 00 4c 39 84 24 20 01 00 00 75 28 4c 8b 84 24  ..L9.$ ...u(L..$
	0x0660 48 01 00 00 4c 39 84 24 28 01 00 00 75 16 4c 8b  H...L9.$(...u.L.
	0x0670 84 24 50 01 00 00 4c 39 84 24 30 01 00 00 41 0f  .$P...L9.$0...A.
	0x0680 94 c0 eb 03 45 31 c0 48 8d 8c 24 18 02 00 00 44  ....E1.H..$....D
	0x0690 0f 11 39 4c 8d 8c 24 28 02 00 00 45 0f 11 39 4c  ..9L..$(...E..9L
	0x06a0 8d 0d 00 00 00 00 4c 89 8c 24 18 02 00 00 0f b6  ......L..$......
	0x06b0 d2 4c 8d 15 00 00 00 00 49 8d 14 d2 48 89 94 24  .L......I...H..$
	0x06c0 20 02 00 00 4c 89 8c 24 28 02 00 00 41 0f b6 d0   ...L..$(...A...
	0x06d0 49 8d 14 d2 48 89 94 24 30 02 00 00 48 8b 1d 00  I...H..$0...H...
	0x06e0 00 00 00 48 8d 05 00 00 00 00 bf 02 00 00 00 48  ...H...........H
	0x06f0 89 fe e8 00 00 00 00 48 8b ac 24 58 02 00 00 48  .......H..$X...H
	0x0700 81 c4 60 02 00 00 c3 e8 00 00 00 00 e9 ef f8 ff  ..`.............
	0x0710 ff                                               .
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[3]int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[3]int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[3]int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[3]int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.int+0
	rel 3+0 t=24 type.int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[11]int+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.bool+0
	rel 3+0 t=24 type.bool+0
	rel 3+0 t=24 type.[3]int+0
	rel 92+4 t=15 type.[3]int+0
	rel 105+4 t=7 runtime.convT2Enoptr+0
	rel 137+4 t=15 os.Stdout+0
	rel 144+4 t=15 go.itab.*os.File,io.Writer+0
	rel 151+4 t=15 go.string."a: %v\n"+0
	rel 178+4 t=7 fmt.Fprintf+0
	rel 275+4 t=15 type.[3]int+0
	rel 289+4 t=7 runtime.convT2Enoptr+0
	rel 321+4 t=15 os.Stdout+0
	rel 328+4 t=15 go.itab.*os.File,io.Writer+0
	rel 335+4 t=15 go.string."b: %v\n"+0
	rel 362+4 t=7 fmt.Fprintf+0
	rel 453+4 t=15 type.[3]int+0
	rel 466+4 t=7 runtime.convT2Enoptr+0
	rel 498+4 t=15 os.Stdout+0
	rel 505+4 t=15 go.itab.*os.File,io.Writer+0
	rel 512+4 t=15 go.string."c: %v\n"+0
	rel 539+4 t=7 fmt.Fprintf+0
	rel 582+4 t=15 type.[3]int+0
	rel 595+4 t=7 runtime.convT2Enoptr+0
	rel 627+4 t=15 os.Stdout+0
	rel 634+4 t=15 go.itab.*os.File,io.Writer+0
	rel 641+4 t=15 go.string."%T, d: %[1]v\n"+0
	rel 673+4 t=7 fmt.Fprintf+0
	rel 786+4 t=15 type.[3]int+0
	rel 801+4 t=7 runtime.convT2Enoptr+0
	rel 833+4 t=15 os.Stdout+0
	rel 840+4 t=15 go.itab.*os.File,io.Writer+0
	rel 847+4 t=15 go.string."e: %v\n"+0
	rel 874+4 t=7 fmt.Fprintf+0
	rel 936+4 t=7 runtime.convT64+0
	rel 954+4 t=7 runtime.convT64+0
	rel 985+4 t=15 type.int+0
	rel 1032+4 t=15 os.Stdout+0
	rel 1039+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1060+4 t=15 go.string."index: %d\tvalue: %d\n"+0
	rel 1065+4 t=7 fmt.Fprintf+0
	rel 1133+4 t=7 runtime.duffzero+336
	rel 1168+4 t=15 type.[11]int+0
	rel 1185+4 t=7 runtime.convT2Enoptr+0
	rel 1217+4 t=15 os.Stdout+0
	rel 1224+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1231+4 t=15 go.string."f type: %T\t, value: %[1]v\n"+0
	rel 1258+4 t=7 fmt.Fprintf+0
	rel 1698+4 t=15 type.bool+0
	rel 1716+4 t=15 runtime.staticuint64s+0
	rel 1759+4 t=15 os.Stdout+0
	rel 1766+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1779+4 t=7 fmt.Fprintln+0
	rel 1800+4 t=7 runtime.morestack_noctxt+0
"".arry_len_le_4 STEXT size=158 args=0x0 locals=0x70 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:63)	TEXT	"".arry_len_le_4(SB), ABIInternal, $112-0
	0x0000 00000 (06-data_types/07-array.go:63)	CMPQ	SP, 16(R14)
	0x0004 00004 (06-data_types/07-array.go:63)	PCDATA	$0, $-2
	0x0004 00004 (06-data_types/07-array.go:63)	JLS	148
	0x000a 00010 (06-data_types/07-array.go:63)	PCDATA	$0, $-1
	0x000a 00010 (06-data_types/07-array.go:63)	SUBQ	$112, SP
	0x000e 00014 (06-data_types/07-array.go:63)	MOVQ	BP, 104(SP)
	0x0013 00019 (06-data_types/07-array.go:63)	LEAQ	104(SP), BP
	0x0018 00024 (06-data_types/07-array.go:63)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0018 00024 (06-data_types/07-array.go:63)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0018 00024 (06-data_types/07-array.go:63)	FUNCDATA	$2, "".arry_len_le_4.stkobj(SB)
	0x0018 00024 (06-data_types/07-array.go:75)	MOVQ	$1, ""..autotmp_8+56(SP)
	0x0021 00033 (06-data_types/07-array.go:75)	MOVQ	$2, ""..autotmp_8+64(SP)
	0x002a 00042 (06-data_types/07-array.go:75)	MOVQ	$3, ""..autotmp_8+72(SP)
	0x0033 00051 (06-data_types/07-array.go:75)	MOVQ	$4, ""..autotmp_8+80(SP)
	0x003c 00060 (06-data_types/07-array.go:75)	LEAQ	type.[4]int64(SB), AX
	0x0043 00067 (06-data_types/07-array.go:75)	LEAQ	""..autotmp_8+56(SP), BX
	0x0048 00072 (06-data_types/07-array.go:75)	PCDATA	$1, $0
	0x0048 00072 (06-data_types/07-array.go:75)	CALL	runtime.convT2Enoptr(SB)
	0x004d 00077 (06-data_types/07-array.go:75)	MOVUPS	X15, ""..autotmp_12+88(SP)
	0x0053 00083 (06-data_types/07-array.go:75)	MOVQ	AX, ""..autotmp_12+88(SP)
	0x0058 00088 (06-data_types/07-array.go:75)	MOVQ	BX, ""..autotmp_12+96(SP)
	0x005d 00093 (<unknown line number>)	NOP
	0x005d 00093 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0064 00100 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x006b 00107 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a: %v\n"(SB), CX
	0x0072 00114 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0077 00119 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_12+88(SP), SI
	0x007c 00124 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0082 00130 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0085 00133 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x008a 00138 (06-data_types/07-array.go:76)	MOVQ	104(SP), BP
	0x008f 00143 (06-data_types/07-array.go:76)	ADDQ	$112, SP
	0x0093 00147 (06-data_types/07-array.go:76)	RET
	0x0094 00148 (06-data_types/07-array.go:76)	NOP
	0x0094 00148 (06-data_types/07-array.go:63)	PCDATA	$1, $-1
	0x0094 00148 (06-data_types/07-array.go:63)	PCDATA	$0, $-2
	0x0094 00148 (06-data_types/07-array.go:63)	CALL	runtime.morestack_noctxt(SB)
	0x0099 00153 (06-data_types/07-array.go:63)	PCDATA	$0, $-1
	0x0099 00153 (06-data_types/07-array.go:63)	JMP	0
	0x0000 49 3b 66 10 0f 86 8a 00 00 00 48 83 ec 70 48 89  I;f.......H..pH.
	0x0010 6c 24 68 48 8d 6c 24 68 48 c7 44 24 38 01 00 00  l$hH.l$hH.D$8...
	0x0020 00 48 c7 44 24 40 02 00 00 00 48 c7 44 24 48 03  .H.D$@....H.D$H.
	0x0030 00 00 00 48 c7 44 24 50 04 00 00 00 48 8d 05 00  ...H.D$P....H...
	0x0040 00 00 00 48 8d 5c 24 38 e8 00 00 00 00 44 0f 11  ...H.\$8.....D..
	0x0050 7c 24 58 48 89 44 24 58 48 89 5c 24 60 48 8b 1d  |$XH.D$XH.\$`H..
	0x0060 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00  ....H......H....
	0x0070 00 00 bf 06 00 00 00 48 8d 74 24 58 41 b8 01 00  .......H.t$XA...
	0x0080 00 00 4d 89 c1 e8 00 00 00 00 48 8b 6c 24 68 48  ..M.......H.l$hH
	0x0090 83 c4 70 c3 e8 00 00 00 00 e9 62 ff ff ff        ..p.......b...
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 63+4 t=15 type.[4]int64+0
	rel 73+4 t=7 runtime.convT2Enoptr+0
	rel 96+4 t=15 os.Stdout+0
	rel 103+4 t=15 go.itab.*os.File,io.Writer+0
	rel 110+4 t=15 go.string."a: %v\n"+0
	rel 134+4 t=7 fmt.Fprintf+0
	rel 149+4 t=7 runtime.morestack_noctxt+0
"".arry_len_more_4 STEXT size=234 args=0x0 locals=0xa0 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:78)	TEXT	"".arry_len_more_4(SB), ABIInternal, $160-0
	0x0000 00000 (06-data_types/07-array.go:78)	LEAQ	-32(SP), R12
	0x0005 00005 (06-data_types/07-array.go:78)	CMPQ	R12, 16(R14)
	0x0009 00009 (06-data_types/07-array.go:78)	PCDATA	$0, $-2
	0x0009 00009 (06-data_types/07-array.go:78)	JLS	222
	0x000f 00015 (06-data_types/07-array.go:78)	PCDATA	$0, $-1
	0x000f 00015 (06-data_types/07-array.go:78)	SUBQ	$160, SP
	0x0016 00022 (06-data_types/07-array.go:78)	MOVQ	BP, 152(SP)
	0x001e 00030 (06-data_types/07-array.go:78)	LEAQ	152(SP), BP
	0x0026 00038 (06-data_types/07-array.go:78)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0026 00038 (06-data_types/07-array.go:78)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0026 00038 (06-data_types/07-array.go:78)	FUNCDATA	$2, "".arry_len_more_4.stkobj(SB)
	0x0026 00038 (06-data_types/07-array.go:92)	MOVQ	$1, "".a+56(SP)
	0x002f 00047 (06-data_types/07-array.go:92)	MOVQ	$2, "".a+64(SP)
	0x0038 00056 (06-data_types/07-array.go:92)	MOVQ	$3, "".a+72(SP)
	0x0041 00065 (06-data_types/07-array.go:92)	MOVQ	$4, "".a+80(SP)
	0x004a 00074 (06-data_types/07-array.go:92)	MOVQ	$5, "".a+88(SP)
	0x0053 00083 (06-data_types/07-array.go:94)	MOVQ	"".a+56(SP), CX
	0x0058 00088 (06-data_types/07-array.go:94)	MOVQ	CX, ""..autotmp_8+96(SP)
	0x005d 00093 (06-data_types/07-array.go:94)	MOVUPS	"".a+64(SP), X0
	0x0062 00098 (06-data_types/07-array.go:94)	MOVUPS	X0, ""..autotmp_8+104(SP)
	0x0067 00103 (06-data_types/07-array.go:94)	MOVUPS	"".a+80(SP), X0
	0x006c 00108 (06-data_types/07-array.go:94)	MOVUPS	X0, ""..autotmp_8+120(SP)
	0x0071 00113 (06-data_types/07-array.go:94)	LEAQ	type.[5]int64(SB), AX
	0x0078 00120 (06-data_types/07-array.go:94)	LEAQ	""..autotmp_8+96(SP), BX
	0x007d 00125 (06-data_types/07-array.go:94)	PCDATA	$1, $0
	0x007d 00125 (06-data_types/07-array.go:94)	NOP
	0x0080 00128 (06-data_types/07-array.go:94)	CALL	runtime.convT2Enoptr(SB)
	0x0085 00133 (06-data_types/07-array.go:94)	MOVUPS	X15, ""..autotmp_12+136(SP)
	0x008e 00142 (06-data_types/07-array.go:94)	MOVQ	AX, ""..autotmp_12+136(SP)
	0x0096 00150 (06-data_types/07-array.go:94)	MOVQ	BX, ""..autotmp_12+144(SP)
	0x009e 00158 (<unknown line number>)	NOP
	0x009e 00158 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x00a5 00165 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00ac 00172 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a: %v\n"(SB), CX
	0x00b3 00179 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x00b8 00184 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_12+136(SP), SI
	0x00c0 00192 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x00c6 00198 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x00c9 00201 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x00ce 00206 (06-data_types/07-array.go:95)	MOVQ	152(SP), BP
	0x00d6 00214 (06-data_types/07-array.go:95)	ADDQ	$160, SP
	0x00dd 00221 (06-data_types/07-array.go:95)	RET
	0x00de 00222 (06-data_types/07-array.go:95)	NOP
	0x00de 00222 (06-data_types/07-array.go:78)	PCDATA	$1, $-1
	0x00de 00222 (06-data_types/07-array.go:78)	PCDATA	$0, $-2
	0x00de 00222 (06-data_types/07-array.go:78)	NOP
	0x00e0 00224 (06-data_types/07-array.go:78)	CALL	runtime.morestack_noctxt(SB)
	0x00e5 00229 (06-data_types/07-array.go:78)	PCDATA	$0, $-1
	0x00e5 00229 (06-data_types/07-array.go:78)	JMP	0
	0x0000 4c 8d 64 24 e0 4d 3b 66 10 0f 86 cf 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec a0 00 00 00 48 89 ac 24 98 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 98 00 00 00 48 c7 44 24 38 01 00 00 00 48  .$....H.D$8....H
	0x0030 c7 44 24 40 02 00 00 00 48 c7 44 24 48 03 00 00  .D$@....H.D$H...
	0x0040 00 48 c7 44 24 50 04 00 00 00 48 c7 44 24 58 05  .H.D$P....H.D$X.
	0x0050 00 00 00 48 8b 4c 24 38 48 89 4c 24 60 0f 10 44  ...H.L$8H.L$`..D
	0x0060 24 40 0f 11 44 24 68 0f 10 44 24 50 0f 11 44 24  $@..D$h..D$P..D$
	0x0070 78 48 8d 05 00 00 00 00 48 8d 5c 24 60 0f 1f 00  xH......H.\$`...
	0x0080 e8 00 00 00 00 44 0f 11 bc 24 88 00 00 00 48 89  .....D...$....H.
	0x0090 84 24 88 00 00 00 48 89 9c 24 90 00 00 00 48 8b  .$....H..$....H.
	0x00a0 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x00b0 00 00 00 bf 06 00 00 00 48 8d b4 24 88 00 00 00  ........H..$....
	0x00c0 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b  A.....M.......H.
	0x00d0 ac 24 98 00 00 00 48 81 c4 a0 00 00 00 c3 66 90  .$....H.......f.
	0x00e0 e8 00 00 00 00 e9 16 ff ff ff                    ..........
	rel 3+0 t=24 type.[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 116+4 t=15 type.[5]int64+0
	rel 129+4 t=7 runtime.convT2Enoptr+0
	rel 161+4 t=15 os.Stdout+0
	rel 168+4 t=15 go.itab.*os.File,io.Writer+0
	rel 175+4 t=15 go.string."a: %v\n"+0
	rel 202+4 t=7 fmt.Fprintf+0
	rel 225+4 t=7 runtime.morestack_noctxt+0
"".array_pointer STEXT size=2743 args=0x0 locals=0x230 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:97)	TEXT	"".array_pointer(SB), ABIInternal, $560-0
	0x0000 00000 (06-data_types/07-array.go:97)	LEAQ	-432(SP), R12
	0x0008 00008 (06-data_types/07-array.go:97)	CMPQ	R12, 16(R14)
	0x000c 00012 (06-data_types/07-array.go:97)	PCDATA	$0, $-2
	0x000c 00012 (06-data_types/07-array.go:97)	JLS	2733
	0x0012 00018 (06-data_types/07-array.go:97)	PCDATA	$0, $-1
	0x0012 00018 (06-data_types/07-array.go:97)	SUBQ	$560, SP
	0x0019 00025 (06-data_types/07-array.go:97)	MOVQ	BP, 552(SP)
	0x0021 00033 (06-data_types/07-array.go:97)	LEAQ	552(SP), BP
	0x0029 00041 (06-data_types/07-array.go:97)	FUNCDATA	$0, gclocals·fcf5af2016adf65a97b579a67730f1b6(SB)
	0x0029 00041 (06-data_types/07-array.go:97)	FUNCDATA	$1, gclocals·12b6f77e0db330efe75dfd918d12c72e(SB)
	0x0029 00041 (06-data_types/07-array.go:97)	FUNCDATA	$2, "".array_pointer.stkobj(SB)
	0x0029 00041 (06-data_types/07-array.go:98)	LEAQ	type.[4]int64(SB), AX
	0x0030 00048 (06-data_types/07-array.go:98)	PCDATA	$1, $0
	0x0030 00048 (06-data_types/07-array.go:98)	CALL	runtime.newobject(SB)
	0x0035 00053 (06-data_types/07-array.go:98)	MOVQ	AX, "".&a+160(SP)
	0x003d 00061 (06-data_types/07-array.go:98)	MOVQ	$1, (AX)
	0x0044 00068 (06-data_types/07-array.go:98)	MOVQ	$2, 8(AX)
	0x004c 00076 (06-data_types/07-array.go:98)	MOVQ	$3, 16(AX)
	0x0054 00084 (06-data_types/07-array.go:98)	MOVQ	$4, 24(AX)
	0x005c 00092 (06-data_types/07-array.go:99)	LEAQ	type.[4]int64(SB), AX
	0x0063 00099 (06-data_types/07-array.go:99)	PCDATA	$1, $1
	0x0063 00099 (06-data_types/07-array.go:99)	CALL	runtime.newobject(SB)
	0x0068 00104 (06-data_types/07-array.go:99)	MOVQ	AX, "".&b+152(SP)
	0x0070 00112 (06-data_types/07-array.go:99)	MOVQ	"".&a+160(SP), CX
	0x0078 00120 (06-data_types/07-array.go:99)	MOVUPS	(CX), X0
	0x007b 00123 (06-data_types/07-array.go:99)	MOVUPS	X0, (AX)
	0x007e 00126 (06-data_types/07-array.go:99)	MOVUPS	16(CX), X0
	0x0082 00130 (06-data_types/07-array.go:99)	MOVUPS	X0, 16(AX)
	0x0086 00134 (06-data_types/07-array.go:100)	LEAQ	type.*[4]int64(SB), AX
	0x008d 00141 (06-data_types/07-array.go:100)	PCDATA	$1, $2
	0x008d 00141 (06-data_types/07-array.go:100)	CALL	runtime.newobject(SB)
	0x0092 00146 (06-data_types/07-array.go:100)	MOVQ	AX, "".&c+144(SP)
	0x009a 00154 (06-data_types/07-array.go:100)	PCDATA	$0, $-2
	0x009a 00154 (06-data_types/07-array.go:100)	CMPL	runtime.writeBarrier(SB), $0
	0x00a1 00161 (06-data_types/07-array.go:100)	JNE	176
	0x00a3 00163 (06-data_types/07-array.go:100)	MOVQ	"".&a+160(SP), DX
	0x00ab 00171 (06-data_types/07-array.go:100)	MOVQ	DX, (AX)
	0x00ae 00174 (06-data_types/07-array.go:100)	JMP	197
	0x00b0 00176 (06-data_types/07-array.go:100)	MOVQ	AX, DI
	0x00b3 00179 (06-data_types/07-array.go:100)	MOVQ	"".&a+160(SP), DX
	0x00bb 00187 (06-data_types/07-array.go:100)	NOP
	0x00c0 00192 (06-data_types/07-array.go:100)	CALL	runtime.gcWriteBarrierDX(SB)
	0x00c5 00197 (06-data_types/07-array.go:102)	PCDATA	$0, $-1
	0x00c5 00197 (06-data_types/07-array.go:102)	LEAQ	go.string."="(SB), AX
	0x00cc 00204 (06-data_types/07-array.go:102)	MOVL	$1, BX
	0x00d1 00209 (06-data_types/07-array.go:102)	MOVL	$64, CX
	0x00d6 00214 (06-data_types/07-array.go:102)	PCDATA	$1, $3
	0x00d6 00214 (06-data_types/07-array.go:102)	CALL	strings.Repeat(SB)
	0x00db 00219 (06-data_types/07-array.go:102)	NOP
	0x00e0 00224 (06-data_types/07-array.go:102)	CALL	runtime.convTstring(SB)
	0x00e5 00229 (06-data_types/07-array.go:102)	MOVUPS	X15, ""..autotmp_180+536(SP)
	0x00ee 00238 (06-data_types/07-array.go:102)	LEAQ	type.string(SB), DX
	0x00f5 00245 (06-data_types/07-array.go:102)	MOVQ	DX, ""..autotmp_180+536(SP)
	0x00fd 00253 (06-data_types/07-array.go:102)	MOVQ	AX, ""..autotmp_180+544(SP)
	0x0105 00261 (<unknown line number>)	NOP
	0x0105 00261 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x010c 00268 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0113 00275 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_180+536(SP), CX
	0x011b 00283 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0120 00288 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0123 00291 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0128 00296 (06-data_types/07-array.go:103)	MOVQ	"".&a+160(SP), DX
	0x0130 00304 (06-data_types/07-array.go:103)	MOVUPS	(DX), X0
	0x0133 00307 (06-data_types/07-array.go:103)	MOVUPS	X0, ""..autotmp_172+56(SP)
	0x0138 00312 (06-data_types/07-array.go:103)	MOVUPS	16(DX), X0
	0x013c 00316 (06-data_types/07-array.go:103)	MOVUPS	X0, ""..autotmp_172+72(SP)
	0x0141 00321 (06-data_types/07-array.go:103)	LEAQ	type.[4]int64(SB), AX
	0x0148 00328 (06-data_types/07-array.go:103)	LEAQ	""..autotmp_172+56(SP), BX
	0x014d 00333 (06-data_types/07-array.go:103)	CALL	runtime.convT2Enoptr(SB)
	0x0152 00338 (06-data_types/07-array.go:103)	MOVUPS	X15, ""..autotmp_182+520(SP)
	0x015b 00347 (06-data_types/07-array.go:103)	MOVQ	AX, ""..autotmp_182+520(SP)
	0x0163 00355 (06-data_types/07-array.go:103)	MOVQ	BX, ""..autotmp_182+528(SP)
	0x016b 00363 (<unknown line number>)	NOP
	0x016b 00363 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0172 00370 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0179 00377 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a: %v\n"(SB), CX
	0x0180 00384 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0185 00389 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_182+520(SP), SI
	0x018d 00397 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0193 00403 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0196 00406 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x019b 00411 (06-data_types/07-array.go:104)	MOVQ	"".&b+152(SP), DX
	0x01a3 00419 (06-data_types/07-array.go:104)	MOVUPS	(DX), X0
	0x01a6 00422 (06-data_types/07-array.go:104)	MOVUPS	X0, ""..autotmp_172+56(SP)
	0x01ab 00427 (06-data_types/07-array.go:104)	MOVUPS	16(DX), X0
	0x01af 00431 (06-data_types/07-array.go:104)	MOVUPS	X0, ""..autotmp_172+72(SP)
	0x01b4 00436 (06-data_types/07-array.go:104)	LEAQ	type.[4]int64(SB), AX
	0x01bb 00443 (06-data_types/07-array.go:104)	LEAQ	""..autotmp_172+56(SP), BX
	0x01c0 00448 (06-data_types/07-array.go:104)	CALL	runtime.convT2Enoptr(SB)
	0x01c5 00453 (06-data_types/07-array.go:104)	MOVUPS	X15, ""..autotmp_184+504(SP)
	0x01ce 00462 (06-data_types/07-array.go:104)	MOVQ	AX, ""..autotmp_184+504(SP)
	0x01d6 00470 (06-data_types/07-array.go:104)	MOVQ	BX, ""..autotmp_184+512(SP)
	0x01de 00478 (<unknown line number>)	NOP
	0x01de 00478 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x01e5 00485 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x01ec 00492 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."b: %v\n"(SB), CX
	0x01f3 00499 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x01f8 00504 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_184+504(SP), SI
	0x0200 00512 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0206 00518 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0209 00521 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x020e 00526 (06-data_types/07-array.go:106)	MOVQ	"".&b+152(SP), DX
	0x0216 00534 (06-data_types/07-array.go:106)	MOVQ	$100, 8(DX)
	0x021e 00542 (06-data_types/07-array.go:107)	MOVUPS	(DX), X0
	0x0221 00545 (06-data_types/07-array.go:107)	MOVUPS	X0, ""..autotmp_172+56(SP)
	0x0226 00550 (06-data_types/07-array.go:107)	MOVUPS	16(DX), X0
	0x022a 00554 (06-data_types/07-array.go:107)	MOVUPS	X0, ""..autotmp_172+72(SP)
	0x022f 00559 (06-data_types/07-array.go:107)	LEAQ	type.[4]int64(SB), AX
	0x0236 00566 (06-data_types/07-array.go:107)	LEAQ	""..autotmp_172+56(SP), BX
	0x023b 00571 (06-data_types/07-array.go:107)	NOP
	0x0240 00576 (06-data_types/07-array.go:107)	CALL	runtime.convT2Enoptr(SB)
	0x0245 00581 (06-data_types/07-array.go:107)	MOVUPS	X15, ""..autotmp_186+488(SP)
	0x024e 00590 (06-data_types/07-array.go:107)	MOVQ	AX, ""..autotmp_186+488(SP)
	0x0256 00598 (06-data_types/07-array.go:107)	MOVQ	BX, ""..autotmp_186+496(SP)
	0x025e 00606 (<unknown line number>)	NOP
	0x025e 00606 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0265 00613 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x026c 00620 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."b修改后: %v\n"(SB), CX
	0x0273 00627 ($GOROOT/src/fmt/print.go:213)	MOVL	$15, DI
	0x0278 00632 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_186+488(SP), SI
	0x0280 00640 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0286 00646 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0289 00649 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x028e 00654 (06-data_types/07-array.go:108)	MOVQ	"".&a+160(SP), DX
	0x0296 00662 (06-data_types/07-array.go:108)	MOVUPS	(DX), X0
	0x0299 00665 (06-data_types/07-array.go:108)	MOVUPS	X0, ""..autotmp_172+56(SP)
	0x029e 00670 (06-data_types/07-array.go:108)	MOVUPS	16(DX), X0
	0x02a2 00674 (06-data_types/07-array.go:108)	MOVUPS	X0, ""..autotmp_172+72(SP)
	0x02a7 00679 (06-data_types/07-array.go:108)	LEAQ	type.[4]int64(SB), AX
	0x02ae 00686 (06-data_types/07-array.go:108)	LEAQ	""..autotmp_172+56(SP), BX
	0x02b3 00691 (06-data_types/07-array.go:108)	CALL	runtime.convT2Enoptr(SB)
	0x02b8 00696 (06-data_types/07-array.go:108)	MOVUPS	X15, ""..autotmp_188+472(SP)
	0x02c1 00705 (06-data_types/07-array.go:108)	MOVQ	AX, ""..autotmp_188+472(SP)
	0x02c9 00713 (06-data_types/07-array.go:108)	MOVQ	BX, ""..autotmp_188+480(SP)
	0x02d1 00721 (<unknown line number>)	NOP
	0x02d1 00721 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x02d8 00728 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x02df 00735 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."b修改后， a的值: %v\n"(SB), CX
	0x02e6 00742 ($GOROOT/src/fmt/print.go:213)	MOVL	$26, DI
	0x02eb 00747 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_188+472(SP), SI
	0x02f3 00755 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x02f9 00761 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x02fc 00764 ($GOROOT/src/fmt/print.go:213)	NOP
	0x0300 00768 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0305 00773 (06-data_types/07-array.go:109)	MOVUPS	X15, ""..autotmp_190+456(SP)
	0x030e 00782 (06-data_types/07-array.go:109)	LEAQ	type.*[4]int64(SB), DX
	0x0315 00789 (06-data_types/07-array.go:109)	MOVQ	DX, ""..autotmp_190+456(SP)
	0x031d 00797 (06-data_types/07-array.go:109)	MOVQ	"".&a+160(SP), SI
	0x0325 00805 (06-data_types/07-array.go:109)	MOVQ	SI, ""..autotmp_190+464(SP)
	0x032d 00813 (<unknown line number>)	NOP
	0x032d 00813 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0334 00820 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x033b 00827 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a 的内存地址: %p\n"(SB), CX
	0x0342 00834 ($GOROOT/src/fmt/print.go:213)	MOVL	$22, DI
	0x0347 00839 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x034d 00845 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0350 00848 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_190+456(SP), SI
	0x0358 00856 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x035d 00861 (06-data_types/07-array.go:110)	MOVUPS	X15, ""..autotmp_192+440(SP)
	0x0366 00870 (06-data_types/07-array.go:110)	LEAQ	type.*[4]int64(SB), DX
	0x036d 00877 (06-data_types/07-array.go:110)	MOVQ	DX, ""..autotmp_192+440(SP)
	0x0375 00885 (06-data_types/07-array.go:110)	MOVQ	"".&b+152(SP), SI
	0x037d 00893 (06-data_types/07-array.go:110)	MOVQ	SI, ""..autotmp_192+448(SP)
	0x0385 00901 (<unknown line number>)	NOP
	0x0385 00901 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x038c 00908 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0393 00915 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."b 的内存地址: %p\n"(SB), CX
	0x039a 00922 ($GOROOT/src/fmt/print.go:213)	MOVL	$22, DI
	0x039f 00927 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_192+440(SP), SI
	0x03a7 00935 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x03ad 00941 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x03b0 00944 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $4
	0x03b0 00944 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x03b5 00949 (06-data_types/07-array.go:116)	LEAQ	go.string."-"(SB), AX
	0x03bc 00956 (06-data_types/07-array.go:116)	MOVL	$1, BX
	0x03c1 00961 (06-data_types/07-array.go:116)	MOVL	$32, CX
	0x03c6 00966 (06-data_types/07-array.go:116)	CALL	strings.Repeat(SB)
	0x03cb 00971 (06-data_types/07-array.go:116)	CALL	runtime.convTstring(SB)
	0x03d0 00976 (06-data_types/07-array.go:116)	MOVUPS	X15, ""..autotmp_195+424(SP)
	0x03d9 00985 (06-data_types/07-array.go:116)	LEAQ	type.string(SB), DX
	0x03e0 00992 (06-data_types/07-array.go:116)	MOVQ	DX, ""..autotmp_195+424(SP)
	0x03e8 01000 (06-data_types/07-array.go:116)	MOVQ	AX, ""..autotmp_195+432(SP)
	0x03f0 01008 (<unknown line number>)	NOP
	0x03f0 01008 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x03f7 01015 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x03fe 01022 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_195+424(SP), CX
	0x0406 01030 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x040b 01035 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x040e 01038 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0413 01043 (06-data_types/07-array.go:118)	MOVQ	"".&c+144(SP), DX
	0x041b 01051 (06-data_types/07-array.go:118)	MOVQ	(DX), SI
	0x041e 01054 (06-data_types/07-array.go:118)	MOVQ	$200, 16(SI)
	0x0426 01062 (06-data_types/07-array.go:119)	MOVQ	(DX), BX
	0x0429 01065 (06-data_types/07-array.go:119)	TESTB	AL, (BX)
	0x042b 01067 (06-data_types/07-array.go:119)	LEAQ	type.[4]int64(SB), AX
	0x0432 01074 (06-data_types/07-array.go:119)	CALL	runtime.convT2Enoptr(SB)
	0x0437 01079 (06-data_types/07-array.go:119)	MOVUPS	X15, ""..autotmp_198+408(SP)
	0x0440 01088 (06-data_types/07-array.go:119)	MOVQ	AX, ""..autotmp_198+408(SP)
	0x0448 01096 (06-data_types/07-array.go:119)	MOVQ	BX, ""..autotmp_198+416(SP)
	0x0450 01104 (<unknown line number>)	NOP
	0x0450 01104 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0457 01111 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x045e 01118 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."c: %v\n"(SB), CX
	0x0465 01125 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x046a 01130 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_198+408(SP), SI
	0x0472 01138 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0478 01144 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x047b 01147 ($GOROOT/src/fmt/print.go:213)	NOP
	0x0480 01152 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0485 01157 (06-data_types/07-array.go:120)	MOVQ	"".&a+160(SP), DX
	0x048d 01165 (06-data_types/07-array.go:120)	MOVUPS	(DX), X0
	0x0490 01168 (06-data_types/07-array.go:120)	MOVUPS	X0, ""..autotmp_172+56(SP)
	0x0495 01173 (06-data_types/07-array.go:120)	MOVUPS	16(DX), X0
	0x0499 01177 (06-data_types/07-array.go:120)	MOVUPS	X0, ""..autotmp_172+72(SP)
	0x049e 01182 (06-data_types/07-array.go:120)	LEAQ	type.[4]int64(SB), AX
	0x04a5 01189 (06-data_types/07-array.go:120)	LEAQ	""..autotmp_172+56(SP), BX
	0x04aa 01194 (06-data_types/07-array.go:120)	CALL	runtime.convT2Enoptr(SB)
	0x04af 01199 (06-data_types/07-array.go:120)	MOVUPS	X15, ""..autotmp_200+392(SP)
	0x04b8 01208 (06-data_types/07-array.go:120)	MOVQ	AX, ""..autotmp_200+392(SP)
	0x04c0 01216 (06-data_types/07-array.go:120)	MOVQ	BX, ""..autotmp_200+400(SP)
	0x04c8 01224 (<unknown line number>)	NOP
	0x04c8 01224 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x04cf 01231 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x04d6 01238 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a: %v\n"(SB), CX
	0x04dd 01245 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x04e2 01250 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_200+392(SP), SI
	0x04ea 01258 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x04f0 01264 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x04f3 01267 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x04f8 01272 (06-data_types/07-array.go:121)	MOVUPS	X15, ""..autotmp_202+376(SP)
	0x0501 01281 (06-data_types/07-array.go:121)	LEAQ	type.*[4]int64(SB), DX
	0x0508 01288 (06-data_types/07-array.go:121)	MOVQ	DX, ""..autotmp_202+376(SP)
	0x0510 01296 (06-data_types/07-array.go:121)	MOVQ	"".&a+160(SP), SI
	0x0518 01304 (06-data_types/07-array.go:121)	MOVQ	SI, ""..autotmp_202+384(SP)
	0x0520 01312 (<unknown line number>)	NOP
	0x0520 01312 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0527 01319 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x052e 01326 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."&a: %p\n"(SB), CX
	0x0535 01333 ($GOROOT/src/fmt/print.go:213)	MOVL	$7, DI
	0x053a 01338 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_202+376(SP), SI
	0x0542 01346 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0548 01352 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x054b 01355 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $5
	0x054b 01355 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0550 01360 (06-data_types/07-array.go:122)	MOVQ	"".&c+144(SP), DX
	0x0558 01368 (06-data_types/07-array.go:122)	MOVQ	(DX), SI
	0x055b 01371 (06-data_types/07-array.go:122)	MOVUPS	X15, ""..autotmp_204+360(SP)
	0x0564 01380 (06-data_types/07-array.go:122)	LEAQ	type.*[4]int64(SB), DI
	0x056b 01387 (06-data_types/07-array.go:122)	MOVQ	DI, ""..autotmp_204+360(SP)
	0x0573 01395 (06-data_types/07-array.go:122)	MOVQ	SI, ""..autotmp_204+368(SP)
	0x057b 01403 (<unknown line number>)	NOP
	0x057b 01403 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0582 01410 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0589 01417 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."c: %p，以%%p打印会打印出它的值：数组a的首地址\n"(SB), CX
	0x0590 01424 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_204+360(SP), SI
	0x0598 01432 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x059e 01438 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x05a1 01441 ($GOROOT/src/fmt/print.go:213)	MOVL	$64, DI
	0x05a6 01446 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x05ab 01451 (06-data_types/07-array.go:123)	MOVQ	"".&c+144(SP), DX
	0x05b3 01459 (06-data_types/07-array.go:123)	MOVQ	(DX), SI
	0x05b6 01462 (06-data_types/07-array.go:123)	MOVUPS	X15, ""..autotmp_206+344(SP)
	0x05bf 01471 (06-data_types/07-array.go:123)	LEAQ	type.*[4]int64(SB), DI
	0x05c6 01478 (06-data_types/07-array.go:123)	MOVQ	DI, ""..autotmp_206+344(SP)
	0x05ce 01486 (06-data_types/07-array.go:123)	MOVQ	SI, ""..autotmp_206+352(SP)
	0x05d6 01494 (<unknown line number>)	NOP
	0x05d6 01494 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x05dd 01501 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x05e4 01508 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string..gostring.107.fc25da9a357bd0ef5c90db92d19f19d232fe9dcb48a70860742a11e5d80e44c2(SB), CX
	0x05eb 01515 ($GOROOT/src/fmt/print.go:213)	MOVL	$107, DI
	0x05f0 01520 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_206+344(SP), SI
	0x05f8 01528 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x05fe 01534 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0601 01537 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0606 01542 (06-data_types/07-array.go:124)	MOVUPS	X15, ""..autotmp_208+328(SP)
	0x060f 01551 (06-data_types/07-array.go:124)	LEAQ	type.**[4]int64(SB), DX
	0x0616 01558 (06-data_types/07-array.go:124)	MOVQ	DX, ""..autotmp_208+328(SP)
	0x061e 01566 (06-data_types/07-array.go:124)	MOVQ	"".&c+144(SP), DX
	0x0626 01574 (06-data_types/07-array.go:124)	MOVQ	DX, ""..autotmp_208+336(SP)
	0x062e 01582 (<unknown line number>)	NOP
	0x062e 01582 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0635 01589 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x063c 01596 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."&c: %p，变量c在栈上的地址，它的内容时指向a的指针\n"(SB), CX
	0x0643 01603 ($GOROOT/src/fmt/print.go:213)	MOVL	$69, DI
	0x0648 01608 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_208+328(SP), SI
	0x0650 01616 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0656 01622 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0659 01625 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x065e 01630 (06-data_types/07-array.go:125)	MOVQ	"".&c+144(SP), DX
	0x0666 01638 (06-data_types/07-array.go:125)	MOVQ	(DX), BX
	0x0669 01641 (06-data_types/07-array.go:125)	TESTB	AL, (BX)
	0x066b 01643 (06-data_types/07-array.go:125)	LEAQ	type.[4]int64(SB), AX
	0x0672 01650 (06-data_types/07-array.go:125)	PCDATA	$1, $0
	0x0672 01650 (06-data_types/07-array.go:125)	CALL	runtime.convT2Enoptr(SB)
	0x0677 01655 (06-data_types/07-array.go:125)	MOVUPS	X15, ""..autotmp_211+312(SP)
	0x0680 01664 (06-data_types/07-array.go:125)	MOVQ	AX, ""..autotmp_211+312(SP)
	0x0688 01672 (06-data_types/07-array.go:125)	MOVQ	BX, ""..autotmp_211+320(SP)
	0x0690 01680 (<unknown line number>)	NOP
	0x0690 01680 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0697 01687 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x069e 01694 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."*c: %v，打印*c会打出c指向的内容，即a的内容\n"(SB), CX
	0x06a5 01701 ($GOROOT/src/fmt/print.go:213)	MOVL	$59, DI
	0x06aa 01706 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_211+312(SP), SI
	0x06b2 01714 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x06b8 01720 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x06bb 01723 ($GOROOT/src/fmt/print.go:213)	NOP
	0x06c0 01728 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x06c5 01733 (06-data_types/07-array.go:127)	LEAQ	go.string."-"(SB), AX
	0x06cc 01740 (06-data_types/07-array.go:127)	MOVL	$1, BX
	0x06d1 01745 (06-data_types/07-array.go:127)	MOVL	$64, CX
	0x06d6 01750 (06-data_types/07-array.go:127)	CALL	strings.Repeat(SB)
	0x06db 01755 (06-data_types/07-array.go:127)	NOP
	0x06e0 01760 (06-data_types/07-array.go:127)	CALL	runtime.convTstring(SB)
	0x06e5 01765 (06-data_types/07-array.go:127)	MOVUPS	X15, ""..autotmp_214+296(SP)
	0x06ee 01774 (06-data_types/07-array.go:127)	LEAQ	type.string(SB), DX
	0x06f5 01781 (06-data_types/07-array.go:127)	MOVQ	DX, ""..autotmp_214+296(SP)
	0x06fd 01789 (06-data_types/07-array.go:127)	MOVQ	AX, ""..autotmp_214+304(SP)
	0x0705 01797 (<unknown line number>)	NOP
	0x0705 01797 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x070c 01804 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0713 01811 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_214+296(SP), CX
	0x071b 01819 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0720 01824 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0723 01827 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0728 01832 (06-data_types/07-array.go:129)	LEAQ	type.[5]int64(SB), AX
	0x072f 01839 (06-data_types/07-array.go:129)	CALL	runtime.newobject(SB)
	0x0734 01844 (06-data_types/07-array.go:129)	MOVQ	AX, "".&d+136(SP)
	0x073c 01852 (06-data_types/07-array.go:129)	MOVQ	$1, (AX)
	0x0743 01859 (06-data_types/07-array.go:129)	MOVQ	$2, 8(AX)
	0x074b 01867 (06-data_types/07-array.go:129)	MOVQ	$3, 16(AX)
	0x0753 01875 (06-data_types/07-array.go:129)	MOVQ	$4, 24(AX)
	0x075b 01883 (06-data_types/07-array.go:129)	MOVQ	$5, 32(AX)
	0x0763 01891 (06-data_types/07-array.go:130)	LEAQ	type.*[5]int64(SB), AX
	0x076a 01898 (06-data_types/07-array.go:130)	PCDATA	$1, $6
	0x076a 01898 (06-data_types/07-array.go:130)	CALL	runtime.newobject(SB)
	0x076f 01903 (06-data_types/07-array.go:130)	MOVQ	AX, "".&f+128(SP)
	0x0777 01911 (06-data_types/07-array.go:130)	PCDATA	$0, $-2
	0x0777 01911 (06-data_types/07-array.go:130)	CMPL	runtime.writeBarrier(SB), $0
	0x077e 01918 (06-data_types/07-array.go:130)	NOP
	0x0780 01920 (06-data_types/07-array.go:130)	JNE	1935
	0x0782 01922 (06-data_types/07-array.go:130)	MOVQ	"".&d+136(SP), CX
	0x078a 01930 (06-data_types/07-array.go:130)	MOVQ	CX, (AX)
	0x078d 01933 (06-data_types/07-array.go:130)	JMP	1951
	0x078f 01935 (06-data_types/07-array.go:130)	MOVQ	AX, DI
	0x0792 01938 (06-data_types/07-array.go:130)	MOVQ	"".&d+136(SP), CX
	0x079a 01946 (06-data_types/07-array.go:130)	CALL	runtime.gcWriteBarrierCX(SB)
	0x079f 01951 (06-data_types/07-array.go:131)	PCDATA	$0, $-1
	0x079f 01951 (06-data_types/07-array.go:131)	MOVQ	(CX), DX
	0x07a2 01954 (06-data_types/07-array.go:131)	MOVQ	DX, ""..autotmp_175+88(SP)
	0x07a7 01959 (06-data_types/07-array.go:131)	MOVUPS	8(CX), X0
	0x07ab 01963 (06-data_types/07-array.go:131)	MOVUPS	X0, ""..autotmp_175+96(SP)
	0x07b0 01968 (06-data_types/07-array.go:131)	MOVUPS	24(CX), X0
	0x07b4 01972 (06-data_types/07-array.go:131)	MOVUPS	X0, ""..autotmp_175+112(SP)
	0x07b9 01977 (06-data_types/07-array.go:131)	LEAQ	type.[5]int64(SB), AX
	0x07c0 01984 (06-data_types/07-array.go:131)	LEAQ	""..autotmp_175+88(SP), BX
	0x07c5 01989 (06-data_types/07-array.go:131)	PCDATA	$1, $7
	0x07c5 01989 (06-data_types/07-array.go:131)	CALL	runtime.convT2Enoptr(SB)
	0x07ca 01994 (06-data_types/07-array.go:131)	MOVUPS	X15, ""..autotmp_216+280(SP)
	0x07d3 02003 (06-data_types/07-array.go:131)	MOVQ	AX, ""..autotmp_216+280(SP)
	0x07db 02011 (06-data_types/07-array.go:131)	MOVQ	BX, ""..autotmp_216+288(SP)
	0x07e3 02019 (<unknown line number>)	NOP
	0x07e3 02019 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x07ea 02026 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x07f1 02033 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."d: %v\n"(SB), CX
	0x07f8 02040 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x07fd 02045 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_216+280(SP), SI
	0x0805 02053 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x080b 02059 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x080e 02062 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0813 02067 (06-data_types/07-array.go:132)	MOVQ	"".&f+128(SP), CX
	0x081b 02075 (06-data_types/07-array.go:132)	MOVQ	(CX), DX
	0x081e 02078 (06-data_types/07-array.go:132)	MOVUPS	X15, ""..autotmp_218+264(SP)
	0x0827 02087 (06-data_types/07-array.go:132)	LEAQ	type.*[5]int64(SB), SI
	0x082e 02094 (06-data_types/07-array.go:132)	MOVQ	SI, ""..autotmp_218+264(SP)
	0x0836 02102 (06-data_types/07-array.go:132)	MOVQ	DX, ""..autotmp_218+272(SP)
	0x083e 02110 (<unknown line number>)	NOP
	0x083e 02110 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0845 02117 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x084c 02124 ($GOROOT/src/fmt/print.go:213)	MOVL	$6, DI
	0x0851 02129 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0857 02135 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x085a 02138 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."f: %v\n"(SB), CX
	0x0861 02145 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_218+264(SP), SI
	0x0869 02153 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x086e 02158 (06-data_types/07-array.go:133)	MOVUPS	X15, ""..autotmp_220+248(SP)
	0x0877 02167 (06-data_types/07-array.go:133)	LEAQ	type.*[5]int64(SB), CX
	0x087e 02174 (06-data_types/07-array.go:133)	MOVQ	CX, ""..autotmp_220+248(SP)
	0x0886 02182 (06-data_types/07-array.go:133)	MOVQ	"".&d+136(SP), DX
	0x088e 02190 (06-data_types/07-array.go:133)	MOVQ	DX, ""..autotmp_220+256(SP)
	0x0896 02198 (<unknown line number>)	NOP
	0x0896 02198 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x089d 02205 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x08a4 02212 ($GOROOT/src/fmt/print.go:213)	MOVL	$7, DI
	0x08a9 02217 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_220+248(SP), SI
	0x08b1 02225 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x08b7 02231 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x08ba 02234 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."&d: %p\n"(SB), CX
	0x08c1 02241 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $8
	0x08c1 02241 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x08c6 02246 (06-data_types/07-array.go:134)	MOVQ	"".&f+128(SP), CX
	0x08ce 02254 (06-data_types/07-array.go:134)	MOVQ	(CX), DX
	0x08d1 02257 (06-data_types/07-array.go:134)	MOVUPS	X15, ""..autotmp_222+232(SP)
	0x08da 02266 (06-data_types/07-array.go:134)	LEAQ	type.*[5]int64(SB), SI
	0x08e1 02273 (06-data_types/07-array.go:134)	MOVQ	SI, ""..autotmp_222+232(SP)
	0x08e9 02281 (06-data_types/07-array.go:134)	MOVQ	DX, ""..autotmp_222+240(SP)
	0x08f1 02289 (<unknown line number>)	NOP
	0x08f1 02289 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x08f8 02296 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x08ff 02303 ($GOROOT/src/fmt/print.go:213)	MOVL	$64, DI
	0x0904 02308 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x090a 02314 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x090d 02317 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."f: %p，以%%p打印会打印出它的值：数组d的首地址\n"(SB), CX
	0x0914 02324 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_222+232(SP), SI
	0x091c 02332 ($GOROOT/src/fmt/print.go:213)	NOP
	0x0920 02336 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0925 02341 (06-data_types/07-array.go:135)	MOVQ	"".&f+128(SP), CX
	0x092d 02349 (06-data_types/07-array.go:135)	MOVQ	(CX), DX
	0x0930 02352 (06-data_types/07-array.go:135)	MOVUPS	X15, ""..autotmp_224+216(SP)
	0x0939 02361 (06-data_types/07-array.go:135)	LEAQ	type.*[5]int64(SB), SI
	0x0940 02368 (06-data_types/07-array.go:135)	MOVQ	SI, ""..autotmp_224+216(SP)
	0x0948 02376 (06-data_types/07-array.go:135)	MOVQ	DX, ""..autotmp_224+224(SP)
	0x0950 02384 (<unknown line number>)	NOP
	0x0950 02384 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0957 02391 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x095e 02398 ($GOROOT/src/fmt/print.go:213)	MOVL	$107, DI
	0x0963 02403 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_224+216(SP), SI
	0x096b 02411 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0971 02417 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0974 02420 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string..gostring.107.958de85d687a58aae8b5f361b1c484f3c0d64dab64309d7e9d0b8869bff01f62(SB), CX
	0x097b 02427 ($GOROOT/src/fmt/print.go:213)	NOP
	0x0980 02432 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0985 02437 (06-data_types/07-array.go:136)	MOVUPS	X15, ""..autotmp_226+200(SP)
	0x098e 02446 (06-data_types/07-array.go:136)	LEAQ	type.**[5]int64(SB), CX
	0x0995 02453 (06-data_types/07-array.go:136)	MOVQ	CX, ""..autotmp_226+200(SP)
	0x099d 02461 (06-data_types/07-array.go:136)	MOVQ	"".&f+128(SP), CX
	0x09a5 02469 (06-data_types/07-array.go:136)	MOVQ	CX, ""..autotmp_226+208(SP)
	0x09ad 02477 (<unknown line number>)	NOP
	0x09ad 02477 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x09b4 02484 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x09bb 02491 ($GOROOT/src/fmt/print.go:213)	MOVL	$69, DI
	0x09c0 02496 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_226+200(SP), SI
	0x09c8 02504 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x09ce 02510 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x09d1 02513 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."&f: %p，变量f在栈上的地址，它的内容是指向d的指针\n"(SB), CX
	0x09d8 02520 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x09dd 02525 (06-data_types/07-array.go:137)	MOVQ	"".&f+128(SP), CX
	0x09e5 02533 (06-data_types/07-array.go:137)	MOVQ	(CX), BX
	0x09e8 02536 (06-data_types/07-array.go:137)	TESTB	AL, (BX)
	0x09ea 02538 (06-data_types/07-array.go:137)	LEAQ	type.[5]int64(SB), AX
	0x09f1 02545 (06-data_types/07-array.go:137)	PCDATA	$1, $0
	0x09f1 02545 (06-data_types/07-array.go:137)	CALL	runtime.convT2Enoptr(SB)
	0x09f6 02550 (06-data_types/07-array.go:137)	MOVUPS	X15, ""..autotmp_229+184(SP)
	0x09ff 02559 (06-data_types/07-array.go:137)	MOVQ	AX, ""..autotmp_229+184(SP)
	0x0a07 02567 (06-data_types/07-array.go:137)	MOVQ	BX, ""..autotmp_229+192(SP)
	0x0a0f 02575 (<unknown line number>)	NOP
	0x0a0f 02575 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x0a16 02582 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0a1d 02589 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."*f: %v，打印*f会打出f指向的内容，即d的内容\n"(SB), CX
	0x0a24 02596 ($GOROOT/src/fmt/print.go:213)	MOVL	$59, DI
	0x0a29 02601 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_229+184(SP), SI
	0x0a31 02609 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x0a37 02615 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x0a3a 02618 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0a3f 02623 (06-data_types/07-array.go:138)	LEAQ	go.string."="(SB), AX
	0x0a46 02630 (06-data_types/07-array.go:138)	MOVL	$1, BX
	0x0a4b 02635 (06-data_types/07-array.go:138)	MOVL	$64, CX
	0x0a50 02640 (06-data_types/07-array.go:138)	CALL	strings.Repeat(SB)
	0x0a55 02645 (06-data_types/07-array.go:138)	CALL	runtime.convTstring(SB)
	0x0a5a 02650 (06-data_types/07-array.go:138)	MOVUPS	X15, ""..autotmp_232+168(SP)
	0x0a63 02659 (06-data_types/07-array.go:138)	LEAQ	type.string(SB), CX
	0x0a6a 02666 (06-data_types/07-array.go:138)	MOVQ	CX, ""..autotmp_232+168(SP)
	0x0a72 02674 (06-data_types/07-array.go:138)	MOVQ	AX, ""..autotmp_232+176(SP)
	0x0a7a 02682 (<unknown line number>)	NOP
	0x0a7a 02682 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x0a81 02689 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0a88 02696 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_232+168(SP), CX
	0x0a90 02704 ($GOROOT/src/fmt/print.go:274)	MOVL	$1, DI
	0x0a95 02709 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0a98 02712 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0a9d 02717 (06-data_types/07-array.go:139)	MOVQ	552(SP), BP
	0x0aa5 02725 (06-data_types/07-array.go:139)	ADDQ	$560, SP
	0x0aac 02732 (06-data_types/07-array.go:139)	RET
	0x0aad 02733 (06-data_types/07-array.go:139)	NOP
	0x0aad 02733 (06-data_types/07-array.go:97)	PCDATA	$1, $-1
	0x0aad 02733 (06-data_types/07-array.go:97)	PCDATA	$0, $-2
	0x0aad 02733 (06-data_types/07-array.go:97)	CALL	runtime.morestack_noctxt(SB)
	0x0ab2 02738 (06-data_types/07-array.go:97)	PCDATA	$0, $-1
	0x0ab2 02738 (06-data_types/07-array.go:97)	JMP	0
	0x0000 4c 8d a4 24 50 fe ff ff 4d 3b 66 10 0f 86 9b 0a  L..$P...M;f.....
	0x0010 00 00 48 81 ec 30 02 00 00 48 89 ac 24 28 02 00  ..H..0...H..$(..
	0x0020 00 48 8d ac 24 28 02 00 00 48 8d 05 00 00 00 00  .H..$(...H......
	0x0030 e8 00 00 00 00 48 89 84 24 a0 00 00 00 48 c7 00  .....H..$....H..
	0x0040 01 00 00 00 48 c7 40 08 02 00 00 00 48 c7 40 10  ....H.@.....H.@.
	0x0050 03 00 00 00 48 c7 40 18 04 00 00 00 48 8d 05 00  ....H.@.....H...
	0x0060 00 00 00 e8 00 00 00 00 48 89 84 24 98 00 00 00  ........H..$....
	0x0070 48 8b 8c 24 a0 00 00 00 0f 10 01 0f 11 00 0f 10  H..$............
	0x0080 41 10 0f 11 40 10 48 8d 05 00 00 00 00 e8 00 00  A...@.H.........
	0x0090 00 00 48 89 84 24 90 00 00 00 83 3d 00 00 00 00  ..H..$.....=....
	0x00a0 00 75 0d 48 8b 94 24 a0 00 00 00 48 89 10 eb 15  .u.H..$....H....
	0x00b0 48 89 c7 48 8b 94 24 a0 00 00 00 0f 1f 44 00 00  H..H..$......D..
	0x00c0 e8 00 00 00 00 48 8d 05 00 00 00 00 bb 01 00 00  .....H..........
	0x00d0 00 b9 40 00 00 00 e8 00 00 00 00 0f 1f 44 00 00  ..@..........D..
	0x00e0 e8 00 00 00 00 44 0f 11 bc 24 18 02 00 00 48 8d  .....D...$....H.
	0x00f0 15 00 00 00 00 48 89 94 24 18 02 00 00 48 89 84  .....H..$....H..
	0x0100 24 20 02 00 00 48 8b 1d 00 00 00 00 48 8d 05 00  $ ...H......H...
	0x0110 00 00 00 48 8d 8c 24 18 02 00 00 bf 01 00 00 00  ...H..$.........
	0x0120 48 89 fe e8 00 00 00 00 48 8b 94 24 a0 00 00 00  H.......H..$....
	0x0130 0f 10 02 0f 11 44 24 38 0f 10 42 10 0f 11 44 24  .....D$8..B...D$
	0x0140 48 48 8d 05 00 00 00 00 48 8d 5c 24 38 e8 00 00  HH......H.\$8...
	0x0150 00 00 44 0f 11 bc 24 08 02 00 00 48 89 84 24 08  ..D...$....H..$.
	0x0160 02 00 00 48 89 9c 24 10 02 00 00 48 8b 1d 00 00  ...H..$....H....
	0x0170 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00 00 00  ..H......H......
	0x0180 bf 06 00 00 00 48 8d b4 24 08 02 00 00 41 b8 01  .....H..$....A..
	0x0190 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b 94 24 98  ...M.......H..$.
	0x01a0 00 00 00 0f 10 02 0f 11 44 24 38 0f 10 42 10 0f  ........D$8..B..
	0x01b0 11 44 24 48 48 8d 05 00 00 00 00 48 8d 5c 24 38  .D$HH......H.\$8
	0x01c0 e8 00 00 00 00 44 0f 11 bc 24 f8 01 00 00 48 89  .....D...$....H.
	0x01d0 84 24 f8 01 00 00 48 89 9c 24 00 02 00 00 48 8b  .$....H..$....H.
	0x01e0 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x01f0 00 00 00 bf 06 00 00 00 48 8d b4 24 f8 01 00 00  ........H..$....
	0x0200 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b  A.....M.......H.
	0x0210 94 24 98 00 00 00 48 c7 42 08 64 00 00 00 0f 10  .$....H.B.d.....
	0x0220 02 0f 11 44 24 38 0f 10 42 10 0f 11 44 24 48 48  ...D$8..B...D$HH
	0x0230 8d 05 00 00 00 00 48 8d 5c 24 38 0f 1f 44 00 00  ......H.\$8..D..
	0x0240 e8 00 00 00 00 44 0f 11 bc 24 e8 01 00 00 48 89  .....D...$....H.
	0x0250 84 24 e8 01 00 00 48 89 9c 24 f0 01 00 00 48 8b  .$....H..$....H.
	0x0260 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x0270 00 00 00 bf 0f 00 00 00 48 8d b4 24 e8 01 00 00  ........H..$....
	0x0280 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b  A.....M.......H.
	0x0290 94 24 a0 00 00 00 0f 10 02 0f 11 44 24 38 0f 10  .$.........D$8..
	0x02a0 42 10 0f 11 44 24 48 48 8d 05 00 00 00 00 48 8d  B...D$HH......H.
	0x02b0 5c 24 38 e8 00 00 00 00 44 0f 11 bc 24 d8 01 00  \$8.....D...$...
	0x02c0 00 48 89 84 24 d8 01 00 00 48 89 9c 24 e0 01 00  .H..$....H..$...
	0x02d0 00 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48  .H......H......H
	0x02e0 8d 0d 00 00 00 00 bf 1a 00 00 00 48 8d b4 24 d8  ...........H..$.
	0x02f0 01 00 00 41 b8 01 00 00 00 4d 89 c1 0f 1f 40 00  ...A.....M....@.
	0x0300 e8 00 00 00 00 44 0f 11 bc 24 c8 01 00 00 48 8d  .....D...$....H.
	0x0310 15 00 00 00 00 48 89 94 24 c8 01 00 00 48 8b b4  .....H..$....H..
	0x0320 24 a0 00 00 00 48 89 b4 24 d0 01 00 00 48 8b 1d  $....H..$....H..
	0x0330 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00  ....H......H....
	0x0340 00 00 bf 16 00 00 00 41 b8 01 00 00 00 4d 89 c1  .......A.....M..
	0x0350 48 8d b4 24 c8 01 00 00 e8 00 00 00 00 44 0f 11  H..$.........D..
	0x0360 bc 24 b8 01 00 00 48 8d 15 00 00 00 00 48 89 94  .$....H......H..
	0x0370 24 b8 01 00 00 48 8b b4 24 98 00 00 00 48 89 b4  $....H..$....H..
	0x0380 24 c0 01 00 00 48 8b 1d 00 00 00 00 48 8d 05 00  $....H......H...
	0x0390 00 00 00 48 8d 0d 00 00 00 00 bf 16 00 00 00 48  ...H...........H
	0x03a0 8d b4 24 b8 01 00 00 41 b8 01 00 00 00 4d 89 c1  ..$....A.....M..
	0x03b0 e8 00 00 00 00 48 8d 05 00 00 00 00 bb 01 00 00  .....H..........
	0x03c0 00 b9 20 00 00 00 e8 00 00 00 00 e8 00 00 00 00  .. .............
	0x03d0 44 0f 11 bc 24 a8 01 00 00 48 8d 15 00 00 00 00  D...$....H......
	0x03e0 48 89 94 24 a8 01 00 00 48 89 84 24 b0 01 00 00  H..$....H..$....
	0x03f0 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x0400 8c 24 a8 01 00 00 bf 01 00 00 00 48 89 fe e8 00  .$.........H....
	0x0410 00 00 00 48 8b 94 24 90 00 00 00 48 8b 32 48 c7  ...H..$....H.2H.
	0x0420 46 10 c8 00 00 00 48 8b 1a 84 03 48 8d 05 00 00  F.....H....H....
	0x0430 00 00 e8 00 00 00 00 44 0f 11 bc 24 98 01 00 00  .......D...$....
	0x0440 48 89 84 24 98 01 00 00 48 89 9c 24 a0 01 00 00  H..$....H..$....
	0x0450 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x0460 0d 00 00 00 00 bf 06 00 00 00 48 8d b4 24 98 01  ..........H..$..
	0x0470 00 00 41 b8 01 00 00 00 4d 89 c1 0f 1f 44 00 00  ..A.....M....D..
	0x0480 e8 00 00 00 00 48 8b 94 24 a0 00 00 00 0f 10 02  .....H..$.......
	0x0490 0f 11 44 24 38 0f 10 42 10 0f 11 44 24 48 48 8d  ..D$8..B...D$HH.
	0x04a0 05 00 00 00 00 48 8d 5c 24 38 e8 00 00 00 00 44  .....H.\$8.....D
	0x04b0 0f 11 bc 24 88 01 00 00 48 89 84 24 88 01 00 00  ...$....H..$....
	0x04c0 48 89 9c 24 90 01 00 00 48 8b 1d 00 00 00 00 48  H..$....H......H
	0x04d0 8d 05 00 00 00 00 48 8d 0d 00 00 00 00 bf 06 00  ......H.........
	0x04e0 00 00 48 8d b4 24 88 01 00 00 41 b8 01 00 00 00  ..H..$....A.....
	0x04f0 4d 89 c1 e8 00 00 00 00 44 0f 11 bc 24 78 01 00  M.......D...$x..
	0x0500 00 48 8d 15 00 00 00 00 48 89 94 24 78 01 00 00  .H......H..$x...
	0x0510 48 8b b4 24 a0 00 00 00 48 89 b4 24 80 01 00 00  H..$....H..$....
	0x0520 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x0530 0d 00 00 00 00 bf 07 00 00 00 48 8d b4 24 78 01  ..........H..$x.
	0x0540 00 00 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00  ..A.....M.......
	0x0550 48 8b 94 24 90 00 00 00 48 8b 32 44 0f 11 bc 24  H..$....H.2D...$
	0x0560 68 01 00 00 48 8d 3d 00 00 00 00 48 89 bc 24 68  h...H.=....H..$h
	0x0570 01 00 00 48 89 b4 24 70 01 00 00 48 8b 1d 00 00  ...H..$p...H....
	0x0580 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00 00 00  ..H......H......
	0x0590 48 8d b4 24 68 01 00 00 41 b8 01 00 00 00 4d 89  H..$h...A.....M.
	0x05a0 c1 bf 40 00 00 00 e8 00 00 00 00 48 8b 94 24 90  ..@........H..$.
	0x05b0 00 00 00 48 8b 32 44 0f 11 bc 24 58 01 00 00 48  ...H.2D...$X...H
	0x05c0 8d 3d 00 00 00 00 48 89 bc 24 58 01 00 00 48 89  .=....H..$X...H.
	0x05d0 b4 24 60 01 00 00 48 8b 1d 00 00 00 00 48 8d 05  .$`...H......H..
	0x05e0 00 00 00 00 48 8d 0d 00 00 00 00 bf 6b 00 00 00  ....H.......k...
	0x05f0 48 8d b4 24 58 01 00 00 41 b8 01 00 00 00 4d 89  H..$X...A.....M.
	0x0600 c1 e8 00 00 00 00 44 0f 11 bc 24 48 01 00 00 48  ......D...$H...H
	0x0610 8d 15 00 00 00 00 48 89 94 24 48 01 00 00 48 8b  ......H..$H...H.
	0x0620 94 24 90 00 00 00 48 89 94 24 50 01 00 00 48 8b  .$....H..$P...H.
	0x0630 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00  .....H......H...
	0x0640 00 00 00 bf 45 00 00 00 48 8d b4 24 48 01 00 00  ....E...H..$H...
	0x0650 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48 8b  A.....M.......H.
	0x0660 94 24 90 00 00 00 48 8b 1a 84 03 48 8d 05 00 00  .$....H....H....
	0x0670 00 00 e8 00 00 00 00 44 0f 11 bc 24 38 01 00 00  .......D...$8...
	0x0680 48 89 84 24 38 01 00 00 48 89 9c 24 40 01 00 00  H..$8...H..$@...
	0x0690 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d  H......H......H.
	0x06a0 0d 00 00 00 00 bf 3b 00 00 00 48 8d b4 24 38 01  ......;...H..$8.
	0x06b0 00 00 41 b8 01 00 00 00 4d 89 c1 0f 1f 44 00 00  ..A.....M....D..
	0x06c0 e8 00 00 00 00 48 8d 05 00 00 00 00 bb 01 00 00  .....H..........
	0x06d0 00 b9 40 00 00 00 e8 00 00 00 00 0f 1f 44 00 00  ..@..........D..
	0x06e0 e8 00 00 00 00 44 0f 11 bc 24 28 01 00 00 48 8d  .....D...$(...H.
	0x06f0 15 00 00 00 00 48 89 94 24 28 01 00 00 48 89 84  .....H..$(...H..
	0x0700 24 30 01 00 00 48 8b 1d 00 00 00 00 48 8d 05 00  $0...H......H...
	0x0710 00 00 00 48 8d 8c 24 28 01 00 00 bf 01 00 00 00  ...H..$(........
	0x0720 48 89 fe e8 00 00 00 00 48 8d 05 00 00 00 00 e8  H.......H.......
	0x0730 00 00 00 00 48 89 84 24 88 00 00 00 48 c7 00 01  ....H..$....H...
	0x0740 00 00 00 48 c7 40 08 02 00 00 00 48 c7 40 10 03  ...H.@.....H.@..
	0x0750 00 00 00 48 c7 40 18 04 00 00 00 48 c7 40 20 05  ...H.@.....H.@ .
	0x0760 00 00 00 48 8d 05 00 00 00 00 e8 00 00 00 00 48  ...H...........H
	0x0770 89 84 24 80 00 00 00 83 3d 00 00 00 00 00 66 90  ..$.....=.....f.
	0x0780 75 0d 48 8b 8c 24 88 00 00 00 48 89 08 eb 10 48  u.H..$....H....H
	0x0790 89 c7 48 8b 8c 24 88 00 00 00 e8 00 00 00 00 48  ..H..$.........H
	0x07a0 8b 11 48 89 54 24 58 0f 10 41 08 0f 11 44 24 60  ..H.T$X..A...D$`
	0x07b0 0f 10 41 18 0f 11 44 24 70 48 8d 05 00 00 00 00  ..A...D$pH......
	0x07c0 48 8d 5c 24 58 e8 00 00 00 00 44 0f 11 bc 24 18  H.\$X.....D...$.
	0x07d0 01 00 00 48 89 84 24 18 01 00 00 48 89 9c 24 20  ...H..$....H..$ 
	0x07e0 01 00 00 48 8b 1d 00 00 00 00 48 8d 05 00 00 00  ...H......H.....
	0x07f0 00 48 8d 0d 00 00 00 00 bf 06 00 00 00 48 8d b4  .H...........H..
	0x0800 24 18 01 00 00 41 b8 01 00 00 00 4d 89 c1 e8 00  $....A.....M....
	0x0810 00 00 00 48 8b 8c 24 80 00 00 00 48 8b 11 44 0f  ...H..$....H..D.
	0x0820 11 bc 24 08 01 00 00 48 8d 35 00 00 00 00 48 89  ..$....H.5....H.
	0x0830 b4 24 08 01 00 00 48 89 94 24 10 01 00 00 48 8b  .$....H..$....H.
	0x0840 1d 00 00 00 00 48 8d 05 00 00 00 00 bf 06 00 00  .....H..........
	0x0850 00 41 b8 01 00 00 00 4d 89 c1 48 8d 0d 00 00 00  .A.....M..H.....
	0x0860 00 48 8d b4 24 08 01 00 00 e8 00 00 00 00 44 0f  .H..$.........D.
	0x0870 11 bc 24 f8 00 00 00 48 8d 0d 00 00 00 00 48 89  ..$....H......H.
	0x0880 8c 24 f8 00 00 00 48 8b 94 24 88 00 00 00 48 89  .$....H..$....H.
	0x0890 94 24 00 01 00 00 48 8b 1d 00 00 00 00 48 8d 05  .$....H......H..
	0x08a0 00 00 00 00 bf 07 00 00 00 48 8d b4 24 f8 00 00  .........H..$...
	0x08b0 00 41 b8 01 00 00 00 4d 89 c1 48 8d 0d 00 00 00  .A.....M..H.....
	0x08c0 00 e8 00 00 00 00 48 8b 8c 24 80 00 00 00 48 8b  ......H..$....H.
	0x08d0 11 44 0f 11 bc 24 e8 00 00 00 48 8d 35 00 00 00  .D...$....H.5...
	0x08e0 00 48 89 b4 24 e8 00 00 00 48 89 94 24 f0 00 00  .H..$....H..$...
	0x08f0 00 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 bf  .H......H.......
	0x0900 40 00 00 00 41 b8 01 00 00 00 4d 89 c1 48 8d 0d  @...A.....M..H..
	0x0910 00 00 00 00 48 8d b4 24 e8 00 00 00 0f 1f 40 00  ....H..$......@.
	0x0920 e8 00 00 00 00 48 8b 8c 24 80 00 00 00 48 8b 11  .....H..$....H..
	0x0930 44 0f 11 bc 24 d8 00 00 00 48 8d 35 00 00 00 00  D...$....H.5....
	0x0940 48 89 b4 24 d8 00 00 00 48 89 94 24 e0 00 00 00  H..$....H..$....
	0x0950 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 bf 6b  H......H.......k
	0x0960 00 00 00 48 8d b4 24 d8 00 00 00 41 b8 01 00 00  ...H..$....A....
	0x0970 00 4d 89 c1 48 8d 0d 00 00 00 00 0f 1f 44 00 00  .M..H........D..
	0x0980 e8 00 00 00 00 44 0f 11 bc 24 c8 00 00 00 48 8d  .....D...$....H.
	0x0990 0d 00 00 00 00 48 89 8c 24 c8 00 00 00 48 8b 8c  .....H..$....H..
	0x09a0 24 80 00 00 00 48 89 8c 24 d0 00 00 00 48 8b 1d  $....H..$....H..
	0x09b0 00 00 00 00 48 8d 05 00 00 00 00 bf 45 00 00 00  ....H.......E...
	0x09c0 48 8d b4 24 c8 00 00 00 41 b8 01 00 00 00 4d 89  H..$....A.....M.
	0x09d0 c1 48 8d 0d 00 00 00 00 e8 00 00 00 00 48 8b 8c  .H...........H..
	0x09e0 24 80 00 00 00 48 8b 19 84 03 48 8d 05 00 00 00  $....H....H.....
	0x09f0 00 e8 00 00 00 00 44 0f 11 bc 24 b8 00 00 00 48  ......D...$....H
	0x0a00 89 84 24 b8 00 00 00 48 89 9c 24 c0 00 00 00 48  ..$....H..$....H
	0x0a10 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d  ......H......H..
	0x0a20 00 00 00 00 bf 3b 00 00 00 48 8d b4 24 b8 00 00  .....;...H..$...
	0x0a30 00 41 b8 01 00 00 00 4d 89 c1 e8 00 00 00 00 48  .A.....M.......H
	0x0a40 8d 05 00 00 00 00 bb 01 00 00 00 b9 40 00 00 00  ............@...
	0x0a50 e8 00 00 00 00 e8 00 00 00 00 44 0f 11 bc 24 a8  ..........D...$.
	0x0a60 00 00 00 48 8d 0d 00 00 00 00 48 89 8c 24 a8 00  ...H......H..$..
	0x0a70 00 00 48 89 84 24 b0 00 00 00 48 8b 1d 00 00 00  ..H..$....H.....
	0x0a80 00 48 8d 05 00 00 00 00 48 8d 8c 24 a8 00 00 00  .H......H..$....
	0x0a90 bf 01 00 00 00 48 89 fe e8 00 00 00 00 48 8b ac  .....H.......H..
	0x0aa0 24 28 02 00 00 48 81 c4 30 02 00 00 c3 e8 00 00  $(...H..0.......
	0x0ab0 00 00 e9 49 f5 ff ff                             ...I...
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.**[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[4]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.*[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.**[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.*os.File+0
	rel 44+4 t=15 type.[4]int64+0
	rel 49+4 t=7 runtime.newobject+0
	rel 95+4 t=15 type.[4]int64+0
	rel 100+4 t=7 runtime.newobject+0
	rel 137+4 t=15 type.*[4]int64+0
	rel 142+4 t=7 runtime.newobject+0
	rel 156+4 t=15 runtime.writeBarrier+-1
	rel 193+4 t=7 runtime.gcWriteBarrierDX+0
	rel 200+4 t=15 go.string."="+0
	rel 215+4 t=7 strings.Repeat+0
	rel 225+4 t=7 runtime.convTstring+0
	rel 241+4 t=15 type.string+0
	rel 264+4 t=15 os.Stdout+0
	rel 271+4 t=15 go.itab.*os.File,io.Writer+0
	rel 292+4 t=7 fmt.Fprintln+0
	rel 324+4 t=15 type.[4]int64+0
	rel 334+4 t=7 runtime.convT2Enoptr+0
	rel 366+4 t=15 os.Stdout+0
	rel 373+4 t=15 go.itab.*os.File,io.Writer+0
	rel 380+4 t=15 go.string."a: %v\n"+0
	rel 407+4 t=7 fmt.Fprintf+0
	rel 439+4 t=15 type.[4]int64+0
	rel 449+4 t=7 runtime.convT2Enoptr+0
	rel 481+4 t=15 os.Stdout+0
	rel 488+4 t=15 go.itab.*os.File,io.Writer+0
	rel 495+4 t=15 go.string."b: %v\n"+0
	rel 522+4 t=7 fmt.Fprintf+0
	rel 562+4 t=15 type.[4]int64+0
	rel 577+4 t=7 runtime.convT2Enoptr+0
	rel 609+4 t=15 os.Stdout+0
	rel 616+4 t=15 go.itab.*os.File,io.Writer+0
	rel 623+4 t=15 go.string."b修改后: %v\n"+0
	rel 650+4 t=7 fmt.Fprintf+0
	rel 682+4 t=15 type.[4]int64+0
	rel 692+4 t=7 runtime.convT2Enoptr+0
	rel 724+4 t=15 os.Stdout+0
	rel 731+4 t=15 go.itab.*os.File,io.Writer+0
	rel 738+4 t=15 go.string."b修改后， a的值: %v\n"+0
	rel 769+4 t=7 fmt.Fprintf+0
	rel 785+4 t=15 type.*[4]int64+0
	rel 816+4 t=15 os.Stdout+0
	rel 823+4 t=15 go.itab.*os.File,io.Writer+0
	rel 830+4 t=15 go.string."a 的内存地址: %p\n"+0
	rel 857+4 t=7 fmt.Fprintf+0
	rel 873+4 t=15 type.*[4]int64+0
	rel 904+4 t=15 os.Stdout+0
	rel 911+4 t=15 go.itab.*os.File,io.Writer+0
	rel 918+4 t=15 go.string."b 的内存地址: %p\n"+0
	rel 945+4 t=7 fmt.Fprintf+0
	rel 952+4 t=15 go.string."-"+0
	rel 967+4 t=7 strings.Repeat+0
	rel 972+4 t=7 runtime.convTstring+0
	rel 988+4 t=15 type.string+0
	rel 1011+4 t=15 os.Stdout+0
	rel 1018+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1039+4 t=7 fmt.Fprintln+0
	rel 1070+4 t=15 type.[4]int64+0
	rel 1075+4 t=7 runtime.convT2Enoptr+0
	rel 1107+4 t=15 os.Stdout+0
	rel 1114+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1121+4 t=15 go.string."c: %v\n"+0
	rel 1153+4 t=7 fmt.Fprintf+0
	rel 1185+4 t=15 type.[4]int64+0
	rel 1195+4 t=7 runtime.convT2Enoptr+0
	rel 1227+4 t=15 os.Stdout+0
	rel 1234+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1241+4 t=15 go.string."a: %v\n"+0
	rel 1268+4 t=7 fmt.Fprintf+0
	rel 1284+4 t=15 type.*[4]int64+0
	rel 1315+4 t=15 os.Stdout+0
	rel 1322+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1329+4 t=15 go.string."&a: %p\n"+0
	rel 1356+4 t=7 fmt.Fprintf+0
	rel 1383+4 t=15 type.*[4]int64+0
	rel 1406+4 t=15 os.Stdout+0
	rel 1413+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1420+4 t=15 go.string."c: %p，以%%p打印会打印出它的值：数组a的首地址\n"+0
	rel 1447+4 t=7 fmt.Fprintf+0
	rel 1474+4 t=15 type.*[4]int64+0
	rel 1497+4 t=15 os.Stdout+0
	rel 1504+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1511+4 t=15 go.string..gostring.107.fc25da9a357bd0ef5c90db92d19f19d232fe9dcb48a70860742a11e5d80e44c2+0
	rel 1538+4 t=7 fmt.Fprintf+0
	rel 1554+4 t=15 type.**[4]int64+0
	rel 1585+4 t=15 os.Stdout+0
	rel 1592+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1599+4 t=15 go.string."&c: %p，变量c在栈上的地址，它的内容时指向a的指针\n"+0
	rel 1626+4 t=7 fmt.Fprintf+0
	rel 1646+4 t=15 type.[4]int64+0
	rel 1651+4 t=7 runtime.convT2Enoptr+0
	rel 1683+4 t=15 os.Stdout+0
	rel 1690+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1697+4 t=15 go.string."*c: %v，打印*c会打出c指向的内容，即a的内容\n"+0
	rel 1729+4 t=7 fmt.Fprintf+0
	rel 1736+4 t=15 go.string."-"+0
	rel 1751+4 t=7 strings.Repeat+0
	rel 1761+4 t=7 runtime.convTstring+0
	rel 1777+4 t=15 type.string+0
	rel 1800+4 t=15 os.Stdout+0
	rel 1807+4 t=15 go.itab.*os.File,io.Writer+0
	rel 1828+4 t=7 fmt.Fprintln+0
	rel 1835+4 t=15 type.[5]int64+0
	rel 1840+4 t=7 runtime.newobject+0
	rel 1894+4 t=15 type.*[5]int64+0
	rel 1899+4 t=7 runtime.newobject+0
	rel 1913+4 t=15 runtime.writeBarrier+-1
	rel 1947+4 t=7 runtime.gcWriteBarrierCX+0
	rel 1980+4 t=15 type.[5]int64+0
	rel 1990+4 t=7 runtime.convT2Enoptr+0
	rel 2022+4 t=15 os.Stdout+0
	rel 2029+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2036+4 t=15 go.string."d: %v\n"+0
	rel 2063+4 t=7 fmt.Fprintf+0
	rel 2090+4 t=15 type.*[5]int64+0
	rel 2113+4 t=15 os.Stdout+0
	rel 2120+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2141+4 t=15 go.string."f: %v\n"+0
	rel 2154+4 t=7 fmt.Fprintf+0
	rel 2170+4 t=15 type.*[5]int64+0
	rel 2201+4 t=15 os.Stdout+0
	rel 2208+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2237+4 t=15 go.string."&d: %p\n"+0
	rel 2242+4 t=7 fmt.Fprintf+0
	rel 2269+4 t=15 type.*[5]int64+0
	rel 2292+4 t=15 os.Stdout+0
	rel 2299+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2320+4 t=15 go.string."f: %p，以%%p打印会打印出它的值：数组d的首地址\n"+0
	rel 2337+4 t=7 fmt.Fprintf+0
	rel 2364+4 t=15 type.*[5]int64+0
	rel 2387+4 t=15 os.Stdout+0
	rel 2394+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2423+4 t=15 go.string..gostring.107.958de85d687a58aae8b5f361b1c484f3c0d64dab64309d7e9d0b8869bff01f62+0
	rel 2433+4 t=7 fmt.Fprintf+0
	rel 2449+4 t=15 type.**[5]int64+0
	rel 2480+4 t=15 os.Stdout+0
	rel 2487+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2516+4 t=15 go.string."&f: %p，变量f在栈上的地址，它的内容是指向d的指针\n"+0
	rel 2521+4 t=7 fmt.Fprintf+0
	rel 2541+4 t=15 type.[5]int64+0
	rel 2546+4 t=7 runtime.convT2Enoptr+0
	rel 2578+4 t=15 os.Stdout+0
	rel 2585+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2592+4 t=15 go.string."*f: %v，打印*f会打出f指向的内容，即d的内容\n"+0
	rel 2619+4 t=7 fmt.Fprintf+0
	rel 2626+4 t=15 go.string."="+0
	rel 2641+4 t=7 strings.Repeat+0
	rel 2646+4 t=7 runtime.convTstring+0
	rel 2662+4 t=15 type.string+0
	rel 2685+4 t=15 os.Stdout+0
	rel 2692+4 t=15 go.itab.*os.File,io.Writer+0
	rel 2713+4 t=7 fmt.Fprintln+0
	rel 2734+4 t=7 runtime.morestack_noctxt+0
"".array_pointer2 STEXT size=234 args=0x0 locals=0x68 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:141)	TEXT	"".array_pointer2(SB), ABIInternal, $104-0
	0x0000 00000 (06-data_types/07-array.go:141)	CMPQ	SP, 16(R14)
	0x0004 00004 (06-data_types/07-array.go:141)	PCDATA	$0, $-2
	0x0004 00004 (06-data_types/07-array.go:141)	JLS	223
	0x000a 00010 (06-data_types/07-array.go:141)	PCDATA	$0, $-1
	0x000a 00010 (06-data_types/07-array.go:141)	SUBQ	$104, SP
	0x000e 00014 (06-data_types/07-array.go:141)	MOVQ	BP, 96(SP)
	0x0013 00019 (06-data_types/07-array.go:141)	LEAQ	96(SP), BP
	0x0018 00024 (06-data_types/07-array.go:141)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0018 00024 (06-data_types/07-array.go:141)	FUNCDATA	$1, gclocals·d527b79a98f329c2ba624a68e7df03d6(SB)
	0x0018 00024 (06-data_types/07-array.go:141)	FUNCDATA	$2, "".array_pointer2.stkobj(SB)
	0x0018 00024 (06-data_types/07-array.go:142)	LEAQ	type.[3]int64(SB), AX
	0x001f 00031 (06-data_types/07-array.go:142)	PCDATA	$1, $0
	0x001f 00031 (06-data_types/07-array.go:142)	NOP
	0x0020 00032 (06-data_types/07-array.go:142)	CALL	runtime.newobject(SB)
	0x0025 00037 (06-data_types/07-array.go:142)	MOVQ	AX, "".&a+56(SP)
	0x002a 00042 (06-data_types/07-array.go:142)	MOVQ	$1, (AX)
	0x0031 00049 (06-data_types/07-array.go:142)	MOVQ	$2, 8(AX)
	0x0039 00057 (06-data_types/07-array.go:142)	MOVQ	$3, 16(AX)
	0x0041 00065 (06-data_types/07-array.go:143)	LEAQ	type.[5]int64(SB), AX
	0x0048 00072 (06-data_types/07-array.go:143)	PCDATA	$1, $1
	0x0048 00072 (06-data_types/07-array.go:143)	CALL	runtime.newobject(SB)
	0x004d 00077 (06-data_types/07-array.go:143)	MOVQ	$1, (AX)
	0x0054 00084 (06-data_types/07-array.go:143)	MOVQ	$2, 8(AX)
	0x005c 00092 (06-data_types/07-array.go:143)	MOVQ	$3, 16(AX)
	0x0064 00100 (06-data_types/07-array.go:143)	MOVQ	$4, 24(AX)
	0x006c 00108 (06-data_types/07-array.go:143)	MOVQ	$5, 32(AX)
	0x0074 00116 (06-data_types/07-array.go:144)	LEAQ	""..autotmp_15+64(SP), SI
	0x0079 00121 (06-data_types/07-array.go:144)	MOVUPS	X15, (SI)
	0x007d 00125 (06-data_types/07-array.go:144)	LEAQ	""..autotmp_15+80(SP), CX
	0x0082 00130 (06-data_types/07-array.go:144)	MOVUPS	X15, (CX)
	0x0086 00134 (06-data_types/07-array.go:144)	LEAQ	type.*[3]int64(SB), CX
	0x008d 00141 (06-data_types/07-array.go:144)	MOVQ	CX, ""..autotmp_15+64(SP)
	0x0092 00146 (06-data_types/07-array.go:144)	MOVQ	"".&a+56(SP), CX
	0x0097 00151 (06-data_types/07-array.go:144)	MOVQ	CX, ""..autotmp_15+72(SP)
	0x009c 00156 (06-data_types/07-array.go:144)	LEAQ	type.*[5]int64(SB), CX
	0x00a3 00163 (06-data_types/07-array.go:144)	MOVQ	CX, ""..autotmp_15+80(SP)
	0x00a8 00168 (06-data_types/07-array.go:144)	MOVQ	AX, ""..autotmp_15+88(SP)
	0x00ad 00173 (<unknown line number>)	NOP
	0x00ad 00173 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x00b4 00180 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00bb 00187 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."&a: %p, &d: %p\n"(SB), CX
	0x00c2 00194 ($GOROOT/src/fmt/print.go:213)	MOVL	$15, DI
	0x00c7 00199 ($GOROOT/src/fmt/print.go:213)	MOVL	$2, R8
	0x00cd 00205 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x00d0 00208 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $0
	0x00d0 00208 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x00d5 00213 (06-data_types/07-array.go:149)	MOVQ	96(SP), BP
	0x00da 00218 (06-data_types/07-array.go:149)	ADDQ	$104, SP
	0x00de 00222 (06-data_types/07-array.go:149)	RET
	0x00df 00223 (06-data_types/07-array.go:149)	NOP
	0x00df 00223 (06-data_types/07-array.go:141)	PCDATA	$1, $-1
	0x00df 00223 (06-data_types/07-array.go:141)	PCDATA	$0, $-2
	0x00df 00223 (06-data_types/07-array.go:141)	NOP
	0x00e0 00224 (06-data_types/07-array.go:141)	CALL	runtime.morestack_noctxt(SB)
	0x00e5 00229 (06-data_types/07-array.go:141)	PCDATA	$0, $-1
	0x00e5 00229 (06-data_types/07-array.go:141)	JMP	0
	0x0000 49 3b 66 10 0f 86 d5 00 00 00 48 83 ec 68 48 89  I;f.......H..hH.
	0x0010 6c 24 60 48 8d 6c 24 60 48 8d 05 00 00 00 00 90  l$`H.l$`H.......
	0x0020 e8 00 00 00 00 48 89 44 24 38 48 c7 00 01 00 00  .....H.D$8H.....
	0x0030 00 48 c7 40 08 02 00 00 00 48 c7 40 10 03 00 00  .H.@.....H.@....
	0x0040 00 48 8d 05 00 00 00 00 e8 00 00 00 00 48 c7 00  .H...........H..
	0x0050 01 00 00 00 48 c7 40 08 02 00 00 00 48 c7 40 10  ....H.@.....H.@.
	0x0060 03 00 00 00 48 c7 40 18 04 00 00 00 48 c7 40 20  ....H.@.....H.@ 
	0x0070 05 00 00 00 48 8d 74 24 40 44 0f 11 3e 48 8d 4c  ....H.t$@D..>H.L
	0x0080 24 50 44 0f 11 39 48 8d 0d 00 00 00 00 48 89 4c  $PD..9H......H.L
	0x0090 24 40 48 8b 4c 24 38 48 89 4c 24 48 48 8d 0d 00  $@H.L$8H.L$HH...
	0x00a0 00 00 00 48 89 4c 24 50 48 89 44 24 58 48 8b 1d  ...H.L$PH.D$XH..
	0x00b0 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00  ....H......H....
	0x00c0 00 00 bf 0f 00 00 00 41 b8 02 00 00 00 4d 89 c1  .......A.....M..
	0x00d0 e8 00 00 00 00 48 8b 6c 24 60 48 83 c4 68 c3 90  .....H.l$`H..h..
	0x00e0 e8 00 00 00 00 e9 16 ff ff ff                    ..........
	rel 3+0 t=24 type.*[3]int64+0
	rel 3+0 t=24 type.*[5]int64+0
	rel 3+0 t=24 type.*os.File+0
	rel 27+4 t=15 type.[3]int64+0
	rel 33+4 t=7 runtime.newobject+0
	rel 68+4 t=15 type.[5]int64+0
	rel 73+4 t=7 runtime.newobject+0
	rel 137+4 t=15 type.*[3]int64+0
	rel 159+4 t=15 type.*[5]int64+0
	rel 176+4 t=15 os.Stdout+0
	rel 183+4 t=15 go.itab.*os.File,io.Writer+0
	rel 190+4 t=15 go.string."&a: %p, &d: %p\n"+0
	rel 209+4 t=7 fmt.Fprintf+0
	rel 225+4 t=7 runtime.morestack_noctxt+0
"".main STEXT size=40 args=0x0 locals=0x8 funcid=0x0
	0x0000 00000 (06-data_types/07-array.go:151)	TEXT	"".main(SB), ABIInternal, $8-0
	0x0000 00000 (06-data_types/07-array.go:151)	CMPQ	SP, 16(R14)
	0x0004 00004 (06-data_types/07-array.go:151)	PCDATA	$0, $-2
	0x0004 00004 (06-data_types/07-array.go:151)	JLS	33
	0x0006 00006 (06-data_types/07-array.go:151)	PCDATA	$0, $-1
	0x0006 00006 (06-data_types/07-array.go:151)	SUBQ	$8, SP
	0x000a 00010 (06-data_types/07-array.go:151)	MOVQ	BP, (SP)
	0x000e 00014 (06-data_types/07-array.go:151)	LEAQ	(SP), BP
	0x0012 00018 (06-data_types/07-array.go:151)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (06-data_types/07-array.go:151)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (06-data_types/07-array.go:156)	PCDATA	$1, $0
	0x0012 00018 (06-data_types/07-array.go:156)	CALL	"".array_pointer2(SB)
	0x0017 00023 (06-data_types/07-array.go:157)	MOVQ	(SP), BP
	0x001b 00027 (06-data_types/07-array.go:157)	ADDQ	$8, SP
	0x001f 00031 (06-data_types/07-array.go:157)	NOP
	0x0020 00032 (06-data_types/07-array.go:157)	RET
	0x0021 00033 (06-data_types/07-array.go:157)	NOP
	0x0021 00033 (06-data_types/07-array.go:151)	PCDATA	$1, $-1
	0x0021 00033 (06-data_types/07-array.go:151)	PCDATA	$0, $-2
	0x0021 00033 (06-data_types/07-array.go:151)	CALL	runtime.morestack_noctxt(SB)
	0x0026 00038 (06-data_types/07-array.go:151)	PCDATA	$0, $-1
	0x0026 00038 (06-data_types/07-array.go:151)	JMP	0
	0x0000 49 3b 66 10 76 1b 48 83 ec 08 48 89 2c 24 48 8d  I;f.v.H...H.,$H.
	0x0010 2c 24 e8 00 00 00 00 48 8b 2c 24 48 83 c4 08 90  ,$.....H.,$H....
	0x0020 c3 e8 00 00 00 00 eb d8                          ........
	rel 19+4 t=7 "".array_pointer2+0
	rel 34+4 t=7 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok size=86 args=0x8 locals=0x10 funcid=0x16
	0x0000 00000 (<autogenerated>:1)	TEXT	os.(*File).close(SB), DUPOK|WRAPPER|ABIInternal, $16-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	52
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	69
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, os.(*File).close.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0025 00037 (<autogenerated>:1)	PCDATA	$1, $1
	0x0025 00037 (<autogenerated>:1)	CALL	os.(*file).close(SB)
	0x002a 00042 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x002f 00047 (<autogenerated>:1)	ADDQ	$16, SP
	0x0033 00051 (<autogenerated>:1)	RET
	0x0034 00052 (<autogenerated>:1)	NOP
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0034 00052 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0034 00052 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0039 00057 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x003e 00062 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0043 00067 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0043 00067 (<autogenerated>:1)	JMP	0
	0x0045 00069 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x004a 00074 (<autogenerated>:1)	CMPQ	(R12), R13
	0x004e 00078 (<autogenerated>:1)	JNE	29
	0x0050 00080 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0054 00084 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 2e 48 83 ec 10 48 89 6c 24 08 48  I;f.v.H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 28 48 89 44  .l$.M.f M..u(H.D
	0x0020 24 18 48 8b 00 e8 00 00 00 00 48 8b 6c 24 08 48  $.H.......H.l$.H
	0x0030 83 c4 10 c3 48 89 44 24 08 e8 00 00 00 00 48 8b  ....H.D$......H.
	0x0040 44 24 08 eb bb 4c 8d 6c 24 18 4d 39 2c 24 75 cd  D$...L.l$.M9,$u.
	0x0050 49 89 24 24 eb c7                                I.$$..
	rel 38+4 t=7 os.(*file).close+0
	rel 58+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=40
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 fmt..inittask+0
	rel 32+8 t=1 strings..inittask+0
go.info.fmt.Printf$abstract SDWARFABSFCN dupok size=54
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 66 00 01 01 11 66  .fmt.Printf....f
	0x0010 6f 72 6d 61 74 00 00 00 00 00 00 11 61 00 00 00  ormat.......a...
	0x0020 00 00 00 11 6e 00 01 00 00 00 00 11 65 72 72 00  ....n.......err.
	0x0030 01 00 00 00 00 00                                ......
	rel 0+0 t=23 type.[]interface {}+0
	rel 0+0 t=23 type.error+0
	rel 0+0 t=23 type.int+0
	rel 0+0 t=23 type.string+0
	rel 23+4 t=31 go.info.string+0
	rel 31+4 t=31 go.info.[]interface {}+0
	rel 39+4 t=31 go.info.int+0
	rel 49+4 t=31 go.info.error+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=23 type.[]interface {}+0
	rel 0+0 t=23 type.error+0
	rel 0+0 t=23 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
go.string."a: %v\n" SRODATA dupok size=6
	0x0000 61 3a 20 25 76 0a                                a: %v.
go.string."b: %v\n" SRODATA dupok size=6
	0x0000 62 3a 20 25 76 0a                                b: %v.
go.string."c: %v\n" SRODATA dupok size=6
	0x0000 63 3a 20 25 76 0a                                c: %v.
go.string."%T, d: %[1]v\n" SRODATA dupok size=13
	0x0000 25 54 2c 20 64 3a 20 25 5b 31 5d 76 0a           %T, d: %[1]v.
go.string."e: %v\n" SRODATA dupok size=6
	0x0000 65 3a 20 25 76 0a                                e: %v.
go.string."index: %d\tvalue: %d\n" SRODATA dupok size=20
	0x0000 69 6e 64 65 78 3a 20 25 64 09 76 61 6c 75 65 3a  index: %d.value:
	0x0010 20 25 64 0a                                       %d.
go.string."f type: %T\t, value: %[1]v\n" SRODATA dupok size=26
	0x0000 66 20 74 79 70 65 3a 20 25 54 09 2c 20 76 61 6c  f type: %T., val
	0x0010 75 65 3a 20 25 5b 31 5d 76 0a                    ue: %[1]v.
go.string."=" SRODATA dupok size=1
	0x0000 3d                                               =
go.string."b修改后: %v\n" SRODATA dupok size=15
	0x0000 62 e4 bf ae e6 94 b9 e5 90 8e 3a 20 25 76 0a     b.........: %v.
go.string."b修改后， a的值: %v\n" SRODATA dupok size=26
	0x0000 62 e4 bf ae e6 94 b9 e5 90 8e ef bc 8c 20 61 e7  b............ a.
	0x0010 9a 84 e5 80 bc 3a 20 25 76 0a                    .....: %v.
go.string."a 的内存地址: %p\n" SRODATA dupok size=22
	0x0000 61 20 e7 9a 84 e5 86 85 e5 ad 98 e5 9c b0 e5 9d  a ..............
	0x0010 80 3a 20 25 70 0a                                .: %p.
go.string."b 的内存地址: %p\n" SRODATA dupok size=22
	0x0000 62 20 e7 9a 84 e5 86 85 e5 ad 98 e5 9c b0 e5 9d  b ..............
	0x0010 80 3a 20 25 70 0a                                .: %p.
go.string."-" SRODATA dupok size=1
	0x0000 2d                                               -
go.string."&a: %p\n" SRODATA dupok size=7
	0x0000 26 61 3a 20 25 70 0a                             &a: %p.
go.string."c: %p，以%%p打印会打印出它的值：数组a的首地址\n" SRODATA dupok size=64
	0x0000 63 3a 20 25 70 ef bc 8c e4 bb a5 25 25 70 e6 89  c: %p......%%p..
	0x0010 93 e5 8d b0 e4 bc 9a e6 89 93 e5 8d b0 e5 87 ba  ................
	0x0020 e5 ae 83 e7 9a 84 e5 80 bc ef bc 9a e6 95 b0 e7  ................
	0x0030 bb 84 61 e7 9a 84 e9 a6 96 e5 9c b0 e5 9d 80 0a  ..a.............
go.string..gostring.107.fc25da9a357bd0ef5c90db92d19f19d232fe9dcb48a70860742a11e5d80e44c2 SRODATA dupok size=107
	0x0000 63 3a 20 25 76 ef bc 8c e4 bb a5 25 25 76 e6 89  c: %v......%%v..
	0x0010 93 e5 8d b0 e4 bc 9a e6 89 93 e5 8d b0 e5 87 ba  ................
	0x0020 e5 ae 83 e5 bc 95 e7 94 a8 e7 9a 84 e5 80 bc ef  ................
	0x0030 bc 9a e6 95 b0 e7 bb 84 61 e7 9a 84 e5 86 85 e5  ........a.......
	0x0040 ae b9 e3 80 82 e7 94 a8 26 e4 bb a3 e8 a1 a8 e5  ........&.......
	0x0050 bc 95 e7 94 a8 ef bc 8c e4 b8 8d e4 bc 9a e6 98  ................
	0x0060 be e7 a4 ba e6 8c 87 e9 92 88 0a                 ...........
go.string."&c: %p，变量c在栈上的地址，它的内容时指向a的指针\n" SRODATA dupok size=69
	0x0000 26 63 3a 20 25 70 ef bc 8c e5 8f 98 e9 87 8f 63  &c: %p.........c
	0x0010 e5 9c a8 e6 a0 88 e4 b8 8a e7 9a 84 e5 9c b0 e5  ................
	0x0020 9d 80 ef bc 8c e5 ae 83 e7 9a 84 e5 86 85 e5 ae  ................
	0x0030 b9 e6 97 b6 e6 8c 87 e5 90 91 61 e7 9a 84 e6 8c  ..........a.....
	0x0040 87 e9 92 88 0a                                   .....
go.string."*c: %v，打印*c会打出c指向的内容，即a的内容\n" SRODATA dupok size=59
	0x0000 2a 63 3a 20 25 76 ef bc 8c e6 89 93 e5 8d b0 2a  *c: %v.........*
	0x0010 63 e4 bc 9a e6 89 93 e5 87 ba 63 e6 8c 87 e5 90  c.........c.....
	0x0020 91 e7 9a 84 e5 86 85 e5 ae b9 ef bc 8c e5 8d b3  ................
	0x0030 61 e7 9a 84 e5 86 85 e5 ae b9 0a                 a..........
go.string."d: %v\n" SRODATA dupok size=6
	0x0000 64 3a 20 25 76 0a                                d: %v.
go.string."f: %v\n" SRODATA dupok size=6
	0x0000 66 3a 20 25 76 0a                                f: %v.
go.string."&d: %p\n" SRODATA dupok size=7
	0x0000 26 64 3a 20 25 70 0a                             &d: %p.
go.string."f: %p，以%%p打印会打印出它的值：数组d的首地址\n" SRODATA dupok size=64
	0x0000 66 3a 20 25 70 ef bc 8c e4 bb a5 25 25 70 e6 89  f: %p......%%p..
	0x0010 93 e5 8d b0 e4 bc 9a e6 89 93 e5 8d b0 e5 87 ba  ................
	0x0020 e5 ae 83 e7 9a 84 e5 80 bc ef bc 9a e6 95 b0 e7  ................
	0x0030 bb 84 64 e7 9a 84 e9 a6 96 e5 9c b0 e5 9d 80 0a  ..d.............
go.string..gostring.107.958de85d687a58aae8b5f361b1c484f3c0d64dab64309d7e9d0b8869bff01f62 SRODATA dupok size=107
	0x0000 66 3a 20 25 76 ef bc 8c e4 bb a5 25 25 76 e6 89  f: %v......%%v..
	0x0010 93 e5 8d b0 e4 bc 9a e6 89 93 e5 8d b0 e5 87 ba  ................
	0x0020 e5 ae 83 e5 bc 95 e7 94 a8 e7 9a 84 e5 80 bc ef  ................
	0x0030 bc 9a e6 95 b0 e7 bb 84 64 e7 9a 84 e5 86 85 e5  ........d.......
	0x0040 ae b9 e3 80 82 e7 94 a8 26 e4 bb a3 e8 a1 a8 e5  ........&.......
	0x0050 bc 95 e7 94 a8 ef bc 8c e4 b8 8d e4 bc 9a e6 98  ................
	0x0060 be e7 a4 ba e6 8c 87 e9 92 88 0a                 ...........
go.string."&f: %p，变量f在栈上的地址，它的内容是指向d的指针\n" SRODATA dupok size=69
	0x0000 26 66 3a 20 25 70 ef bc 8c e5 8f 98 e9 87 8f 66  &f: %p.........f
	0x0010 e5 9c a8 e6 a0 88 e4 b8 8a e7 9a 84 e5 9c b0 e5  ................
	0x0020 9d 80 ef bc 8c e5 ae 83 e7 9a 84 e5 86 85 e5 ae  ................
	0x0030 b9 e6 98 af e6 8c 87 e5 90 91 64 e7 9a 84 e6 8c  ..........d.....
	0x0040 87 e9 92 88 0a                                   .....
go.string."*f: %v，打印*f会打出f指向的内容，即d的内容\n" SRODATA dupok size=59
	0x0000 2a 66 3a 20 25 76 ef bc 8c e6 89 93 e5 8d b0 2a  *f: %v.........*
	0x0010 66 e4 bc 9a e6 89 93 e5 87 ba 66 e6 8c 87 e5 90  f.........f.....
	0x0020 91 e7 9a 84 e5 86 85 e5 ae b9 ef bc 8c e5 8d b3  ................
	0x0030 64 e7 9a 84 e5 86 85 e5 ae b9 0a                 d..........
go.string."&a: %p, &d: %p\n" SRODATA dupok size=15
	0x0000 26 61 3a 20 25 70 2c 20 26 64 3a 20 25 70 0a     &a: %p, &d: %p.
""..stmp_0 SRODATA static size=40
	0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 03 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0020 05 00 00 00 00 00 00 00                          ........
""..stmp_1 SRODATA static size=32
	0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 03 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
""..stmp_2 SRODATA static size=40
	0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 03 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0020 05 00 00 00 00 00 00 00                          ........
""..stmp_3 SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 03 00 00 00 00 00 00 00                          ........
""..stmp_4 SRODATA static size=40
	0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 03 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0020 05 00 00 00 00 00 00 00                          ........
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]int64- SRODATA dupok size=10
	0x0000 00 08 2a 5b 5d 69 6e 74 36 34                    ..*[]int64
type.*[]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2c 4e ee 91 08 08 08 36 00 00 00 00 00 00 00 00  ,N.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int64-+0
	rel 48+8 t=1 type.[]int64+0
type.[]int64 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 96 8e 76 88 02 08 08 17 00 00 00 00 00 00 00 00  ..v.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int64-+0
	rel 44+4 t=-32763 type.*[]int64+0
	rel 48+8 t=1 type.int64+0
type..eqfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memequal_varlen+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*[4]int64- SRODATA dupok size=11
	0x0000 00 09 2a 5b 34 5d 69 6e 74 36 34                 ..*[4]int64
type.[4]int64 SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x0010 04 b6 ce b5 0a 08 08 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc32+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[4]int64-+0
	rel 44+4 t=-32763 type.*[4]int64+0
	rel 48+8 t=1 type.int64+0
	rel 56+8 t=1 type.[]int64+0
type.*[4]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 27 22 8e 0c 08 08 08 36 00 00 00 00 00 00 00 00  '".....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[4]int64-+0
	rel 44+4 t=-32763 type.**[4]int64+0
	rel 48+8 t=1 type.[4]int64+0
type..namedata.**[4]int64- SRODATA dupok size=12
	0x0000 00 0a 2a 2a 5b 34 5d 69 6e 74 36 34              ..**[4]int64
type.**[4]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 51 66 82 08 08 08 36 00 00 00 00 00 00 00 00  .Qf....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**[4]int64-+0
	rel 48+8 t=1 type.*[4]int64+0
type..eqfunc40 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[5]int64- SRODATA dupok size=11
	0x0000 00 09 2a 5b 35 5d 69 6e 74 36 34                 ..*[5]int64
type.[5]int64 SRODATA dupok size=72
	0x0000 28 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  (...............
	0x0010 2c 7e 64 df 0a 08 08 11 00 00 00 00 00 00 00 00  ,~d.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 05 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc40+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[5]int64-+0
	rel 44+4 t=-32763 type.*[5]int64+0
	rel 48+8 t=1 type.int64+0
	rel 56+8 t=1 type.[]int64+0
type.*[5]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 75 ff 00 e3 08 08 08 36 00 00 00 00 00 00 00 00  u......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[5]int64-+0
	rel 44+4 t=-32763 type.**[5]int64+0
	rel 48+8 t=1 type.[5]int64+0
type..namedata.**[5]int64- SRODATA dupok size=12
	0x0000 00 0a 2a 2a 5b 35 5d 69 6e 74 36 34              ..**[5]int64
type.**[5]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 50 f3 0a fd 08 08 08 36 00 00 00 00 00 00 00 00  P......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**[5]int64-+0
	rel 48+8 t=1 type.*[5]int64+0
type..eqfunc24 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[3]int64- SRODATA dupok size=11
	0x0000 00 09 2a 5b 33 5d 69 6e 74 36 34                 ..*[3]int64
type.[3]int64 SRODATA dupok size=72
	0x0000 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 1c ef 86 42 0a 08 08 11 00 00 00 00 00 00 00 00  ...B............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc24+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[3]int64-+0
	rel 44+4 t=-32763 type.*[3]int64+0
	rel 48+8 t=1 type.int64+0
	rel 56+8 t=1 type.[]int64+0
type.*[3]int64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 81 ee ae 6b 08 08 08 36 00 00 00 00 00 00 00 00  ...k...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]int64-+0
	rel 48+8 t=1 type.[3]int64+0
type..namedata.*[]int- SRODATA dupok size=8
	0x0000 00 06 2a 5b 5d 69 6e 74                          ..*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=-32763 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..eqfunc88 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 58 00 00 00 00 00 00 00  ........X.......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[11]int- SRODATA dupok size=10
	0x0000 00 08 2a 5b 31 31 5d 69 6e 74                    ..*[11]int
type.*[11]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8d cd 8b d3 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[11]int-+0
	rel 48+8 t=1 type.[11]int+0
type.[11]int SRODATA dupok size=72
	0x0000 58 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  X...............
	0x0010 40 ca 66 be 0a 08 08 11 00 00 00 00 00 00 00 00  @.f.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 0b 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc88+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[11]int-+0
	rel 44+4 t=-32763 type.*[11]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
type..namedata.*[3]int- SRODATA dupok size=9
	0x0000 00 07 2a 5b 33 5d 69 6e 74                       ..*[3]int
type.*[3]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 25 30 9a e3 08 08 08 36 00 00 00 00 00 00 00 00  %0.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]int-+0
	rel 48+8 t=1 type.[3]int+0
type.[3]int SRODATA dupok size=72
	0x0000 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 63 c4 1a 46 0a 08 08 11 00 00 00 00 00 00 00 00  c..F............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc24+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[3]int-+0
	rel 44+4 t=-32763 type.*[3]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
type..importpath.strings. SRODATA dupok size=9
	0x0000 00 07 73 74 72 69 6e 67 73                       ..strings
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9f1f24ff36424f3d0e5fd3a9c6a7b1cf SRODATA dupok size=14
	0x0000 02 00 00 00 15 00 00 00 00 00 00 01 00 00        ..............
"".array_def.stkobj SRODATA static size=200
	0x0000 08 00 00 00 00 00 00 00 60 ff ff ff 10 00 00 00  ........`.......
	0x0010 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 70 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  p...............
	0x0030 00 00 00 00 00 00 00 00 80 ff ff ff 10 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 90 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 a0 ff ff ff 10 00 00 00  ................
	0x0070 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 b0 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 c0 ff ff ff 20 00 00 00  ............ ...
	0x00a0 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x00b0 e0 ff ff ff 20 00 00 00 20 00 00 00 00 00 00 00  .... ... .......
	0x00c0 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.gcbits.02+0
	rel 48+8 t=1 runtime.gcbits.02+0
	rel 72+8 t=1 runtime.gcbits.02+0
	rel 96+8 t=1 runtime.gcbits.02+0
	rel 120+8 t=1 runtime.gcbits.02+0
	rel 144+8 t=1 runtime.gcbits.02+0
	rel 168+8 t=1 runtime.gcbits.0a+0
	rel 192+8 t=1 runtime.gcbits.0a+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
"".arry_len_le_4.stkobj SRODATA static size=32
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.gcbits.02+0
"".arry_len_more_4.stkobj SRODATA static size=32
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.gcbits.02+0
gclocals·fcf5af2016adf65a97b579a67730f1b6 SRODATA dupok size=8
	0x0000 09 00 00 00 00 00 00 00                          ........
gclocals·12b6f77e0db330efe75dfd918d12c72e SRODATA dupok size=71
	0x0000 09 00 00 00 35 00 00 00 00 00 00 00 00 00 00 10  ....5...........
	0x0010 00 00 00 00 00 00 18 00 00 00 00 00 00 1c 00 00  ................
	0x0020 00 00 00 00 14 00 00 00 00 00 00 04 00 00 00 00  ................
	0x0030 00 00 02 00 00 00 00 00 00 03 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00                             .......
"".array_pointer.stkobj SRODATA static size=584
	0x0000 18 00 00 00 00 00 00 00 80 fe ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 90 fe ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 a0 fe ff ff 10 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 b0 fe ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 c0 fe ff ff 10 00 00 00  ................
	0x0070 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 d0 fe ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 e0 fe ff ff 10 00 00 00  ................
	0x00a0 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00b0 f0 fe ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x00c0 00 00 00 00 00 00 00 00 00 ff ff ff 10 00 00 00  ................
	0x00d0 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00e0 10 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x00f0 00 00 00 00 00 00 00 00 20 ff ff ff 10 00 00 00  ........ .......
	0x0100 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0110 30 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  0...............
	0x0120 00 00 00 00 00 00 00 00 40 ff ff ff 10 00 00 00  ........@.......
	0x0130 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0140 50 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  P...............
	0x0150 00 00 00 00 00 00 00 00 60 ff ff ff 10 00 00 00  ........`.......
	0x0160 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0170 70 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  p...............
	0x0180 00 00 00 00 00 00 00 00 80 ff ff ff 10 00 00 00  ................
	0x0190 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x01a0 90 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x01b0 00 00 00 00 00 00 00 00 a0 ff ff ff 10 00 00 00  ................
	0x01c0 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x01d0 b0 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x01e0 00 00 00 00 00 00 00 00 c0 ff ff ff 10 00 00 00  ................
	0x01f0 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0200 d0 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0210 00 00 00 00 00 00 00 00 e0 ff ff ff 10 00 00 00  ................
	0x0220 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0230 f0 ff ff ff 10 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0240 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.gcbits.02+0
	rel 48+8 t=1 runtime.gcbits.02+0
	rel 72+8 t=1 runtime.gcbits.02+0
	rel 96+8 t=1 runtime.gcbits.02+0
	rel 120+8 t=1 runtime.gcbits.02+0
	rel 144+8 t=1 runtime.gcbits.02+0
	rel 168+8 t=1 runtime.gcbits.02+0
	rel 192+8 t=1 runtime.gcbits.02+0
	rel 216+8 t=1 runtime.gcbits.02+0
	rel 240+8 t=1 runtime.gcbits.02+0
	rel 264+8 t=1 runtime.gcbits.02+0
	rel 288+8 t=1 runtime.gcbits.02+0
	rel 312+8 t=1 runtime.gcbits.02+0
	rel 336+8 t=1 runtime.gcbits.02+0
	rel 360+8 t=1 runtime.gcbits.02+0
	rel 384+8 t=1 runtime.gcbits.02+0
	rel 408+8 t=1 runtime.gcbits.02+0
	rel 432+8 t=1 runtime.gcbits.02+0
	rel 456+8 t=1 runtime.gcbits.02+0
	rel 480+8 t=1 runtime.gcbits.02+0
	rel 504+8 t=1 runtime.gcbits.02+0
	rel 528+8 t=1 runtime.gcbits.02+0
	rel 552+8 t=1 runtime.gcbits.02+0
	rel 576+8 t=1 runtime.gcbits.02+0
gclocals·d527b79a98f329c2ba624a68e7df03d6 SRODATA dupok size=10
	0x0000 02 00 00 00 05 00 00 00 00 01                    ..........
"".array_pointer2.stkobj SRODATA static size=32
	0x0000 01 00 00 00 00 00 00 00 e0 ff ff ff 20 00 00 00  ............ ...
	0x0010 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	rel 24+8 t=1 runtime.gcbits.0a+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
os.(*File).close.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
