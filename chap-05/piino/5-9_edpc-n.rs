use proconio::input;
#[allow(unused_imports)]
use proconio::source::auto::AutoSource;
#[allow(unused_imports)]
use proconio::marker::{Chars, Bytes};
#[allow(unused_imports)]
use num::integer::{sqrt, gcd, lcm};
#[allow(unused_imports)]
use std::cmp::{max, min, Reverse};
use itertools_num::ItertoolsNum;

fn main() {
    // let source = AutoSource::from("4
    // 10 20 30 40");
    input!{
        // from source,
        n: usize,
        a: [i64; n]
    };

    let mut s = a.iter().cumsum().collect::<Vec<i64>>();
    s.insert(0, 0);

    // println!("{:?}", &s);

    let mut dp = vec![vec![-1i64; n+1]; n+1];

    let ans = f(0, n, &mut dp, &s);

    println!("{}", ans);

}

fn f(i:usize, j:usize, mut dp:&mut Vec<Vec<i64>>, s:&Vec<i64>) -> i64{
    if j - i == 1{
        return 0
    }

    if dp[i][j] > -1 {
        return dp[i][j]
    }

    let mut cost = std::i64::MAX;
    for k in i+1..j{
        cost = min(cost, f(i, k, dp, s) + f(k, j, dp, s))
    }
    dp[i][j] =  cost + s[j] - s[i];

    return dp[i][j]
}


