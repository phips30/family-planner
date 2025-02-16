import {StyleSheet, ScrollView, View} from 'react-native';
import {
    MD3LightTheme as DefaultTheme, Checkbox, TextInput, Button, Divider, Portal, Dialog, PaperProvider, Text
} from 'react-native-paper';
import React, {useEffect, useState} from 'react';
import {router, Stack} from "expo-router";
import { SafeAreaView } from 'react-native-safe-area-context';
import {InMemoryDb} from "@/app/in-memory-key-value-store";

interface ShoppingItem {
    id: string;
    name: string;
    addedBy: string;
    addedAt: Date;
    bought: boolean;
    sorter: number;
}

export namespace ShoppingItem {
    export function createShoppingItem(name: string, sorter: number): ShoppingItem {
        return {id: "", bought: false, name: name, sorter: sorter} as ShoppingItem;
    }
}

export default function ShoppingList() {
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
                ShoppingItem.createShoppingItem(newItemName, shoppingList.length++),
            ];
            InMemoryDb.storeObject("shopping-list", updatedShoppingList).then(e => console.log("list saved"));

            setShoppingList(updatedShoppingList);
            setNewItemName('');
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
            setShoppingList(shoppingListCopy);
        }
    }

    return (
        <>
            <Stack.Screen
                options={{
                    headerRight: () => <Button onPress={() => openShareModal()}>Share</Button>,
                }}
            />
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
                                <Checkbox.Item label={item.name}
                                               status={item.bought ? 'checked' : 'unchecked'}
                                               onPress={() => handleItemBought(item)}/>
                                <Divider />
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