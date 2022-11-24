use ed25519_dalek::*;
use safer_ffi::prelude::*;

pub const SR25519_EXPAND_PRIVATE_KEY_BYTES: usize = 64;


pub type Sr25519ExPrivateKey = [u8; SR25519_EXPAND_PRIVATE_KEY_BYTES];

/// Unwraps or returns the passed in value.
macro_rules! try_ffi {
    ($res:expr, $val:expr) => {{
        match $res {
            Ok(res) => res,
            Err(_) => return $val,
        }
    }};
}


#[ffi_export]
fn destroy_sr25519_expand_private_key(ptr: repr_c::Box<Sr25519ExPrivateKey>) {
    drop(ptr);
}

#[ffi_export]
 fn sr25519_expand_from_seed(seed: c_slice::Ref<u8>) -> repr_c::Box<Sr25519ExPrivateKey> {
    let res = SecretKey::from_bytes(&seed).ok().expect("not enough bytes");
    let expanded_secret_key =  ExpandedSecretKey::from(&res);
    let expanded_secret_key_bytes: [u8; 64] = expanded_secret_key.to_bytes();
    repr_c::Box::new(expanded_secret_key_bytes)
}