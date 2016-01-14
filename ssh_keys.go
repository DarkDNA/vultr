package vultr

// DestroySSHKey destroys a SSH Key from the vultr databases.
func (cli *Client) DestroySSHKey(keyID string) error {
	params := map[string]string{
		"SSHKEYID": keyID,
	}

	_, err := cli.RequestMap(params, "/sshkey/destroy", "POST")
	return err
}

// UpdateSSHKey updates an existing SSH Key.
func (cli *Client) UpdateSSHKey(keyID, name, sshKey string) error {
	params := map[string]string{
		"name":     name,
		"ssh_key":  sshKey,
		"SSHKEYID": keyID,
	}

	_, err := cli.RequestMap(params, "/sshkey/update", "POST")
	return err
}

// CreateSSHKey creates a new SSH Key for sending to hosts.
func (cli *Client) CreateSSHKey(name, sshKey string) (string, error) {
	params := map[string]string{
		"name":    name,
		"ssh_key": sshKey,
	}

	resp, err := cli.RequestMap(params, "/sshkey/create", "POST")
	if err != nil {
		return "", err
	}

	return resp["SSHKEYID"].(string), nil
}
