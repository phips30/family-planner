import React, {createContext, useEffect, useState} from "react";
import {InMemoryDb} from "@/app/in-memory-key-value-store";
import {ActivityIndicator} from "react-native-paper";

export interface User {
    userId: string;
    userName: string;
}

export interface UserContextProperties {
    loggedInUser: User | null;
    setLoggedInUser: (user: User) => void;
}

export const UserContext = createContext({} as UserContextProperties);

export const UserProvider = ({ children }) => {
    const [isLoadingUser, setIsLoadingUser] = useState<boolean>(true)
    const [loggedInUser, setLoggedInUser] = useState<User | null>({} as User);

    useEffect(() => {
        InMemoryDb.getData("user")
            .then(e => {
                console.log("loggedInUser: " + e)

                const userFromDb = JSON.parse(e) as User;
                setLoggedInUser({
                    ...userFromDb
                });
                setIsLoadingUser(false);
                // To reset db user
                if (userFromDb.userId == "82") {
                    setLoggedInUser(null);
                }
            })
            .catch(e => {
                console.error("not found: " + e)
            });

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