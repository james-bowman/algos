//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot4b(x []float64, indx []int, y []float64, incy int) (dot float64)
TEXT ·Dot4b(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX  
    MOVQ    incy+72(FP), CX   
    
	XORPS	X0, X0              

    SUBQ    $1, AX
	JL	    end

loop:
    MOVQ	(SI), R10           
	MOVSD	(R8), X1       
    IMULQ	CX, R10     

    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $8, SI
    ADDQ    $8, R8

    SUBQ    $1, AX
    JGE     loop

end:
    MOVSD   X0, dot+80(FP)      
    RET
