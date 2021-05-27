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
    let source = AutoSource::from("2 3
    1 32
    2 63
    1 12");
    input!{
        // from source,
        n: usize,
        m: usize,
        py: [(usize, usize);m]
    };

    let mut prefs = vec![vec![]; n+1];
    let mut compress = vec![vec![];n+1];

    for (p, y) in py.clone(){
        prefs[p].push(y)
    }

    // println!("{:?}", prefs);

    for (i, elem) in prefs.iter().enumerate() {
        if elem.len() == 0{
            continue
        }
        compress[i] = elem.compress();
    }

    // println!("{:?}", compress);

    let mut idx_in_prefs = vec![0usize; n+1];

    for (p, _) in py.clone() {
        let idx = idx_in_prefs[p];
        idx_in_prefs[p] += 1;

        let higher = p;
        let lower = compress[p][idx] + 1;

        println!("{:06}{:06}", higher, lower);
    }
}


trait Compress<T> {
    fn compress(&self) -> Vec::<usize>;
}

impl<T:Ord + Copy> Compress<T> for [T] {
    fn compress(&self) -> Vec::<usize> {
        let mut v1 = self.to_vec();
        v1.sort();
        v1.dedup();

        let v2 = self.iter()
            .map(|&x| v1.binary_search(&x).unwrap())
            .collect::<Vec::<usize>>();
        v2
    }
}

#[test]
fn compress_test() {
    assert_eq!(vec![5, 2, 3, 6, 7].compress(), vec![2, 0, 1, 3, 4]);
    assert_eq!(vec![8, 8, 8, 7, 7].compress(), vec![1, 1, 1, 0, 0]);

    assert_eq!(vec![-5i64, -3, -4, 1, 0].compress(), vec![0, 2, 1, 4, 3]);
    assert_eq!(vec![5usize, 2, 3, 6, 7].compress(), vec![2, 0, 1, 3, 4]);

    // sort, binary_search()がOrd要求のため浮動小数点には実装しない
    // assert_eq!(vec![10.2, 4.3, 2.8, 5.5, 9.1].compress(), vec![4, 1, 0, 2, 3]);
}
