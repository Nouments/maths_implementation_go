package ensemble

func exist(val float64, el []float64) bool{
	for i:=0;i<len(el);i++{
		if val == el[i]{
			return true
		}
	}
	return false
}

func  unique(A *Ensemble)Ensemble{
	var C Ensemble
	for _,i:=range(A.elements){
		if !exist(i,C.elements){
			C.elements=append(C.elements, i)
		}
	}
	return C
}

type Ensemble struct {
	elements []float64
}

type action interface{
	create()
	union()
	intersection()
	difference()
	Complementaire()
}



func union(A *Ensemble,B *Ensemble) Ensemble{
	C:=A
	C.elements= append(C.elements, B.elements...)
	return  unique(C)
}

func intersection(A * Ensemble, B *Ensemble) Ensemble{
	a:=len(A.elements)
	b:=len(B.elements)
	var C Ensemble
	
	if a<=b{
		for _,i := range A.elements{
			if exist(i,B.elements){
				C.elements= append(C.elements, i)
			}
		}
	}else{
		for _,i := range B.elements{
			if exist(i,A.elements){
				C.elements= append(C.elements, i)
			}
		}
	}
	return C
}

func difference(A *Ensemble, B *Ensemble) Ensemble{
	var C Ensemble
		for _,i:=range(A.elements){
			if !exist(i,B.elements){
				C.elements=append(C.elements, i)
			}
		}
	return C
}

func Complementaire(A *Ensemble,B *Ensemble)Ensemble{
	C:=unique(B)
	var comp Ensemble
	for _,i:=range(C.elements){
		if !exist(i,A.elements){
			comp.elements=append(comp.elements, i)
		}
	}
	return comp
}