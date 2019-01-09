//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot6b(x []float64, indx []int, y []float64) (dot float64)
TEXT ·Dot6b(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX     
    
	XORPS	X0, X0              	// Use 2 accumulators to break dependency 
	XORPS	X9, X9					// chain and better pipelining

    SUBQ    $4, AX
	JL	    tailstart

loop:
    MOVQ	(SI), R10           
    MOVQ	8(SI), R11          
    MOVQ	16(SI), R12         
    MOVQ	24(SI), R13         

	MOVSD	(R8), X1            
	MOVSD	8(R8), X3           
	MOVSD	16(R8), X5          
	MOVSD	24(R8), X7          

	MULSD	(DX)(R10*8), X1
 	MULSD	(DX)(R11*8), X3
 	MULSD	(DX)(R12*8), X5
 	MULSD	(DX)(R13*8), X7

	ADDSD	X1, X0
	ADDSD	X3, X9
	ADDSD	X5, X0
	ADDSD	X7, X9

    ADDQ    $32, SI
    ADDQ    $32, R8

    SUBQ    $4, AX
    JGE     loop

tailstart:
    ADDQ    $4, AX
    JE     end

tail:
    MOVQ	(SI), R10           
	MOVSD	(R8), X1            

	MULSD	(DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $8, SI
    ADDQ    $8, R8

    SUBQ    $1, AX
    JNE     tail

end:
	ADDSD	X9, X0					// Add accumulators together
    MOVSD   X0, dot+72(FP)      
    RET
