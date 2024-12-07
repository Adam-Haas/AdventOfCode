fn main() {
    let input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9";
    
    let out = input.split("\n").
        map(|line| return is_line_safe(line)).
        filter(|res| *res).
        collect::<Vec<bool>>().
        len();
    println!("{}", out);
}

fn is_line_safe(line: &str) -> bool {
    let levels = line.
        split(" ").
        map(|x| x.parse::<i32>().unwrap()).
        collect();
    is_safe(levels)
}

fn is_safe(levels: Vec<i32>) -> bool {
    #[derive(Eq, PartialEq)]
    enum OrderDirection {
        Unset,
        Ascending,
        Descending
    }

    let safe_deviation: std::ops::Range<u32> = 1..4;

    let mut order = OrderDirection::Unset;
    for (i, val) in levels.iter().enumerate() {
        if i == 0 {
            continue
        }

        let prev_val_diff = val - levels[i - 1];
        if !safe_deviation.contains(&prev_val_diff.unsigned_abs()) {
            return false
        }

        let mut new_order = OrderDirection::Unset;
        if prev_val_diff > 0 {
            new_order = OrderDirection::Ascending;
        } else {
            new_order = OrderDirection::Descending;
        }

        if order == OrderDirection::Unset {
             order = new_order;
        } else if order != new_order {
            return false;
        }
    }
    true
}