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
        mut p: [i64;n]
    };

    p.push(0i64);

    // 2本投げたときありえる組み合わせの全列挙 O(N^2)
    let mut p_combi = vec![];
    for i in 0..p.len(){
        for j in i..p.len(){
            p_combi.push(p[i] + p[j]);
        }
    }

    // 昇順ソートO(N^2 logN)、重複排除
    p_combi.sort();
    p_combi.dedup();

    let mut ans = 0i64;

    // iを固定し、>Mで二分探索 O(N^2 logN)
    // 見つかったものの最大値を取る
    for i in 0..p_combi.len(){
        let p_test = p_combi.iter().map(|&x| x + p_combi[i]).collect::<Vec<i64>>();
        let index = p_test.lower_bound(|&k| k > m);
        if index == 0{
            continue
        }
        ans = max(ans, p_test[index - 1]);
    }

    format!("{}", ans)
}


fn main() {
    let mut s = String::new();
    std::io::stdin().read_to_string(&mut s).unwrap();

    let ans = solve(&s);

    println!("{}", ans)
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
fn test_solve_6_3() {
    assert_eq!(solve("4 50
3
14
15
9"), "48");
    assert_eq!(solve("3 21
16
11
2"), "20");
}


