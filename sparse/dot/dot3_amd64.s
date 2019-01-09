//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot3(x []float64, indx []int, y []float64) (dot float64)
TEXT ·Dot3(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI
    MOVQ    y+48(FP), DX   
    MOVQ    indx+32(FP), AX
    
	XORL	R9, R9          
	XORPS	X0, X0          

    CMPQ	R9, AX
	JGE	    end

loop:
    MOVQ	(SI)(R9*8), R10 
	MOVSD	(R8)(R9*8), X1  
    INCQ	R9

    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    CMPQ    R9, AX
    JL      loop

end:
    MOVSD   X0, dot+72(FP)   
    RET
