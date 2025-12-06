# BhamBins [![CI](https://github.com/shmink/BhamBins/actions/workflows/ci.yml/badge.svg)](https://github.com/shmink/BhamBins/actions/workflows/ci.yml)

A quick and dirty data scrapper for the [Birmingham City councils bin collection day lookup](https://www.birmingham.gov.uk/xfp/form/619).

The aim is to run in on an arduino with lights to flash different colours to indicate it's bin day. Thought I would seperate out this bit in case it's of use to anyone else

### Usage
```
❯ ./bhambins -p "B17 0LY" -u "100070285236"

{
    "bins": [
        {
            "name": "Household Collection",
            "webDate": "Monday (12th)",
            "actualDate": "2024-08-12"
        },
        {
            "name": "Recycling Collection",
            "webDate": "Monday (12th)",
            "actualDate": "2024-08-12"
        }
    ]
}
```

- `-p`, `--postcode` is for postcode
- `-u`, `--uprn` is for UPRN which you can [look up here](https://www.findmyaddress.co.uk/search)

If you're not interested in the source code you can just [download the binary](https://github.com/shmink/BhamBins/raw/refs/heads/main/bhambins) and run it like you see above.

### Limitations

- It's really basic. Only seems to work with houses, so flats with shared bins are out.
- Hard coded values, so if the council change their minds of the web page address this won't work but I don't have any power over this. (06/12/2025 - they changed how they return dates for example which borked this program for example)

