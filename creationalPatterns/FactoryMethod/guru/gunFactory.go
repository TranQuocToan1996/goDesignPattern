package creationalPatterns

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "musket" {
		return newM4A1(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
