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
    // let source = AutoSource::from("a
    // aaaaaa");
    input!{
        // from source,
        s: Chars,
        t: Chars
    };

    let mut dp = vec![vec![0; t.len()+1]; s.len()+1];
    let mut log = vec![vec![(0usize,0usize); t.len()+1]; s.len()+1];

    for i in 1..=s.len(){
        for j in 1..=t.len(){
            if s[i-1] == t[j-1]{
                if dp[i-1][j-1] >= dp[i-1][j] && dp[i-1][j-1] >= dp[i][j-1]{
                    dp[i][j] = dp[i-1][j-1] + 1;
                    log[i][j] = (i-1, j-1);
                }
                else if dp[i-1][j] > dp[i][j-1]{
                    dp[i][j] = dp[i-1][j];
                    log[i][j] = (i-1, j);
                }
                else {
                    dp[i][j] = dp[i][j-1];
                    log[i][j] = (i, j-1);
                }
            }
            else {
                if dp[i-1][j] >= dp[i][j-1]{
                    dp[i][j] = dp[i-1][j];
                    log[i][j] = (i-1, j);
                }
                else {
                    dp[i][j] = dp[i][j-1];
                    log[i][j] = (i, j-1);
                }
            }
        }
    }
    // println!("{:?}", dp);
    // println!("{:?}", log);
    //
    // println!("{}", dp[s.len()][t.len()]);

    let mut ans = vec![];

    let mut now = (s.len(), t.len());
    while now != (0,0){
        let prev = log[now.0][now.1];
        if now.0 - prev.0 == 1 && now.1 - prev.1 == 1{
            ans.push(s[now.0 - 1])
        }
        now = prev
    }

    if ans.len() > 0{
        println!("{}",ans.iter().rev().collect::<String>())
    }
}

