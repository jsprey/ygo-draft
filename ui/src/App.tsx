import React from 'react';
import './App.css';
import AppRouter from "./routes/AppRouter";
import AuthProvider from "./auth/AuthProvider";
import ColorThemeProvider from "./core/context/ColorThemeProvider";
import {SnackbarProvider} from "notistack";

function App() {
    return (<>
            <SnackbarProvider
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'center',
                }}
            >
                <AuthProvider>
                    <ColorThemeProvider>
                        <AppRouter/>
                    </ColorThemeProvider>
                </AuthProvider>
            </SnackbarProvider>
        </>
    );
}

export default App;
