"use strict";

const {
  HEDERA_NETWORK,
  HEDERA_ACCOUNT_ID,
  HEDERA_ACCOUNT_PRIVATE_KEY,
  HEDERA_CREET_TOKEN_ADDRESS,
} = process.env;

module.exports = {
  network: HEDERA_NETWORK.toLowerCase(),
  accountId: HEDERA_ACCOUNT_ID,
  privateKey: HEDERA_ACCOUNT_PRIVATE_KEY,
  tokenId: HEDERA_CREET_TOKEN_ADDRESS
};
