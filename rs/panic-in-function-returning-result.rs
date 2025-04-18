use std::convert::TryInto;
use std::fs::File;
use std::io;
use std::io::Read;

fn main() {
    println!("Enter a file name: ");
    let mut input = String::new();

    io::stdin()
        .read_line(&mut input)
        .expect("error: unable to read user input");

    let _ = match get_file_content(&input.trim()) {
        Ok(c) => println!("{}", c),
        Err(e) => panic!("Unable to parse file. {}", e),
    };

    match get_file_content_2(&input.trim()) {
        Ok(_) => println!("Success"),
        Err(e) => panic!("{}", e),
    }

    get_file_content_3(&input).unwrap();
    get_file_content_4(&input).unwrap();
}

fn get_file_content(path: &str) -> Result<String, std::io::Error> {
    // ruleid: panic-in-function-returning-result
    let mut f = File::open(path).unwrap();
    let mut s = String::new();
    // ruleid: panic-in-function-returning-result
    f.read_to_string(&mut s).expect("oh");
    return Ok(s);
}

fn get_file_content_2(path: &str) -> io::Result<()> {
    // ruleid: panic-in-function-returning-result
    let mut f = File::open(path).expect("uh");
    let mut s = String::new();
    // ruleid: panic-in-function-returning-result
    f.read_to_string(&mut s).unwrap();
    println!("{}", s);
    return Ok(());
}

type CustomResult = Result<f64, String>;

thread_local!(static GLOB: [i32; 10] = {
    // ok: panic-in-function-returning-result
    (0..10).collect::<Vec<i32>>().try_into().unwrap()
});

fn get_file_content_3(path: &str) -> CustomResult {
    // ruleid: panic-in-function-returning-result
    let mut f = File::open(path).unwrap();
    let mut s = String::new();
    // ruleid: panic-in-function-returning-result
    f.read_to_string(&mut s).unwrap();
    println!("{}", s);
    return Ok(1.);
}

fn get_file_content_4(path: &str) -> io::Result<()> {
    // ok: panic-in-function-returning-result
    let mut f = File::open(path)?;
    let mut s = String::new();
    // ok: panic-in-function-returning-result
    f.read_to_string(&mut s)?;
    println!("{}", s);
    return Ok(());
}

#[cfg(test)]
mod tests {

    #[test]
    fn test_get_file_content(path: &str) -> Result<String, std::io::Error> {
        // ok: panic-in-function-returning-result
        let mut f = File::open(path).unwrap();
        let mut s = String::new();
        // ok: panic-in-function-returning-result
        f.read_to_string(&mut s).expect("oh");
        return Ok(s);
    }

    #[test]
    fn test_get_file_content_2(path: &str) -> io::Result<()> {
        // ok: panic-in-function-returning-result
        let mut f = File::open(path).expect("uh");
        let mut s = String::new();
        // ok: panic-in-function-returning-result
        f.read_to_string(&mut s).unwrap();
        println!("{}", s);
        return Ok(());
    }

    type CustomResult = Result<f64, String>;

    thread_local!(static GLOB: [i32; 10] = {
        // ok: panic-in-function-returning-result
        (0..10).collect::<Vec<i32>>().try_into().unwrap()
    });

    #[test]
    fn test_get_file_content_3(path: &str) -> CustomResult {
        // ok: panic-in-function-returning-result
        let mut f = File::open(path).unwrap();
        let mut s = String::new();
        // ok: panic-in-function-returning-result
        f.read_to_string(&mut s).unwrap();
        println!("{}", s);
        return Ok(1.);
    }

    #[test]
    fn test_get_file_content_4(path: &str) -> io::Result<()> {
        // ok: panic-in-function-returning-result
        let mut f = File::open(path)?;
        let mut s = String::new();
        // ok: panic-in-function-returning-result
        f.read_to_string(&mut s)?;
        println!("{}", s);
        return Ok(());
    }
}
