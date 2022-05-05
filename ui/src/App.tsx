import React from 'react';
import './App.css';
import {QueryClient, QueryClientProvider,} from 'react-query'
import DeckViewer from "./components/DeckViewer";
import {useCard} from "./components/hooks/useRandomCard";
import {useRandomCards} from "./components/hooks/useRandomCards";
import RandomDeckGenerator from "./components/RandomDeckGenerator";

// Create a client
const queryClient = new QueryClient()

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            <div className="App">
                <RandomDeckGenerator/>
            </div>
        </QueryClientProvider>
    );
}

export default App;
