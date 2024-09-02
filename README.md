# go-temperature-rug

Based on https://github.com/ndewijer/lambda-temperature-rug 
An experiement to see how fast and well I can use LLMs to create the same functionality in a different language I've never used before.
About 4 hours, if you want to know.



This program will generate the lowest and highest temperature of a given KNMI datastation.
You can pass the following variables to the program:

1. start - Startdate(YYYYMMDD)
2. end - Enddate(YYYYMMDD)
3. reverse - Date ordering(true/false)
4. stations - Stations(comma separated list of station numbers)
5. variables - Variables(comma separated list of variables)

Info about how to use the Variables and Stations can be found here: https://www.knmi.nl/kennis-en-datacentrum/achtergrond/data-ophalen-vanuit-een-script (Dutch)

```
----------------------------------------------------------------
| Date       | Min.Temp | Min.Kleur   | Max.Temp | Max.Kleur   |
----------------------------------------------------------------
| 2023-03-01 |   -4.1°C | donkerblauw |    7.1°C | lichtblauw  |
| 2023-03-02 |    1.3°C | blauw       |    6.8°C | lichtblauw  |
| 2023-03-03 |    1.2°C | blauw       |    8.2°C | lichtblauw  |
| 2023-03-04 |    2.2°C | blauw       |    8.5°C | lichtblauw  |
| 2023-03-05 |    0.9°C | blauw       |    6.9°C | lichtblauw  |
| 2023-03-06 |    1.1°C | blauw       |    6.7°C | lichtblauw  |
| 2023-03-07 |   -2.5°C | donkerblauw |    5.4°C | blauw       |
| 2023-03-08 |   -2.7°C | donkerblauw |    3.7°C | blauw       |
| 2023-03-09 |    0.6°C | blauw       |    3.0°C | blauw       |
| 2023-03-10 |    0.0°C | blauw       |    2.6°C | blauw       |
| 2023-03-11 |   -3.0°C | donkerblauw |    7.3°C | lichtblauw  |
| 2023-03-12 |    2.0°C | blauw       |   11.1°C | lichtblauw  |
| 2023-03-13 |    9.1°C | lichtblauw  |   15.3°C | creme       |
| 2023-03-14 |    2.6°C | blauw       |   10.8°C | lichtblauw  |
| 2023-03-15 |    1.1°C | blauw       |    9.1°C | lichtblauw  |
| 2023-03-16 |    5.2°C | blauw       |   13.6°C | creme       |
| 2023-03-17 |    7.7°C | lichtblauw  |   15.8°C | creme       |
| 2023-03-18 |    7.3°C | lichtblauw  |   15.9°C | creme       |
| 2023-03-19 |    7.4°C | lichtblauw  |   11.2°C | lichtblauw  |
| 2023-03-20 |    6.9°C | lichtblauw  |    9.5°C | lichtblauw  |
| 2023-03-21 |    8.8°C | lichtblauw  |   13.2°C | creme       |
----------------------------------------------------------------
```
