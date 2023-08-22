import React from 'react';
import './App.css';
import AppRouter from "./routes/AppRouter";
import AuthProvider from "./auth/AuthProvider";

function App() {
    return (<>
            <AuthProvider>
                <AppRouter/>
            </AuthProvider>
        </>
    );
}

export default App;
