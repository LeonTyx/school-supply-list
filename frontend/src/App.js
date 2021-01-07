import React, {useEffect, useState} from "react";
import {userSession} from "./UserSession";
import Header from "./Components/Header/Header";
import Home from "./Components/Home/Home";
import {HashRouter, Route} from "react-router-dom"
import SupplyList from "./Components/Supply List/SupplyList";
import Account from "./Components/Account/Account";

function App() {
    const [user, setUser] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        // Refresh session every 15 minutes
        setInterval(refreshSession, 900000)
        // Listen to localstorage changes
        // Update user on change
        window.addEventListener('storage', () => {
            setUser(JSON.parse(localStorage.getItem("user")))
        });

        //Fetch user from api
        fetch("/oauth/v1/profile")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
            .then(
                (result) => {
                    setUser(result);
                    localStorage.setItem("user", JSON.stringify(result))
                }, (error) => {
                    setUser(null)
                    localStorage.removeItem("user")
                    setError(error);
                }
            )
    }, [])

    function refreshSession() {
        fetch("/oauth/v1/refresh")
            .then((res) => {
                if (!res.ok) {
                    setError("Unable to refresh session")
                }
            })
    }

    return (
        error === null &&
        <HashRouter>
            <userSession.Provider value={[user, setUser]}>
                <Header/>
                <main>
                    <Route exact path="/" component={Home}/>
                    <Route path="/list/:id" component={SupplyList}/>
                    {user !== null && user !== undefined && (
                        <Route exact path="/account" component={Account}/>
                    )}
                </main>
            </userSession.Provider>
        </HashRouter>
    );
}

export default App