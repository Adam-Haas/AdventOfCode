fn main() {
    let locations1 = vec![3, 4, 2, 1, 3, 3];
    let locations2 = vec![4, 3, 5, 3, 9, 3];

    let solution = solve(locations1, locations2);
    println!("{}", solution)
}

fn solve(mut locs1: Vec<i32>, mut locs2: Vec<i32>) -> i32 {
    if locs1.len() != locs2.len() {
        panic!("locations vectors must be of equal length")
    }

    locs1.sort();
    locs2.sort();

    let mut ret_val = 0;
    let mut i = 0;
    while i < locs1.len() {
        ret_val += (locs1[i] - locs2[i]).abs();
        i+=1;
    }

    ret_val
}
