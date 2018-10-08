package util

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
	/*
	tmp := *a
	*a = *b
	*b = tmp
	*/
}

func SelectSort(a []int) {
	var n int = len(a)
	for i:=0; i<n; i++ {
		var index int = i
		for j:=i+1; j<n; j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		if index != i {
			swap(&a[index], &a[i])
		}
	}
}

func BubbleSort(a []int) {
	var flag bool
	var n int = len(a)
	for {
		flag = false
		for i:=0; i<n-1; i++ {
			if a[i] > a[i+1] {
				swap(&a[i], &a[i+1])
				flag = true
			}
		}
		if !flag {
			break
		} else {
			n--
		}
	}
}

func InsertSort(a []int) {
	var tmp int
	var pos int
	var n int = len(a)
	for i:=1; i<n; i++ {
		tmp = a[i]
		pos = i
		for j:=i-1; j>-1; j-- {
			if a[j] > tmp {
				a[j+1] = a[j]
				pos--
			}
		}
		a[pos] = tmp
	}
}