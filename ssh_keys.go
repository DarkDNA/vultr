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
func (cli *Client) UpdateSSHKey(keyID, name, sshKey string) (string, error) {
	params := map[string]string{
		"name":     name,
		"ssh_key":  sshKey,
		"SSHKEYID": keyID,
	}

	resp, err := cli.RequestMap(params, "/sshkey/update", "POST")
}

// CreateSSHKey creates a new SSH Key for sending to hosts.
func (cli *Client) CreateSSHKey(name, sshKey string) error {
	params := map[string]string{
		"name":    name,
		"ssh_key": sshKey,
	}

	_, err := cli.RequestMap(params, "/sshkey/create", "POST")
	return err
}
