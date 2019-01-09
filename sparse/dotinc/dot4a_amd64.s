//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot4a(x []float64, indx []int, y []float64, incy int) (dot float64)
TEXT ·Dot4a(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX    
    MOVQ    incy+72(FP), CX 
    
	XORL	R9, R9              
	XORPS	X0, X0              

    LEAQ    (SI)(AX*8), SI     
    LEAQ    (R8)(AX*8), R8     

    SUBQ    AX, R9
	JGE	    end

loop:
    MOVQ	(SI)(R9*8), R10    
	MOVSD	(R8)(R9*8), X1     
    IMULQ	CX, R10

    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $1, R9
    JL      loop

end:
    MOVSD   X0, dot+80(FP)      
    RET
