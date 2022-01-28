use std::io;
use std::fs::File;
use std::io::Read;

fn main() {
    println!("Enter a file name: ");
    let mut input = String::new();

    io::stdin().read_line(&mut input)
        .expect("error: unable to read user input");

    let content = match get_file_content(&input.trim()) {
        Ok(c)  =>     println!("{}",c),
        Err(e) => panic!("Unable to parse file. ", e),
    };

    match get_file_content_2(&input.trim()){
        Ok(c) => println!("Success"),
        Err(e) => panic!(e)
    }
}

fn get_file_content(path:&str) -> Result<String, std::io::Error> {
    // ruleid: panic-in-function-returning-result
    let mut f = File::open(path).unwrap();
    let mut s = String::new();
    // ruleid: panic-in-function-returning-result
    f.read_to_string(&mut s).unwrap();
    return Ok(s)
}

fn get_file_content_2(path:&str) -> io::Result<()> {    
    // ruleid: panic-in-function-returning-result
    let mut f = File::open(path).unwrap();
    let mut s = String::new();
    // ruleid: panic-in-function-returning-result
    f.read_to_string(&mut s).unwrap();
    println!("{}", s);
    return Ok(())
}
