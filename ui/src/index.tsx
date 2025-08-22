import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import App from './App';
import {QueryClient, QueryClientProvider} from "react-query";

const PORT = parseInt(process.env.PORT || "8080", 10)
export const PUBLIC_URL = process.env.PUBLIC_URL || `http://localhost:${PORT}`;

// Create a client so that we can use it with the use query client hook
export const YgoQueryClient = new QueryClient()

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);

root.render(
    <div className={"rootContainer overflow-x-hidden"}>
        <React.StrictMode>
            <QueryClientProvider client={YgoQueryClient}>
                <App/>
            </QueryClientProvider>
        </React.StrictMode>
    </div>
);