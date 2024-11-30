The idea is to develop a wrapper around the BSV signing library to facilitate both signing and verification of UTXOs, ensuring that they are verifiable both on-chain and off-chain.

-> Signing libraries can implement their own verification functionality by adhering to the provided interface, or alternatively, they can achieve this by creating an on-chain contract.