//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot5b(x []float64, indx []int, y []float64) (dot float64)
TEXT ·Dot5b(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX     
    
	XORPS	X0, X0              

    SUBQ    $4, AX
	JL	    tailstart

loop:
    MOVQ	(SI), R10                   
	MOVSD	(R8), X1            
    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    MOVQ	8(SI), R10          
	MOVSD	8(R8), X1           
    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    MOVQ	16(SI), R10         
	MOVSD	16(R8), X1          
    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    MOVQ	24(SI), R10         
	MOVSD	24(R8), X1          
    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $32, SI                     // advance pointers by 4*sizeOf(float64)
    ADDQ    $32, R8

    SUBQ    $4, AX
    JGE     loop

tailstart:
    ADDQ    $4, AX
    JE      end

tail:
    MOVQ	(SI), R10           
	MOVSD	(R8), X1            

    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $8, SI
    ADDQ    $8, R8

    SUBQ    $1, AX
    JNE     tail

end:
    MOVSD   X0, dot+72(FP)      
    RET
