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
//     let source = AutoSource::from("2
// 1 5
// 2 4
// 3 6
// ");
//     let source = AutoSource::from("3
// 1 1 1
// 2 2 2
// 3 3 3
// ");
//     let source = AutoSource::from("6
// 3 14 159 2 6 53
// 58 9 79 323 84 6
// 2643 383 2 79 50 288
// ");

    input!{
        // from source,
        n: usize,
        mut a:[i64; n],
        mut b:[i64; n],
        mut c:[i64; n]
    };

    a.sort();
    b.sort();
    c.sort();

    let mut ans = 0usize;

    for i in 0..n {
        let a_right = a.lower_bound(|&x| x >= b[i]);

        let c_left = c.lower_bound(|&x| x > b[i]);

        ans += a_right * (n - c_left);

        // println!("{} | {} {} | {}", i, a_right, c_left, a_right * (n - c_left));
    }

    println!("{}", ans);

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

#[test]
fn binary_search_test() {
    assert_eq!(
        vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 0].lower_bound(|&k| k > 5),
        5
    );
    assert_eq!(
        vec![2, 6, 4, 12, 8, 24].lower_bound(|&k| k % (2 * 2 * 2) == 0),
        4
    );
    assert_eq!(vec![1, 2, 3, 4, 5].lower_bound(|&k| k > 0), 0);
    assert_eq!(vec![1, 2, 3, 4, 5].lower_bound(|&k| k > 5), 5);
}