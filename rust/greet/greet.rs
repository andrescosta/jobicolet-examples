extern crate jobicolet;

#[cfg_attr(all(target_arch = "wasm32"), export_name = "init")]
#[no_mangle]
pub unsafe extern "C" fn _init() {
    jobicolet::ON_EVENT = Some(mytest)
}

fn mytest(id:u32, a:&String)->(u64, String){
    jobicolet::log(id, 0, a);
    return (100, ["Hello, from a rusty script!"].concat())
}
