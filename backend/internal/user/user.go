package user

type User struct {
    Name string
    DeviceId string
}

func Of(name string, deviceId string) *User {
    user := User{Name: name, DeviceId: deviceId}
    return &user
}