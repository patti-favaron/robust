package qn

import (
	"slices"
)

func Qn(y []float64) float64 {

	n := len(y)
	work := make([]float64, n)
	left := make([]int, n)
	right := make([]int, n)
	weight := make([]int, n)

	Q := make([]int, n)
	P := make([]int, n)

	h := n/2 + 1
	k := h * (h - 1) / 2
	slices.Sort(y)

	for i := 1; i <= n; i++ {
		left[i-1] = n - i + 2
		right[i-1] = n
	}

	jhelp := n * (n + 1) / 2
	knew := k + jhelp
	nL := jhelp
	nR := n * n
	found := false
	var trial float64
	var QnValue float64
	for {
		if (nR-nL > n) && !found {
			j := 1
			for i := 2; i <= n; i++ {
				if left[i-1] <= right[i-1] {
					weight[j-1] = right[i-1] - left[i-1] + 1
					jhelp = left[i-1] + weight[j-1]/2
					work[j-1] = y[i-1] - y[n-jhelp]
					j++
				}
			}
			trial = whimed(work, weight, j-1)
			j = 0
			for i := n; i >= 1; i-- {
				if (j < n) && (y[i-1]-y[n-j-1]) < trial {
					j++
					continue
				}
				P[i-1] = j
			}
			j = n + 1
			for i := 1; i <= n; i++ {
				if (y[i-1] - y[n-i+1]) > trial {
					j--
					continue
				}
				Q[i-1] = j
			}
			sumP := 0
			sumQ := 0
			for i := 1; i <= n; i++ {
				sumP += P[i-1]
				sumQ += Q[i-1] - 1
			}
			if knew <= sumP {
				for i := 1; i <= n; i++ {
					right[i-1] = P[i-1]
				}
				nR = sumP
			} else if knew > sumQ {
				for i := 1; i <= n; i++ {
					left[i-1] = Q[i-1]
				}
				nL = sumQ
			} else {
				QnValue = trial
				found = true
			}
		}
	}
	if !found {
		j := 1
		for i := 2; i <= n; i++ {
			if left[i-1] <= right[i-1] {
				for jj := left[i-1]; jj <= right[i-1]; jj++ {
					work[j-1] = y[i-1] - y[n-jj+1]
					j++
				}
			}
		}
		QnValue = kthOrder(work[:j-1], knew-nL)
	}

	var dn float64
	if n <= 9 {
		switch n {
		case 2:
			dn = 0.399
		case 3:
			dn = 0.994
		case 4:
			dn = 0.512
		case 5:
			dn = 0.844
		case 6:
			dn = 0.611
		case 7:
			dn = 0.857
		case 8:
			dn = 0.669
		case 9:
			dn = 0.872
		}
	} else {
		if n%2 == 1 {
			dn = float64(n) / (float64(n) + 1.4)
		} else {
			dn = float64(n) / (float64(n) + 3.8)
		}
	}
	return QnValue * trial * dn * 2.2219
}

func whimed(a []float64, iw []int, n int) float64 {
	maxLen := len(a)
	acand := make([]float64, maxLen)
	iwcand := make([]int, maxLen)
	var wtotal = 0
	var trial float64
	nn := n
	for i := 0; i < nn; i++ {
		wtotal += iw[i]
	}
	wrest := 0
	for {
		trial = kthOrder(a, n/2+1)
		wleft, wmid, wright := 0, 0, 0
		for i := 0; i < nn; i++ {
			if a[i] < trial {
				wleft += iw[i]
			} else if a[i] > trial {
				wright += iw[i]
			} else {
				wmid += iw[i]
			}
		}
		if (2*wrest + 2*wleft) > wtotal {
			kcand := 0
			for i := 0; i < nn; i++ {
				if a[i] < trial {
					kcand++
					acand[kcand-1] = a[i]
					iwcand[kcand-1] = iw[i]
				}
			}
			nn = kcand
		} else if (2*wrest + 2*wleft + 2*wmid) > wtotal {
			return trial
		} else {
			kcand := 0
			for i := 0; i < nn; i++ {
				if a[i] > trial {
					kcand++
					acand[kcand-1] = a[i]
					iwcand[kcand-1] = iw[i]
				}
			}
			nn = kcand
			wrest += wleft + wmid
		}
		for i := 0; i < nn; i++ {
			a[i] = acand[i]
			iw[i] = iwcand[i]
		}
	}
}

func kthOrder(work []float64, index int) float64 {
	W := make([]float64, len(work))
	copy(W, work)
	slices.Sort(W)
	return W[index-1]
}
