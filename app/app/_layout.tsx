import {Stack} from "expo-router";
import {PaperProvider, MD3LightTheme as DefaultTheme} from "react-native-paper";


export default function RootLayout() {

    return (
        <PaperProvider theme={DefaultTheme}>
            <Stack initialRouteName="index"
                screenOptions={{
                    headerStyle: {
                        backgroundColor: '#f4511e',
                    },
                    headerTintColor: '#fff',
                    headerTitleStyle: {
                        fontWeight: 'bold',
                    },
                    headerShown: true
                }}>
                <Stack.Screen name="index" options={{ title: "Family-Planner" }} />
                <Stack.Screen name="details" />
                <Stack.Screen name="pages/shopping-list/index" options={{
                    title: "Shopping-list"
                }} />

            </Stack>
        </PaperProvider>

  );
}
