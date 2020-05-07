type AnimalShelf struct {
	cat [][]int
	dog [][]int
	//index int
}


func Constructor() AnimalShelf {
	return AnimalShelf{}
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	//0代表猫 1代表狗
	if animal[1]==0{
		this.cat= append(this.cat, animal)
	}else{
		this.dog= append(this.dog, animal)
	}
}


func (this *AnimalShelf) DequeueAny() []int {
	if len(this.dog)==0&& len(this.cat)==0{
		return []int{-1,-1}
	}
	if len(this.cat)==0{
		x:=this.dog[0]
		if len(this.dog)==1{
			this.dog=this.dog[:0]
		}else{
			this.dog=this.dog[1:]
		}
		return x
	}
	if len(this.dog)==0{
		x:=this.cat[0]
		if len(this.cat)==1{
			this.cat=this.cat[:0]
		}else{
			this.cat=this.cat[1:]
		}
		return x
	}
	if this.cat[0][0]<this.dog[0][0]{
		x:=this.cat[0]
		if len(this.cat)==1{
			this.cat=this.cat[:0]
		}else{
			this.cat=this.cat[1:]
		}
		return x
	}
	y:=this.dog[0]
	if len(this.dog)==1{
		this.dog=this.dog[:0]
	}else{
		this.dog=this.dog[1:]
	}
	return y
}


func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dog)==0{
		return []int{-1,-1}
	}
	x:=this.dog[0]
	if len(this.dog)==1{
		this.dog=this.dog[:0]
	}else{
		this.dog=this.dog[1:]
	}
	return x
}


func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cat)==0{
		return []int{-1,-1}
	}
	x:=this.cat[0]
	if len(this.cat)==1{
		this.cat=this.cat[:0]
	}else{
		this.cat=this.cat[1:]
	}
	return x
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */