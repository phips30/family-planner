import {StyleSheet, ScrollView} from 'react-native';
import React, {useContext, useEffect, useState} from 'react';
import { SafeAreaView } from 'react-native-safe-area-context';
import {User, UserContext} from "@/app/user.provider";
import axios from "axios";
import {API_URL} from "@/app/constants";
import {ActivityIndicator, DataTable} from 'react-native-paper';
import {Stack} from "expo-router";

interface Group {
    groupId: string;
    name: string;
    members: GroupMember[];
}

interface GroupMember {
    user: User;
    joinedAt: Date
}

export default function Group() {
    const { loggedInUser } = useContext(UserContext);
    const [group, setGroup] = useState<Group>();

    useEffect(() => {
        axios.get<Group>(`${API_URL}/group/${loggedInUser.email}`)
            .then(response => {
                setGroup({...response.data});
            })
            .catch(err => {
                console.error(err);
            });
    }, [])

    if (!group) {
        return <ActivityIndicator />;
    }
    return (
        <>
            <Stack.Screen
                options={{
                    title: `Group - ${group?.name}`,
                }}
            />
            <SafeAreaView
                style={[
                    styles.container,
                    {
                        flexDirection: 'column',
                    },
                ]}>
                <SafeAreaView style={{flex: 3}}>
                    <ScrollView>
                        <DataTable>
                            <DataTable.Header>
                                <DataTable.Title>Name</DataTable.Title>
                                <DataTable.Title>Email</DataTable.Title>
                                <DataTable.Title>Joined at</DataTable.Title>
                            </DataTable.Header>

                            {group.members.map((member, index) => (
                                <DataTable.Row key={index}>
                                    <DataTable.Cell>{member.user.name}</DataTable.Cell>
                                    <DataTable.Cell>{member.user.email}</DataTable.Cell>
                                    <DataTable.Cell>{member.joinedAt.toString()}</DataTable.Cell>
                                </DataTable.Row>
                            ))}
                        </DataTable>
                    </ScrollView>
                </SafeAreaView>
            </SafeAreaView>
        </>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 10,
    }
});