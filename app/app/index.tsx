import {Button, ActivityIndicator, Text, TextInput} from 'react-native-paper';

import React, {useEffect, useState} from "react";
import {Link} from "expo-router";
import {InMemoryDb} from "@/app/in-memory-key-value-store";
import {View, StyleSheet} from "react-native";

interface User {
    userId: string;
    userName: string;
}

function CreateNewUserForm({createUser}) {
    const [userName, setUserName] = useState('');

    const handleCreateNewUserClick = () => {
        createUser(userName);
    };

    return (
        <View style={styles.containerCenter}>
            <TextInput
                label="Name"
                value={userName}
                onChangeText={text => setUserName(text)}
                style={styles.input}
            />
            <Button mode="contained" onPress={handleCreateNewUserClick} style={styles.button}>Submit</Button>
        </View>
    )
}

export default function HomeScreen() {
    const [loggedInUser, setLoggedInUser] = useState<User | null>(null)
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        InMemoryDb.getData("user")
            .then(e => {
                console.log("loggedInUser: " + e)

                const userFromDb = JSON.parse(e) as User;
                setLoggedInUser({
                    ...userFromDb
                });
                setIsLoading(false);
                // To reset db user
                if(userFromDb.userId == "35") {
                    setLoggedInUser(null);
                }
            })
            .catch(e => {
                console.error("not found: " + e)
            });
    }, []);

    const hasUserData = () => {
        return !isLoading && loggedInUser;
    }

    function storeUser(userName: any) {
        let newUser = {
            userId: Math.floor((Math.random() * 100) + 1).toString(),
            userName: userName
        } as User;

        InMemoryDb.storeData("user", JSON.stringify(newUser))
            .then(e => {
                setLoggedInUser({
                    ...newUser
                });
                console.log("stored", e);
            });
    }

    return (
            <View style={styles.containerCenter}>
            {isLoading ?
                <ActivityIndicator animating={true} /> :
                hasUserData() ?
                    <>
                        <Text variant="displayMedium">Hello {loggedInUser?.userName} {loggedInUser?.userId}</Text>

                        <Link href="/pages/shopping-list" asChild>
                            <Button mode="outlined" style={styles.button}>
                                <Text>Go to shopping list</Text>
                            </Button>
                        </Link>
                    </>
                    :
                    <>
                        <Text variant="displayMedium">Hello Stranger!</Text>
                        <Text variant="displaySmall">Looks like I don´t know you yet.</Text>
                        <Text variant="displaySmall">Please provide your name so we can get started</Text>
                        <CreateNewUserForm createUser={(username: string) => storeUser(username)}/>
                    </>
            }
            </View>
  );
}


const styles = StyleSheet.create({
    containerCenter: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 16,
    },
    input: {
        marginBottom: 16,
    },
    button: {
        marginTop: 8,
    },
});