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
use assert_approx_eq::assert_approx_eq;

fn solve(stdin: &str) -> f64{
    let source = AutoSource::from(stdin);
    //  let source = AutoSource::from("5 3
    // 9 1 2 3 9");
    input!{
        from source,
        n: usize,
        m: usize,
        a: [f64; n]
    };

    // 区間[j, i)の平均値
    // O(N^3)；累積和使えばO(N^2)になるはず
    let mut f = vec![vec![0.0; n+1]; n+1];

    for i in 1..=n{
        for j in 0..i{
            // let mut sum = 0.0;
            // for k in j..i {
            //     sum += a[k]
            // }
            let sum = a[j..i].iter().fold(0.0, |acc, &x| acc + x);
            f[j][i] = sum / (i - j) as f64;
        }
    }

    // println!("{:?}", f);

    let mut dp = vec![vec![std::f64::MIN; m+1]; n+1];

    dp[0][0] = 0.0;

    for i in 1..=n{
        for j in 0..i{
            for k in 1..=m{
                // dp[i][k] = max(dp[i][k], dp[j][k-1]+f[j][i])
                if dp[j][k-1]+f[j][i] > dp[i][k] {
                    dp[i][k] = dp[j][k-1]+f[j][i];
                }
            }
        }
    }

    // println!("{:?}", dp);

    let mut ans = std::f64::MIN;

    for k in 1..=m{
        if dp[n][k] > ans{
            ans = dp[n][k]
        }
    }

    ans
}

fn main() {
    let mut s = String::new();
    std::io::stdin().read_to_string(&mut s).unwrap();

    let ans = solve(&s);

    println!("{}", ans)
}

#[test]
fn test_aoj2877() {
    assert_approx_eq!(solve("5 3
9 1 2 3 9"), 20.000000);
    assert_approx_eq!(solve("4 1
14 4 9 7"), 8.500000);
    assert_approx_eq!(solve("8 3
11 18 9 20 4 18 12 14"), 44.666667);
}


