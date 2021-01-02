import React, {useEffect, useState} from "react";
import {userSession} from "./UserSession";
import Header from "./Header/Header";
import Home from "./Home";
import {HashRouter, Route} from "react-router-dom"


function App() {
    const [user, setUser] = useState(null);
    const [error, setError] = useState(null)
    useEffect(() => {
        fetch("/oauth/v1/profile")
            .then(res => res.json())
            .then(
                (result) => {
                    setUser(result);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setUser(null)
                    setError(error);
                }
            )
    }, [])

    return (
        <HashRouter>
            <userSession.Provider value={[user, setUser]}>
                <Header/>
                <main>
                    <Route exact path="/" component={Home}/>
                    <Route exact path="/list/:id" component={Home}/>
                </main>
            </userSession.Provider>
        </HashRouter>
    );
}


export default App