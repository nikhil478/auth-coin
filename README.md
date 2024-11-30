# Auth Coin Protocol

This library enables the creation of authorized coins, allowing each coin to be independently verified without needing to trace back to the origin contract. It is compatible with any blockchain that operates on the Bitcoin Virtual Machine (BVM) and supports basic opcodes.

In a typical UTXO model, control is held by the private key holder, and the smart contract is used for verifying token flow by examining its image. In contrast, our approach sits between these two methods.

### Key Features:

- **Control of Auth Coin**: The private key holder retains full control over the coin.
- **Smart Contract (Non-technical Definition)**: For uniqueness, we define the private key as the "smart contract." The private key holder signs each transaction in the chain, with each child UTXO referencing its parent UTXO via a signature (parentTxID + outputIndex). This structure ensures each parent UTXO is verifiable, without needing to trace back to the root.

### Benefits:

1. **No Need for Origin Traceback**: This eliminates the need to trace back to the origin, which is particularly advantageous for ticketing solutions.
2. **Support for Combined Signatures**: This feature is useful for gaming applications, allowing for combined signatures and transfers.
3. **Flexibility in Script Usage**: You can add additional op-scripts, enabling the implementation of conditions such as security tokens, where only the issuer’s private key can burn or transfer the token. This also facilitates on-chain verification.

### Potential Downsides (or Advantages, depending on Use Case):

1. **Full Control in the Hands of the Private Key Holder**: The private key holder has complete control over the token, including the ability to increase or decrease the supply. This is beneficial for business models or applications that require burning tokens, such as stablecoins or ticketing systems, where supply can be adjusted over time.
2. **Risk of Unauthorized Burns**: If multisig is not implemented on-chain, anyone could potentially burn the token.

### Approach Overview:

The protocol concept is similar to embedding the `<sig>` of the issuer’s private key into each child UTXO’s locking script. This approach effectively treats the issuer’s private key as a unique off-chain smart contract, which can be authorized using a ledger. Ultimately, the idea is to use cryptography and the ledger for authorization, as opposed to relying solely on op-scripts and the ledger.

The entire concept will be open source, allowing anyone to integrate it. For any inquiries or questions, feel free to reach out at nikhilmatta10@gmail.com.
