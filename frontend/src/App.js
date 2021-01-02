import React, {useState} from "react";
import {userSession} from "./UserSession";
import Main from "./Main";
import Header from "./Header/Header";

function App() {
    const [user, setUser] = useState({name:null});

    return (
        <userSession.Provider value={[user, setUser]}>
            <Header/>
            <Main />
        </userSession.Provider>
    );
}


export default App