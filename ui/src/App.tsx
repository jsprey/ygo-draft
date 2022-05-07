import React from 'react';
import './App.css';
import YgoNavbar from "./core/YgoNavbar";
import {Route, Routes} from "react-router-dom";
import Home from "./home/Home";
import YgoNavbar2 from "./core/YgoNavbar2";

function App() {
    return (<>
            <YgoNavbar/>
            <div className={"SiteContainer"}>
                <Routes>
                    <Route path='/' element={<Home/>}/>
                    <Route path="/deck" element={<h1>MyDeck</h1>}/>
                </Routes>
            </div>
        </>
    );
}

export default App;
