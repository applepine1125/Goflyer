package Goflyer

func (a *API) GetPermissions() ([]byte, error) {
	path := "/v1/me/getpermissions"

	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")
	return byteSlice, err
}
