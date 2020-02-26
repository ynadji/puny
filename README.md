# `puny`

Encode and decode domains with [Punycode](https://en.wikipedia.org/wiki/Punycode).

## Usage

```
$ echo "mejodastío.com
中国.com
اَلْفُصْحَىٰ.com" | puny
xn--mejodasto-n5a.com
xn--fiqs8s.com
xn--mgbmx1co7a4aeg2ad77a.com

$ echo "xn--mejodasto-n5a.com
xn--fiqs8s.com
xn--mgbmx1co7a4aeg2ad77a.com" | puny --decode
mejodastío.com
中国.com
اَلْفُصْحَىٰ.com

$ echo "mejodastío.com
中国.com
اَلْفُصْحَىٰ.com" | puny | puny --decode
mejodastío.com
中国.com
اَلْفُصْحَىٰ.com
```
