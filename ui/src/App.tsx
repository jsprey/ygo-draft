import React from 'react';
import './App.css';
import AppRouter from "./routes/AppRouter";
import AuthProvider from "./auth/AuthProvider";
import ColorThemeProvider from "./core/context/ColorThemeProvider";

function App() {
    return (<>
            <AuthProvider>
                <ColorThemeProvider>
                    <AppRouter/>
                </ColorThemeProvider>
            </AuthProvider>
        </>
    );
}

export default App;
