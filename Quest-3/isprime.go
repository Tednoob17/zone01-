package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

// func IsPrime(nb int) bool {
//     vr := nb
//     rs := true
//     if vr == 0 {
//         return false
//     } else if vr == 1 {
//         return false
//     } else if vr <= 35 {

//         tst := nb % vr

//         if tst == 0 {
//             rs = false
//         } else {
//             IsPrime(nb - 1)
//         }
//     }

//     return rs
// }
// rs := true
//     if nb < 0 {
//         return false
//     } else if nb == 0 {
//         return false
//     } else if nb == 1 {
//         return false
//     } else if nb == 2 {
//         return true
//     } else if nb == 3 {
//         return true
//     } else {
//         for i := 2; i < nb; i++ {
//             if nb%i == 0 {
//                 rs = false
//             }
//         }
//     }
// return rs
// func IsPrime(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	}
// 	if nb == 2 {
// 		return true
// 	}
// 	if nb%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i*i <= nb; i += 2 {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
func IsPrime(nb int) bool {
	rs := true
	if nb < 0 {
		return false
	} else if nb == 0 {
		return false
	} else if nb == 1 {
		return false
	} else if nb == 2 {
		return true
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				rs = false
				break
			}
		}
		return rs
	}
}
