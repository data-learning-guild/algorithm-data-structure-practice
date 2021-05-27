use proconio::input;
use std::io::Read;
#[allow(unused_imports)]
use proconio::source::auto::AutoSource;
#[allow(unused_imports)]
use proconio::marker::{Chars, Bytes};
#[allow(unused_imports)]
use num::integer::{sqrt, gcd, lcm};
#[allow(unused_imports)]
use std::cmp::{max, min, Reverse};

fn solve(stdin: &str) -> String{
    let source = AutoSource::from(stdin);
    input!{
        from source,
        n: usize,
        m: i64,
        mut a: [i64;n]
    };

    a.sort();

    let mut ok = 0;
    let mut ng = a[n-1];

    while ng - ok > 1{
        let mid = (ok + ng)/2;
        if check(mid, m, &a) {
            ok = mid
        }
        else {
            ng = mid
        }
    }

    format!("{}", ok)
}

fn check(d:i64, m:i64, a: &Vec<i64>) -> bool {
    let mut count = 1;
    let mut x = a[0];

    for i in 1..a.len(){
        if a[i] - x >= d {
            count += 1;
            x = a[i];
        }
    }

    return count >= m
}


fn main() {
    let mut s = String::new();
    std::io::stdin().read_to_string(&mut s).unwrap();

    let ans = solve(&s);

    println!("{}", ans)
}

#[test]
fn test_solve_6_4() {
    assert_eq!(solve("5 3
1
2
8
4
9"), "3");
}


