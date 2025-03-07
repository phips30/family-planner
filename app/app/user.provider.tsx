import React, {createContext, useEffect, useState} from "react";
import {InMemoryDb} from "@/app/in-memory-key-value-store";
import {ActivityIndicator} from "react-native-paper";
import axios from "axios";
import { API_URL } from './constants';

export interface User {
    name: string;
    email: string;
}

export interface UserContextProperties {
    loggedInUser: User;
    setLoggedInUser: (user: User) => void;
}

export const UserContext = createContext({} as UserContextProperties);

export const UserProvider = ({ children }) => {
    const [isLoadingUser, setIsLoadingUser] = useState<boolean>(true)
    const [loggedInUser, setLoggedInUser] = useState<User>({} as User);

    useEffect(() => {
        InMemoryDb.getData('user')
            .then((value) => {
                if (value === null) {
                    console.log('user not found in in-memory db');
                } else {
                    console.log("loggedInUser: " + value)
                    const user = JSON.parse(value) as User;
                    setLoggedInUser({...user});

                    return axios.get<User>(`${API_URL}/user`, {
                        params: {
                            email: user.email
                        }
                    });
                }
            })
            .then((response) => {
                // Change this - e.g. just define that app is in offline mode etc.
                setLoggedInUser(Object.assign({}, response?.data));
            })
            .catch((error) => {
                console.error("Error loading user from server", error);
                console.log("Using in memory user information");
            }).finally(() => setIsLoadingUser(false));
    }, []);


    return (
        <UserContext.Provider value={{loggedInUser, setLoggedInUser}}>
            {isLoadingUser ?
                <ActivityIndicator animating={true} /> :
                children
            }
        </UserContext.Provider>
    );
};