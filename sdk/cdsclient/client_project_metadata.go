package cdsclient

// client.ProjectMetadataSet(v[_ProjectKey], key)
// client.ProjectMetadataList(v[_ProjectKey])
// client.ProjectMetadataDelete(v[_ProjectKey], v["key-name"])

// func (c *client) ProjectMetadataList(key string) ([]sdk.ProjectKey, error) {
// 	k := []sdk.ProjectKey{}
// 	if _, err := c.GetJSON(context.Background(), "/project/"+key+"/metadata", &k); err != nil {
// 		return nil, err
// 	}
// 	return k, nil
// }

// func (c *client) ProjectMetadataSet(projectKey string, keyProject *sdk.ProjectKey) error {
// 	_, err := c.PostJSON(context.Background(), "/project/"+projectKey+"/metadata", keyProject, keyProject)
// 	return err
// }

// func (c *client) ProjectMetadataDelete(projectKey string, keyName string) error {
// 	_, _, _, err := c.Request(context.Background(), "DELETE", "/project/"+projectKey+"/metadata/"+url.QueryEscape(keyName), nil)
// 	return err
// }
