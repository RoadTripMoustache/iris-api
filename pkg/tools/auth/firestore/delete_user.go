package firebase

// DeleteUser - Delete an user in firebase
func (f *FirebaseClient) DeleteUser(userID string) error {
	return f.Client.DeleteUser(f.Context, userID)
}
