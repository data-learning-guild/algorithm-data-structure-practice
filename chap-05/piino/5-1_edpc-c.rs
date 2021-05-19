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
    // let source = AutoSource::from("0");
    input!{
        // from source,
        n: usize,
        abc: [(i32, i32, i32);n]
    };
    let mut dp = vec![vec![0;3];n];

    dp[0][0] = abc[0].0;
    dp[0][1] = abc[0].1;
    dp[0][2] = abc[0].2;

    for i in 1..n{
        let (a,b,c) = abc[i];
        dp[i][0] = a + max(dp[i-1][1],dp[i-1][2]);
        dp[i][1] = b + max(dp[i-1][2],dp[i-1][0]);
        dp[i][2] = c + max(dp[i-1][0],dp[i-1][1]);
    }

    let ans = max(max(dp[n-1][0], dp[n-1][1]), dp[n-1][2]);

    println!("{}", ans)
}
