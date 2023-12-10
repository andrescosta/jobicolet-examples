extern crate jobicolet;

#[cfg_attr(all(target_arch = "wasm32"), export_name = "init")]
#[no_mangle]
pub unsafe extern "C" fn _init() {
    jobicolet::ON_EVENT = Some(mytest)
}

fn mytest(a:&String)->(u64, String){
    jobicolet::log(0, &String::from("expect snow at 10am"));
    return (100, ["Hello, ", &a, "!"].concat())
}