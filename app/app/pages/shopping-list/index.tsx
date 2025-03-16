import {StyleSheet, ScrollView, View} from 'react-native';
import {
    TextInput, Button, Portal, Dialog, Text, Card,
    IconButton, MD3Colors
} from 'react-native-paper';
import React, {useContext, useEffect, useState} from 'react';
import {router} from "expo-router";
import { SafeAreaView } from 'react-native-safe-area-context';
import {InMemoryDb} from "@/app/in-memory-key-value-store";
import {User, UserContext} from "@/app/user.provider";
import axios from "axios";
import {API_URL} from "@/app/constants";

interface ShoppingItem {
    name: string;
    addedBy: User;
    addedAt: Date;
    bought: boolean;
    sorter: number;
}

export namespace ShoppingItem {
    export function createShoppingItem(name: string, addedBy: User, sorter: number): ShoppingItem {
        return {name: name, addedBy: addedBy, addedAt: new Date(), bought: false,  sorter: sorter} as ShoppingItem;
    }
}

export default function ShoppingList() {
    const { loggedInUser } = useContext(UserContext);

    const [newItemName, setNewItemName] = useState('');
    const [shoppingList, setShoppingList] = useState<ShoppingItem[]>([]);
    const [shoppingItemAlreadyExistsDlgVisible, setShoppingItemAlreadyExistsDlgVisible] = useState(false);

    const hideShoppingItemAlreadyExistsDlg = () => setShoppingItemAlreadyExistsDlgVisible(false);

    useEffect(() => {
        InMemoryDb.getData("shopping-list")
            .then(e => {
                console.log("shopping-list: " + e)
                const shoppingListFromDb = JSON.parse(e) as ShoppingItem[];
                setShoppingList(shoppingListFromDb);
            })
            .catch(e => {
                console.error("not found: " + e)
            });
    }, [])

    function openShareModal(): void {
        router.push("./share")
    }

    function addShoppingItem(name: string): void {
        if (shoppingList.findIndex(item => item.name === name) == -1) {
            const updatedShoppingList = [
                ...shoppingList,
                ShoppingItem.createShoppingItem(newItemName, loggedInUser, shoppingList.length++),
            ];
            InMemoryDb.storeObject("shopping-list", updatedShoppingList)
                .then(e => {
                    console.log("list saved", updatedShoppingList)
                    setShoppingList(updatedShoppingList);
                    setNewItemName('');
                    return axios.post<ShoppingItem[]>(`${API_URL}/shopping-list`, updatedShoppingList);
                })
                .then(httpResponse => console.log(httpResponse))
                .catch(e => {console.log(e)});
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
            InMemoryDb.storeObject("shopping-list", shoppingListCopy)
                .then(e => {
                    console.log("list saved", shoppingListCopy)
                    setShoppingList(shoppingListCopy);
                });
        }
    }

    function removeItem(itemToDelete: ShoppingItem, props: { size: number }) {
        setShoppingList(shoppingList.filter(item => item.name !== itemToDelete.name));
        InMemoryDb.storeObject("shopping-list", shoppingList)
            .then(e => console.log("list saved"));
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
                                    subtitle={"Added by: " + item.addedBy.name}
                                    left={(props) =>
                                        <IconButton
                                            icon="delete"
                                            iconColor={MD3Colors.error50}
                                            size={20}
                                            onPress={() => removeItem(item, props)}/>
                                    }
                                    right={(props) =>
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
                        <Dialog visible={shoppingItemAlreadyExistsDlgVisible} onDismiss={hideShoppingItemAlreadyExistsDlg}>
                            <Dialog.Title>Item already exists</Dialog.Title>
                            <Dialog.Content>
                                <Text variant="displaySmall">The item already exists in the list and cannot be added twice.</Text>
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