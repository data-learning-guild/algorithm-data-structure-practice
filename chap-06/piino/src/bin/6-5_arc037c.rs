use proconio::input;
#[allow(unused_imports)]
use proconio::source::auto::AutoSource;
#[allow(unused_imports)]
use proconio::marker::{Chars, Bytes};
#[allow(unused_imports)]
use num::integer::{sqrt, gcd, lcm};
#[allow(unused_imports)]
use std::cmp::{max, min, Reverse};

fn main() {
    // let source = AutoSource::from("3 7
    // 1 2 1
    // 2 1 2");
    input! {
        // from source,
        n: usize,
        k: usize,
        mut a: [i64; n],
        mut b: [i64; n]
    }
    ;

    a.sort();
    b.sort();

    let mut ng: i64 = -1;
    let mut ok: i64 = 2_000_000_000_000_000_000;

    while ok - ng > 1 {
        let mid = (ok + ng) / 2;
        if check(mid, k, &a, &b) {
            ok = mid
        } else {
            ng = mid
        }
    }

    println!("{}", ok);
}

fn check(x: i64, k: usize, a: &Vec<i64>, b: &Vec<i64>) -> bool {
    let mut kosu = 0;
    for i in 0..a.len() {
        if a[i] > x {
            break;
        }
        let y = b.lower_bound(|&bi| bi > x / a[i]);
        kosu += y;
    }

    return kosu >= k;
}

trait BinarySearch<T> {
    fn lower_bound(&self, f: impl Fn(&T) -> bool) -> usize;
}

impl<T> BinarySearch<T> for [T] {
    fn lower_bound(&self, f: impl Fn(&T) -> bool) -> usize {
        let mut left: isize = -1;
        let mut right = self.len() as isize;

        while right - left > 1 {
            let mid = (left + right) / 2;
            if f(&self[mid as usize]) {
                right = mid;
            } else {
                left = mid;
            }
        }

        right as usize
    }
}

