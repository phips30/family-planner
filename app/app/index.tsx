import {Button, Text, TextInput} from 'react-native-paper';

import React, {useContext, useState} from "react";
import {View, StyleSheet} from "react-native";
import { InMemoryDb } from '@/app/in-memory-key-value-store';
import { Link } from 'expo-router';
import {User, UserContext} from "@/app/user.provider";

function CreateNewUserForm({createUser}) {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');

    const handleCreateNewUserClick = () => {
        createUser(name, email);
    };

    return (
        <View style={styles.containerCenter}>
            <TextInput
                label="Name"
                value={name}
                onChangeText={text => setName(text)}
                style={styles.input}
            />
            <TextInput
                label="Email"
                value={email}
                onChangeText={text => setEmail(text)}
                style={styles.input}
            />
            <Button mode="contained" onPress={handleCreateNewUserClick} style={styles.button}>Submit</Button>
        </View>
    )
}

export default function HomeScreen() {

    const { loggedInUser, setLoggedInUser } = useContext(UserContext);

    const hasUserData = (): boolean => {
        return loggedInUser != null;
    }

    function storeUser(name: any, email: string) {
        let newUser = {
            name: name,
            email: email
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
            {hasUserData() ?
                <>
                    <Text variant="displayMedium">Hello {loggedInUser.name} {loggedInUser.email}</Text>

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
                    <CreateNewUserForm createUser={(name: string, email: string) => storeUser(name, email)}/>
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