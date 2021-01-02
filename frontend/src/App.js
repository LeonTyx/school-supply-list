import React from "react";
import {userSession} from "./UserSession";
import Main from "./Main";
import Header from "./Header/Header";

function App() {
    return (
        <userSession.Provider value={{user:{name: "Johnny Test"}}}>
            <Header/>
            <Main />
        </userSession.Provider>
    );
}


export default App