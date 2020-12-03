/*
 * GoCloudFlareInfos
 * Copyright (C) 2020-2020 Dametto Luca <https://damettoluca.com>
 *
 * functions.go is part of GoCloudFlareInfos
 *
 * You should have received a copy of the GNU Affero General Public License v3.0
 * along with GoCloudFlareInfos. If not, see <https://github.com/LucaTheHacker/GoCloudFlareInfos/blob/main/LICENSE>.
 */
package GoCFLocation

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetData returns parsed information 
func GetData(website string) (map[string]string, error) {
	resp, err := http.Get(website + "/cdn-cgi/trace")
	if err != nil {
		return map[string]string{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]string{}, err
	}

	result := map[string]string{}
	parts := strings.Split(string(body), "\n")
	for _, v := range parts {
		values := strings.SplitN(v, "=", 2)
		if len(values) == 2 {
			result[values[0]] = values[1]
		}
	}
	return result, nil
}
