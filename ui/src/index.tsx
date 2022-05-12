import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import {QueryClient, QueryClientProvider} from "react-query";
import 'bootstrap/dist/css/bootstrap.min.css';
import {BrowserRouter} from "react-router-dom";

const PORT = parseInt(process.env.PORT || "8080", 10)
export const  PUBLIC_URL = process.env.PUBLIC_URL || `http://localhost:${PORT}`;

// Create a client
export const YgoQueryClient = new QueryClient()

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <React.StrictMode>
        <BrowserRouter>
            <QueryClientProvider client={YgoQueryClient}>
                    <App/>
            </QueryClientProvider>
        </BrowserRouter>
    </React.StrictMode>
);