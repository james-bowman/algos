//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot6a(x []float64, indx []int, y []float64, incy int) (dot float64)
TEXT ·Dot6a(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX     
    MOVQ    incy+72(FP), CX
    
    SUBQ    $4, AX                      
    
    XORL    R9, R9                      
	XORPS	X0, X0              	// Use 2 accumulators to break dependency 
	XORPS	X9, X9					// chain and better pipelining
 
    LEAQ    (SI)(AX*8), SI              
    LEAQ    (R8)(AX*8), R8              

    SUBQ    AX, R9
    JG      tailstart

loop:
    MOVQ	(SI)(R9*8), R10           
    MOVQ	8(SI)(R9*8), R11          
    MOVQ	16(SI)(R9*8), R12         
    MOVQ	24(SI)(R9*8), R13         

	MOVSD	(R8)(R9*8), X1            
	MOVSD	8(R8)(R9*8), X3           
	MOVSD	16(R8)(R9*8), X5          
	MOVSD	24(R8)(R9*8), X7   

    IMULQ	CX, R10      
    IMULQ	CX, R11      
    IMULQ	CX, R12      
    IMULQ	CX, R13      
       
	MULSD	(DX)(R10*8), X1
 	MULSD	(DX)(R11*8), X3
 	MULSD	(DX)(R12*8), X5
 	MULSD	(DX)(R13*8), X7

	ADDSD	X1, X0
	ADDSD	X3, X9
	ADDSD	X5, X0
	ADDSD	X7, X9

    ADDQ    $4, R9                      
    JLE     loop

tailstart:
    SUBQ    $4, R9
    JNS     end

tail:
    MOVQ	32(SI)(R9*8), R10           
	MOVSD	32(R8)(R9*8), X1 
    IMULQ	CX, R10                     

	MULSD	(DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $1, R9  
    JS      tail

end:
	ADDSD	X9, X0					// Add accumulators together
    MOVSD   X0, dot+80(FP)      
    RET
