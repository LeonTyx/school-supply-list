import React from "react";
import {userContext} from "./userContext";
import Main from "./Main";

function App() {
    return (
        <userContext.Provider value={{user:{name: "hello"}}}>
            <Main />
        </userContext.Provider>
    );
}


export default App