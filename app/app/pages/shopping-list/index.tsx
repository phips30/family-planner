import {StyleSheet, ScrollView, View} from 'react-native';
import {
    TextInput, Button, Portal, Dialog, Text, Card,
    IconButton, MD3Colors, ActivityIndicator
} from 'react-native-paper';
import React, {useContext, useEffect, useState} from 'react';
import { SafeAreaView } from 'react-native-safe-area-context';
import {User, UserContext} from "@/app/user.provider";
import axios from "axios";
import {API_URL} from "@/app/constants";

interface ShoppingItem {
    id: string;
    name: string;
    user: User;
    addedAt: Date;
    bought: boolean;
}

export namespace ShoppingItem {
    export function createShoppingItem(name: string, user: User): ShoppingItem {
        return {name: name, user: user, addedAt: new Date(), bought: false} as ShoppingItem;
    }
}

export default function ShoppingList() {
    const {loggedInUser} = useContext(UserContext);

    const [newItemName, setNewItemName] = useState('');
    const [shoppingList, setShoppingList] = useState<ShoppingItem[]>([]);
    const [shoppingItemAlreadyExistsDlgVisible, setShoppingItemAlreadyExistsDlgVisible] = useState(false);

    const hideShoppingItemAlreadyExistsDlg = () => setShoppingItemAlreadyExistsDlgVisible(false);

    useEffect(() => {
        const shoppinglistUri = `${API_URL}/v1/shopping-list/${loggedInUser.groupId ? 'group/' : 'user/'}${loggedInUser.groupId ? loggedInUser.groupId : loggedInUser.id}`;
        axios.get<ShoppingItem[]>(shoppinglistUri)
            .then(response => {
                setShoppingList(response.data);
            })
            .catch(err => {
                console.error(err);
            });
    }, [])

    function addShoppingItem(name: string): void {
        if (shoppingList.findIndex(item => item.name === name) == -1) {
            const newItem = ShoppingItem.createShoppingItem(newItemName, loggedInUser)
            axios.post<string[]>(`${API_URL}/v1/shopping-list/`, newItem, {
                transformRequest: [(item) => {
                    return JSON.stringify([{
                        name: item.name,
                        userId: item.user.id,
                        groupId: item.user.groupId,
                        addedAt: item.addedAt,
                        bought: item.bought
                    }]);
                }]
            }).then(response => {
                newItem.id = response.data[0];
                setShoppingList([
                    ...shoppingList, newItem
                ]);
            }).catch(err => {
                console.error(err);
            });
        } else {
            setShoppingItemAlreadyExistsDlgVisible(true);
        }
    }

    function handleItemBought(itemToUpdate: ShoppingItem) {
        const shoppingListCopy = [...shoppingList];
        const theItem = shoppingListCopy.find(
            item => item.name === itemToUpdate.name
        )

        if (theItem) {
            theItem.bought = !theItem.bought;
            axios.put<ShoppingItem>(`${API_URL}/v1/shopping-list/`, theItem, {
                transformRequest: [(item) => {
                    return JSON.stringify({
                        id: item.id,
                        name: item.name,
                        bought: item.bought
                    });
                }]
            }).then(() => {
                setShoppingList(shoppingListCopy);
            }).catch(err => {
                console.error(err);
            });
        }
    }

    function removeItem(itemToDelete: ShoppingItem) {
        axios.delete<ShoppingItem>(`${API_URL}/v1/shopping-list/${itemToDelete.id}`)
            .then(() => {
                setShoppingList(shoppingList.filter(item => item.name !== itemToDelete.name));
            }).catch(err => {
                console.error(err);
            });
    }

    if (!shoppingList) {
        return <ActivityIndicator />;
    }
    return (
        <>
            <SafeAreaView
                style={[
                    styles.container,
                    {
                        flexDirection: 'column',
                    },
                ]}>
                <SafeAreaView style={{flex: 9}}>
                    <ScrollView>
                        {shoppingList.map((item, index) => (
                            <View key={index}>
                                <Card>
                                    <Card.Title
                                        title={item.name}
                                        subtitle={"Added by: " + item.user.name}
                                        left={() =>
                                            <IconButton
                                                icon="delete"
                                                iconColor={MD3Colors.error50}
                                                size={20}
                                                onPress={() => removeItem(item)}/>
                                        }
                                        right={() =>
                                            <IconButton
                                                icon={item.bought ? "check" : "crop-square"}
                                                size={20}
                                                onPress={() => handleItemBought(item)}/>
                                        }
                                    />
                                </Card>
                            </View>
                        ))}
                    </ScrollView>
                </SafeAreaView>

                <SafeAreaView style={{flex: 2}}>
                    <TextInput
                        style={styles.input}
                        label="Item"
                        value={newItemName}
                        onChangeText={text => setNewItemName(text)}
                    />

                    <Button icon="plus" mode="contained"
                            style={styles.button}
                            disabled={newItemName.length == 0}
                            onPress={() => addShoppingItem(newItemName)}>
                        Add Item
                    </Button>

                    <Portal>
                        <Dialog visible={shoppingItemAlreadyExistsDlgVisible}
                                onDismiss={hideShoppingItemAlreadyExistsDlg}>
                            <Dialog.Title>Item already exists</Dialog.Title>
                            <Dialog.Content>
                                <Text variant="displaySmall">The item already exists in the list and cannot be added
                                    twice.</Text>
                            </Dialog.Content>
                            <Dialog.Actions>
                                <Button onPress={hideShoppingItemAlreadyExistsDlg}>Ok</Button>
                            </Dialog.Actions>
                        </Dialog>
                    </Portal>


                </SafeAreaView>
            </SafeAreaView>
        </>
    );
}

const styles = StyleSheet.create({
container: {
flex: 1,
padding: 10,
},
input: {
marginBottom: 16,
},
button: {
marginTop: 8,
},
});