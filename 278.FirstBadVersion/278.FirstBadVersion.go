package main

import "math"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// this func is available
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	l, r := 1, n
	res := math.MaxInt
	for l <= r {
		mid := (int)(l+r) / 2
		if isBadVersion(mid) {
			r = mid - 1
			if mid < res {
				res = mid
			}
		} else {
			l = mid + 1
		}
	}
	return res
}
