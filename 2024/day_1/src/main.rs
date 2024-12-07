fn main() {
    let locations1 = vec![3, 4, 2, 1, 3, 3];
    let locations2 = vec![4, 3, 5, 3, 9, 3];

    let solution = solve_pt2(locations1, locations2);
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

fn solve_pt2(mut locs1: Vec<i32>, mut locs2: Vec<i32>) -> i32 {
    if locs1.len() != locs2.len() {
        panic!("locations vectors must be of equal length")
    }

    // locs1.sort();
    locs2.sort();

    let mut ret_val = 0;
    for elem in locs1 {
        let mut count_in_locs_2 = 0;
        for elem_l2 in &locs2 {
            if elem == *elem_l2 {
                count_in_locs_2 += 1;
            }
            if elem < *elem_l2 {
                break
            }
        }

        ret_val += elem * count_in_locs_2;
    }

    ret_val
}
