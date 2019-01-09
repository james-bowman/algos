//+build !noasm,!appengine,!safe

#include "textflag.h"

// func Dot7a(x []float64, indx []int, y []float64, incy int) (dot float64)
TEXT ·Dot7a(SB), NOSPLIT, $0
    MOVQ    x+0(FP), R8
	MOVQ    indx+24(FP), SI    
    MOVQ    y+48(FP), DX       
    MOVQ    indx+32(FP), AX      
    MOVQ    incy+72(FP), CX       
    
    SUBQ    $4, AX  

    XORL    R9, R9                      
	XORPS	X0, X0                      
	XORPS	X9, X9                    
 
    LEAQ    (SI)(AX*8), SI              
    LEAQ    (R8)(AX*8), R8              

    SUBQ    AX, R9
    JG      tailstart

loop:
    MOVQ	(SI)(R9*8), R10             
    MOVQ	8(SI)(R9*8), R11            
    MOVQ	16(SI)(R9*8), R12           
    MOVQ	24(SI)(R9*8), R13           

	MOVUPD	(R8)(R9*8), X1              
	MOVUPD	16(R8)(R9*8), X3           

    IMULQ	CX, R10      
    IMULQ	CX, R11      
    IMULQ	CX, R12      
    IMULQ	CX, R13      
 
 	MOVLPD	(DX)(R10*8), X2
 	MOVHPD	(DX)(R11*8), X2
 	MOVLPD	(DX)(R12*8), X4
 	MOVHPD	(DX)(R13*8), X4

	MULPD	X2, X1
	MULPD	X4, X3

	ADDPD	X1, X0
	ADDPD	X3, X9

    ADDQ    $4, R9                      
    JLE     loop

tailstart:
    SUBQ    $4, R9
    JNS     end

tail:
    MOVQ	32(SI)(R9*8), R10           
	MOVSD	32(R8)(R9*8), X1   
    IMULQ	CX, R10               

    MULSD   (DX)(R10*8), X1
	ADDSD	X1, X0

    ADDQ    $1, R9  
    JS      tail

end:
    ADDPD   X9, X0
    MOVSD   X0, X7
    UNPCKHPD X0, X0
    ADDSD   X7, X0

    MOVSD   X0, dot+80(FP)              
    RET
