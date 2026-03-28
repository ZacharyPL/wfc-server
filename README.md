# rwfc

Retro Wi-Fi Connection is a fork of WiiLink WFC. This repository contains numerous changes to WWFC to support Retro Rewind.

## Changes from WWFC

- Extended and rewritten APIs to facilitate [rwfc-web](https://github.com/Retro-Rewind-Team/rwfc-web/), [wfc-bot](https://github.com/Retro-Rewind-Team/wfc-bot/)
  - Discord linking with a license
  - Improved moderation
    - Querying a profile for bans and other identifiers
    - Querying a Mii for a given profile
    - Searching for users by a given identifiers
  - Setting the Message of the Day without resetting the server
- Extended max VR (Up to 30k, Retro Rewind's higher VR is handled client side)
- Client hash checking to ensure only clients on the most up-to-date version may log in
- Adjustments to ban logic to better facilitate nand-less play
- Improved translations and extended language support
- Reporting of kick/ban reasons to the client
- QR2 kick ordering to force clients to drop kicked clients
- Open Host reporting via the group's API

## Setup

**PREFACE**: All RWFC projects are codependent, in that a mix and match of RWFC
and WFC or vanilla Mario Kart Wii projects and tooling are not guaranteed to
work or be compatible. This fork of wfc-server **WILL NOT** function without
using Retro Rewind's modified payload and modified Pulsar.

You will need:
- A Go compiler (minimum 1.25.5)
- A copy of the payload (see [wfc-patcher-wii](https://github.com/Retro-Rewind-Team/wfc-patcher-wii/) for instructions)
    - The payload's `dist` folder should be copied into the same folder as the executable and named `payload`
- PostgreSQL

1. Create a PostgreSQL database. Note the database name, username, and
   password.
2. Use the `schema.sql` found in the root of this repo and import it into your
   PostgreSQL database.
3. Copy `config-example.xml` to `config.xml` and insert all the correct data.
4. Run `go build`. The executable `wwfc` will appear in the current directory.

#### Hashing

Since clients are required to submit a hash to connect, this must be configured
on launch. You can either disable hash checking (see `enableHashCheck` in
`config.xml`) or use [wfc-bot](https://github.com/Retro-Rewind-Team/wfc-bot/)
to submit hashes.
