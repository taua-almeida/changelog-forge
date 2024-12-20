package versioning

import (
	"fmt"
	"strconv"
	"strings"
)

func IncrementVersion(currentVersion, incrementType string) (string, error) {
	parts := strings.Split(currentVersion, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format")
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch incrementType {
	case "major":
		major++
		minor, patch = 0, 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return "", fmt.Errorf("invalid increment type")
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}
