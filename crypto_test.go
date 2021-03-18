package starsign

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestHash(t *testing.T) {
	for base64Data, base64Expected := range hashes {
		data, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			t.Fatalf("Failed to hash: %v", err)
		}
		expected, err := base64.StdEncoding.DecodeString(base64Expected)
		if err != nil {
			t.Fatalf("Failed to hash: %v", err)
		}
		in := bytes.NewReader(data)
		actual, err := hash(in)
		if err != nil {
			t.Fatalf("Failed to hash: %v", err)
		}
		if !bytes.Equal(actual, expected) {
			t.Error("Hash test failed")
		}
	}
}

var hashes = map[string]string{
	"F1vOFWK4Ypw9FWA4kOpdU3GAPuwIeUyCPBR2T20g/pJ3IWmkOXU1gPBSF0Kgta1hvq6I5RYN3CRH2oOnU6X7XQ==":                                                                                     "l7z2C4V0iPAAF/3hY5KMoK2bdOVzl/t9yq3t28PgkF3OZK226mxlrttfNL9pKISHn5banuHg2ov4zy+d2YLZJA==",
	"8zlv13JCdsMstihNW9D1h3YohB0sxnBW0k8MDY3gdmY6fvqHq+q5Gd7TlndnnuHdblEEWgMZrU8kPfpA6c8b8Pgc/KAX82j2JFn5vV0QiHYH/GxHRgLvH5/osmqXzpdq":                                             "q/euRaFhnm/5dFL1yKyYTeMy9lahnjnpT3lRr8JduSFgPSLo20hQ6W10GFnl3nlY/GZEvvxJ8DovQtNnMHuMvg==",
	"fL+3tmBj0wcQYfTRjldVbXVampqhCmoeD9b8VR7c0E1bl9h59yBw1W80ctJgeveAZ+wuGeeMaCVWP8Hv3QyOCc0rNOkXKe78e9plE9ftUuU8z7WcLyZT3Tgjkh0+URIo5U3jaodqjPcUKw993+jJ8gwakVnnULQC8aswiqLNx+Q=": "vuiey3DEZnna0a9HY7VGFOVpkZ0NLCdB9+Orfsd4NZeSSPcdL1tmoVJQvE6C1vpD/NXGcWR5o1uV+vr08akzew==",
	"Xdcne/uKnOTNamSUBc3j658o13x8b2CpZEFdN8NzeCyGqtiFVu/hZbwhn+5EaY1Rw/BSCig75edcyKNkZck/Y2iOCVjDktbDA51Vc+6xR0TwIY/jNOCofmQ4Pc4E4Ng4BIE+Sj2BBy9FiU/sygpDutOGHDatwnRNDC8fNeZ121XD+d140V/P48eAzKMgujh6YZ9hY8QOFLxYMMzLEht9RbfoGgYLMfoImi23xeNVS6ggTmwcuU2eWkYzTakL0gWn9ihsPwIujszQVoZ3gaZBMXrAorSm0YaBWc6qujegjK/zx4Ho0UGNxN8OzB1q+QmB1x6JcjvNh22sOJn4pxWdkXeqaYh8W7crguRV0+SRNl8u6jEXZOC/5LNLRGOQksKPfuVgJURAW6vGEX/PicAJ8+EccnX/ZAStKBnozzz9e1hX6vlsaZZBQRI1S9twUiZhM6OGnJU/FZ98pi6sqMWYcD1MeR1uM6+amORak4KMXD4YpiTTSY19/v+Wo4Uq4nCaBkKHz9eMI3InD6EYsbGWsGQ+9TszlKwSF0nHLhIx55K+xTC+5gtvlhGb86jNWx1Nn3CTdpM1JGnpVq1e+cMb40I/8hrxO5Xw6ZrI7o2XBsaKgSni0Ghd4TEOUv1sC9iV4CDu2vc1TYHKRDacnftaO1lHVtGVmQBehhOhNWugseY=": "KY7Btdu/cFF1U2WlNGS5ENkEBE+0vlgxnx5Zp+9fIgKGcyDIKXhvwoNrIABixe2BpdTpi/MKwrxy4GeR+rgZcw==",
}
