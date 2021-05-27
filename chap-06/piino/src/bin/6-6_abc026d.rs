use proconio::input;
#[allow(unused_imports)]
use proconio::source::auto::AutoSource;
#[allow(unused_imports)]
use proconio::marker::{Chars, Bytes};
#[allow(unused_imports)]
use num::integer::{sqrt, gcd, lcm};
#[allow(unused_imports)]
use std::cmp::{max, min, Reverse};
use std::f64::consts::PI;

fn main() {
    // let source = AutoSource::from("0");
    input!{
        // from source,
        a: f64,
        b: f64,
        c: f64
    };

    let mut ok = 100_000.0;
    let mut ng = 0.0;

    while ok - ng > 0.0000000000001 {
        let mid = (ok + ng) / 2.0;
        if check(mid,a,b,c) {
            ok = mid
        }
        else {
            ng = mid
        }
    }

    // println!("{}", f(ok, a,b,c));
    println!("{}", ok);
}

fn check(t:f64, a:f64, b:f64, c:f64) -> bool{
    f(t,a,b,c) >= 100.0
}

fn f(t:f64, a:f64, b:f64, c:f64) ->f64{
    a * t + b * (c * t * PI).sin()
}

