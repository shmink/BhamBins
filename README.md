# BhamBins

A quick and dirty data scrapper for the [Birmingham City councils bin collection day lookup](https://www.birmingham.gov.uk/xfp/form/619).

### Usage
```
❯ go get
❯ go run collection.go -p "B17 0LY" -u "100070285236"

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

- `-p` is for postcode
- `-u` is for UPRN which you can [look up here](https://www.findmyaddress.co.uk/search)

### Limitations

- It's really basic. Only seems to work with houses, so flats with shared bins are out.
- Hard coded values, so if the council change their minds of the web page address this won't work