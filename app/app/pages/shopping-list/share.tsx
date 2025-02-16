import {StyleSheet, Image, ScrollView, FlatList, View} from 'react-native';
import {Card, Text, List, Checkbox, TextInput, Button} from 'react-native-paper';

function LogoTitle() {
    return (
        <Image style={styles.image} source={{ uri: 'https://reactnative.dev/img/tiny_logo.png' }} />
    );
}

export default function ShareShoppingList() {

    return (
        <>
            <View style={{flex: 9}}>
                <Text>Share</Text>
            </View>
        </>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 10,
    },
    image: {
        width: 50,
        height: 50,
    },
});
