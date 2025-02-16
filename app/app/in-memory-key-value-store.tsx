import AsyncStorage from '@react-native-async-storage/async-storage';

export class InMemoryDbClass {

    storeData = async (key: string, value: string) => {
        try {
            await AsyncStorage.setItem(key, value);
        } catch (e) {
            // saving error
        }
    };

    getData = async (key: string) => {
        try {
            const value = await AsyncStorage.getItem(key);
            if (value !== null) {
                return value;
            }
        } catch (e) {
            // error reading value
        }
    };
}

const storeData = async (key: string, value: string) => {
    try {
        await AsyncStorage.setItem(key, value);
    } catch (e) {
        // saving error
    }
};

const storeObject = async (key: string, object: any) => {
    try {
        const jsonValue = JSON.stringify(object);
        await AsyncStorage.setItem(key, jsonValue);
    } catch (e) {
        // saving error
    }
};

const getData = async (key: string) => {
    const value = await AsyncStorage.getItem(key);
    if (value !== null) {
        return value;
    }
    throw "Key {key} not found";
};

const getObject = async (key: string) => {
    const jsonValue = await AsyncStorage.getItem(key);
    if (jsonValue !== null) {
        return JSON.parse(jsonValue);
    }
    throw "Key " + key + " not found";
};

export const InMemoryDb = {
    storeData,
    storeObject,
    getData,
    getObject
};
