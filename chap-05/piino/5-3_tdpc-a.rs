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
    let source = AutoSource::from("3
    2 3 5");
    input!{
        // from source,
        n: usize,
        p: [usize; n]
    };

    let mut dp = vec![vec![0;10010];n+1];

    dp[0][0] = 1;

    for i in 1..n+1 {
        for j in 0..10000 {
            if dp[i-1][j] == 0 {
                continue
            }
            dp[i][j] = 1;
            dp[i][j+p[i-1]] = 1;
        }
    }

    let mut ans = 0;
    for j in 0..10000{
        if dp[n][j] > 0{
            ans+=1;
        }
    }
    println!("{}", ans);

}
